package cline

import (
	"os"
	"reflect"
	"testing"
)

func Test_validateCommands(t *testing.T) {
	type args struct {
		commands []Cmd
	}
	tests := []struct {
		name    string
		args    args
		want    []Cmd
		wantErr bool
	}{
		{
			name: "empty command array",
			args: args{
				commands: []Cmd{},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "invalid command name",
			args: args{
				commands: []Cmd{
					{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid command flags",
			args: args{
				commands: []Cmd{
					{Name: "a", Flags: []Flag{}},
				},
			},
			want: []Cmd{
				{
					Name:    "a",
					Summary: "",
					Handler: nil,
					Flags:   nil,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid command flag names",
			args: args{
				commands: []Cmd{
					{Name: "a", Flags: []Flag{FlagBool{Name: ""}}},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid commands with their flags",
			args: args{
				commands: []Cmd{
					{
						Name:    "info",
						Summary: "information",
						Flags: []Flag{
							FlagInt{
								Name:    "trace",
								Summary: "tracing",
								Value:   10,
								Aliases: []string{"t"},
							},
							FlagBool{
								Name:    "verbose",
								Summary: "info details",
								Value:   true,
								Aliases: []string{"v"},
							},
						},
					},
				},
			},
			want: []Cmd{
				{
					Name:    "info",
					Summary: "information",
					Flags: []Flag{
						FlagInt{
							Name:          "trace",
							Summary:       "tracing",
							Value:         10,
							Aliases:       []string{"t"},
							EnvVar:        "",
							zflag:         FlagValue("10"),
							zflagAssigned: false,
						},
						FlagBool{
							Name:          "verbose",
							Summary:       "info details",
							Value:         true,
							Aliases:       []string{"v"},
							EnvVar:        "",
							zflag:         FlagValue("true"),
							zflagAssigned: false,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateCommands(tt.args.commands)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateCommands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateAndInitFlags(t *testing.T) {
	type args struct {
		flags []Flag
	}

	// for test purposes (TEST: `invalid flag names`)
	os.Setenv("ENV_VERBOSE", "true")

	tests := []struct {
		name    string
		args    args
		want    []Flag
		wantErr bool
	}{
		{
			name: "empty flag array",
			args: args{
				flags: []Flag{},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "invalid FlagString name",
			args: args{
				flags: []Flag{
					FlagString{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid FlagBool name",
			args: args{
				flags: []Flag{
					FlagBool{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid FlagInt name",
			args: args{
				flags: []Flag{
					FlagInt{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid FlagStringSlice name",
			args: args{
				flags: []Flag{
					FlagStringSlice{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid flag type value",
			args: args{
				flags: []Flag{
					struct{ ok bool }{false},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid default flag values",
			args: args{
				flags: []Flag{
					FlagString{Name: "str"},
					FlagInt{Name: "int"},
					FlagStringSlice{Name: "slice"},
					FlagBool{Name: "bool"},
				},
			},
			want: []Flag{
				FlagString{
					Name:          "str",
					Summary:       "",
					Aliases:       []string(nil),
					Value:         "",
					EnvVar:        "",
					zflag:         FlagValue(""),
					zflagAssigned: false,
				},
				FlagInt{
					Name:          "int",
					Summary:       "",
					Aliases:       []string(nil),
					Value:         0,
					EnvVar:        "",
					zflag:         FlagValue("0"),
					zflagAssigned: false,
				},
				FlagStringSlice{
					Name:          "slice",
					Summary:       "",
					Aliases:       []string(nil),
					Value:         []string(nil),
					EnvVar:        "",
					zflag:         FlagValue(""),
					zflagAssigned: false,
				},
				FlagBool{
					Name:          "bool",
					Summary:       "",
					Aliases:       []string(nil),
					Value:         false,
					EnvVar:        "",
					zflag:         FlagValue("false"),
					zflagAssigned: false,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid flag names",
			args: args{
				flags: []Flag{
					FlagString{Name: ""},
					FlagBool{Name: ""},
					FlagStringSlice{Name: ""},
					FlagInt{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid flags and env values",
			args: args{
				flags: []Flag{
					FlagString{
						Name:    "file",
						Summary: "summary 1",
						Value:   "default.go",
						Aliases: []string{"f"},
					},
					FlagBool{
						Name:    "verbose",
						Summary: "summary 2",
						Value:   false,
						Aliases: []string{"V"},
						EnvVar:  "ENV_VERBOSE",
					},
				},
			},
			want: []Flag{
				FlagString{
					Name:          "file",
					Summary:       "summary 1",
					Value:         "default.go",
					Aliases:       []string{"f"},
					zflag:         FlagValue("default.go"),
					zflagAssigned: false,
				},
				FlagBool{
					Name:          "verbose",
					Summary:       "summary 2",
					Value:         false,
					Aliases:       []string{"V"},
					EnvVar:        "ENV_VERBOSE",
					zflag:         FlagValue("true"),
					zflagAssigned: false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateFlagsAndInit(tt.args.flags)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateAndInitFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateAndInitFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFlagByKey(t *testing.T) {
	type args struct {
		key   string
		flags []Flag
	}
	tests := []struct {
		name          string
		args          args
		wantIndex     int
		wantItem      Flag
		wantFlagShort bool
	}{
		{
			name: "search on empty flag array",
			args: args{
				key:   "x",
				flags: nil,
			},
			wantIndex: -1,
			wantItem:  nil,
		},
		{
			name: "search a valid FlagStringSlice by name",
			args: args{
				key: "slice",
				flags: []Flag{
					FlagString{Name: "string"},
					FlagInt{Name: "int"},
					FlagStringSlice{Name: "slice"},
					FlagBool{Name: "bool"},
				},
			},
			wantIndex: 2,
			wantItem: FlagStringSlice{
				Name:          "slice",
				Summary:       "",
				Value:         nil,
				Aliases:       nil,
				EnvVar:        "",
				zflag:         FlagValue(""),
				zflagAssigned: false,
			},
		},
		{
			name: "search a valid FlagStringSlice by short name",
			args: args{
				key: "c",
				flags: []Flag{
					FlagString{Name: "string", Aliases: []string{"s"}},
					FlagInt{Name: "int", Aliases: []string{"i"}},
					FlagStringSlice{Name: "slice", Aliases: []string{"c"}},
					FlagBool{Name: "bool", Aliases: []string{"b"}},
				},
			},
			wantIndex: 2,
			wantItem: FlagStringSlice{
				Name:          "slice",
				Summary:       "",
				Value:         nil,
				Aliases:       []string{"c"},
				EnvVar:        "",
				zflag:         FlagValue(""),
				zflagAssigned: false,
			},
			wantFlagShort: true,
		},
		{
			name: "search a valid FlagString by name",
			args: args{
				key: "string",
				flags: []Flag{
					FlagInt{Name: "int"},
					FlagStringSlice{Name: "slice"},
					FlagBool{Name: "bool"},
					FlagString{Name: "string"},
				},
			},
			wantIndex: 3,
			wantItem: FlagString{
				Name:          "string",
				Summary:       "",
				Value:         "",
				Aliases:       nil,
				EnvVar:        "",
				zflag:         FlagValue(""),
				zflagAssigned: false,
			},
		},
		{
			name: "search a valid FlagString by short name",
			args: args{
				key: "s",
				flags: []Flag{
					FlagInt{Name: "int", Aliases: []string{"i"}},
					FlagStringSlice{Name: "slice", Aliases: []string{"c"}},
					FlagBool{Name: "bool", Aliases: []string{"b"}},
					FlagString{Name: "string", Aliases: []string{"s"}},
				},
			},
			wantIndex: 3,
			wantItem: FlagString{
				Name:          "string",
				Summary:       "",
				Value:         "",
				Aliases:       []string{"s"},
				EnvVar:        "",
				zflag:         FlagValue(""),
				zflagAssigned: false,
			},
			wantFlagShort: true,
		},
		{
			name: "search a valid FlagBool by name",
			args: args{
				key: "bool",
				flags: []Flag{
					FlagInt{Name: "int"},
					FlagStringSlice{Name: "slice"},
					FlagBool{Name: "bool"},
					FlagString{Name: "string"},
				},
			},
			wantIndex: 2,
			wantItem: FlagBool{
				Name:          "bool",
				Summary:       "",
				Value:         false,
				Aliases:       nil,
				EnvVar:        "",
				zflag:         FlagValue(""),
				zflagAssigned: false,
			},
		},
		{
			name: "search a valid FlagBool by short name",
			args: args{
				key: "b",
				flags: []Flag{
					FlagInt{Name: "int", Aliases: []string{"i"}},
					FlagStringSlice{Name: "slice", Aliases: []string{"c"}},
					FlagBool{Name: "bool", Aliases: []string{"b"}},
					FlagString{Name: "string", Aliases: []string{"s"}},
				},
			},
			wantIndex: 2,
			wantItem: FlagBool{
				Name:          "bool",
				Summary:       "",
				Value:         false,
				Aliases:       []string{"b"},
				EnvVar:        "",
				zflag:         FlagValue(""),
				zflagAssigned: false,
			},
			wantFlagShort: true,
		},
		{
			name: "search a valid FlagInt by name",
			args: args{
				key: "int",
				flags: []Flag{
					FlagStringSlice{Name: "slice"},
					FlagInt{Name: "int"},
					FlagBool{Name: "bool"},
					FlagString{Name: "string"},
				},
			},
			wantIndex: 1,
			wantItem: FlagInt{
				Name:          "int",
				Summary:       "",
				Value:         0,
				Aliases:       nil,
				EnvVar:        "",
				zflag:         FlagValue(""),
				zflagAssigned: false,
			},
		},
		{
			name: "search a valid FlagInt by short name",
			args: args{
				key: "i",
				flags: []Flag{
					FlagStringSlice{Name: "slice", Aliases: []string{"c"}},
					FlagInt{Name: "int", Aliases: []string{"i"}},
					FlagBool{Name: "bool", Aliases: []string{"b"}},
					FlagString{Name: "string", Aliases: []string{"s"}},
				},
			},
			wantIndex: 1,
			wantItem: FlagInt{
				Name:          "int",
				Summary:       "",
				Value:         0,
				Aliases:       []string{"i"},
				EnvVar:        "",
				zflag:         FlagValue(""),
				zflagAssigned: false,
			},
			wantFlagShort: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := findFlagByKey(tt.args.key, tt.args.flags)
			if got != tt.wantIndex {
				t.Errorf("findFlagByKey() got = %v, want %v", got, tt.wantIndex)
			}
			if !reflect.DeepEqual(got1, tt.wantItem) {
				t.Errorf("findFlagByKey() got1 = %v, want %v", got1, tt.wantItem)
			}
			if got2 != tt.wantFlagShort {
				t.Errorf("findFlagByKey() got2 = %v, want %v", got2, tt.wantFlagShort)
			}
		})
	}
}
