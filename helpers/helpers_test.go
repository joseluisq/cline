package helpers_test

import (
	"os"
	"testing"

	"github.com/joseluisq/cline/app"
	"github.com/joseluisq/cline/flag"
	"github.com/joseluisq/cline/helpers"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateCommands(t *testing.T) {
	type args struct {
		commands []app.Cmd
	}
	tests := []struct {
		name    string
		args    args
		want    []app.Cmd
		wantErr bool
	}{
		{
			name: "empty command array",
			args: args{
				commands: []app.Cmd{},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "invalid command name",
			args: args{
				commands: []app.Cmd{
					{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid command flags",
			args: args{
				commands: []app.Cmd{
					{Name: "a", Flags: []flag.Flag{}},
				},
			},
			want: []app.Cmd{
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
				commands: []app.Cmd{
					{Name: "a", Flags: []flag.Flag{flag.FlagBool{Name: ""}}},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid commands with their flags",
			args: args{
				commands: []app.Cmd{
					{
						Name:    "info",
						Summary: "information",
						Flags: []flag.Flag{
							flag.FlagInt{
								Name:    "trace",
								Summary: "tracing",
								Value:   10,
								Aliases: []string{"t"},
							},
							flag.FlagBool{
								Name:    "verbose",
								Summary: "info details",
								Value:   true,
								Aliases: []string{"v"},
							},
						},
					},
				},
			},
			want: []app.Cmd{
				{
					Name:    "info",
					Summary: "information",
					Flags: []flag.Flag{
						flag.FlagInt{
							Name:         "trace",
							Summary:      "tracing",
							Value:        10,
							Aliases:      []string{"t"},
							EnvVar:       "",
							FlagValue:    flag.Value("10"),
							FlagAssigned: false,
						},
						flag.FlagBool{
							Name:         "verbose",
							Summary:      "info details",
							Value:        true,
							Aliases:      []string{"v"},
							EnvVar:       "",
							FlagValue:    flag.Value("true"),
							FlagAssigned: false,
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actualCmds, actualErr := helpers.ValidateCommands(tt.args.commands); tt.wantErr {
				assert.Error(t, actualErr, "Expected an error but got none")
			} else {
				assert.NoError(t, actualErr, "Expected no error but got one")
				assert.Equal(t, actualCmds, tt.want, "Commands do not match the expected value")
			}
		})
	}
}

func Test_ValidateAndInitFlags(t *testing.T) {
	type args struct {
		flags []flag.Flag
	}

	// for test purposes (TEST: `invalid flag names`)
	os.Setenv("ENV_VERBOSE", "true")

	tests := []struct {
		name    string
		args    args
		want    []flag.Flag
		wantErr bool
	}{
		{
			name: "empty flag array",
			args: args{
				flags: []flag.Flag{},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "invalid FlagString name",
			args: args{
				flags: []flag.Flag{
					flag.FlagString{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid FlagBool name",
			args: args{
				flags: []flag.Flag{
					flag.FlagBool{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid FlagInt name",
			args: args{
				flags: []flag.Flag{
					flag.FlagInt{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid FlagStringSlice name",
			args: args{
				flags: []flag.Flag{
					flag.FlagStringSlice{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid flag type value",
			args: args{
				flags: []flag.Flag{
					struct{ ok bool }{false},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid default flag values",
			args: args{
				flags: []flag.Flag{
					flag.FlagString{Name: "str"},
					flag.FlagInt{Name: "int"},
					flag.FlagStringSlice{Name: "slice"},
					flag.FlagBool{Name: "bool"},
				},
			},
			want: []flag.Flag{
				flag.FlagString{
					Name:         "str",
					Summary:      "",
					Aliases:      []string(nil),
					Value:        "",
					EnvVar:       "",
					FlagValue:    flag.Value(""),
					FlagAssigned: false,
				},
				flag.FlagInt{
					Name:         "int",
					Summary:      "",
					Aliases:      []string(nil),
					Value:        0,
					EnvVar:       "",
					FlagValue:    flag.Value("0"),
					FlagAssigned: false,
				},
				flag.FlagStringSlice{
					Name:         "slice",
					Summary:      "",
					Aliases:      []string(nil),
					Value:        []string(nil),
					EnvVar:       "",
					FlagValue:    flag.Value(""),
					FlagAssigned: false,
				},
				flag.FlagBool{
					Name:         "bool",
					Summary:      "",
					Aliases:      []string(nil),
					Value:        false,
					EnvVar:       "",
					FlagValue:    flag.Value("false"),
					FlagAssigned: false,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid flag names",
			args: args{
				flags: []flag.Flag{
					flag.FlagString{Name: ""},
					flag.FlagBool{Name: ""},
					flag.FlagStringSlice{Name: ""},
					flag.FlagInt{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid flags and env values",
			args: args{
				flags: []flag.Flag{
					flag.FlagString{
						Name:    "file",
						Summary: "summary 1",
						Value:   "default.go",
						Aliases: []string{"f"},
					},
					flag.FlagBool{
						Name:    "verbose",
						Summary: "summary 2",
						Value:   false,
						Aliases: []string{"V"},
						EnvVar:  "ENV_VERBOSE",
					},
				},
			},
			want: []flag.Flag{
				flag.FlagString{
					Name:         "file",
					Summary:      "summary 1",
					Value:        "default.go",
					Aliases:      []string{"f"},
					FlagValue:    flag.Value("default.go"),
					FlagAssigned: false,
				},
				flag.FlagBool{
					Name:         "verbose",
					Summary:      "summary 2",
					Value:        false,
					Aliases:      []string{"V"},
					EnvVar:       "ENV_VERBOSE",
					FlagValue:    flag.Value("true"),
					FlagAssigned: false,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actualFlags, actualErr := helpers.ValidateFlagsAndInit(tt.args.flags); tt.wantErr {
				assert.Error(t, actualErr, "Expected an error but got none")
			} else {
				assert.NoError(t, actualErr, "Expected no error but got one")
				assert.Equal(t, actualFlags, tt.want, "Flags do not match the expected value")
			}
		})
	}
}

func Test_FindFlagByKey(t *testing.T) {
	type args struct {
		key   string
		flags []flag.Flag
	}
	tests := []struct {
		name          string
		args          args
		wantIndex     int
		wantItem      flag.Flag
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
				flags: []flag.Flag{
					flag.FlagString{Name: "string"},
					flag.FlagInt{Name: "int"},
					flag.FlagStringSlice{Name: "slice"},
					flag.FlagBool{Name: "bool"},
				},
			},
			wantIndex: 2,
			wantItem: flag.FlagStringSlice{
				Name:         "slice",
				Summary:      "",
				Value:        nil,
				Aliases:      nil,
				EnvVar:       "",
				FlagValue:    flag.Value(""),
				FlagAssigned: false,
			},
		},
		{
			name: "search a valid FlagStringSlice by short name",
			args: args{
				key: "c",
				flags: []flag.Flag{
					flag.FlagString{Name: "string", Aliases: []string{"s"}},
					flag.FlagInt{Name: "int", Aliases: []string{"i"}},
					flag.FlagStringSlice{Name: "slice", Aliases: []string{"c"}},
					flag.FlagBool{Name: "bool", Aliases: []string{"b"}},
				},
			},
			wantIndex: 2,
			wantItem: flag.FlagStringSlice{
				Name:         "slice",
				Summary:      "",
				Value:        nil,
				Aliases:      []string{"c"},
				EnvVar:       "",
				FlagValue:    flag.Value(""),
				FlagAssigned: false,
			},
			wantFlagShort: true,
		},
		{
			name: "search a valid FlagString by name",
			args: args{
				key: "string",
				flags: []flag.Flag{
					flag.FlagInt{Name: "int"},
					flag.FlagStringSlice{Name: "slice"},
					flag.FlagBool{Name: "bool"},
					flag.FlagString{Name: "string"},
				},
			},
			wantIndex: 3,
			wantItem: flag.FlagString{
				Name:         "string",
				Summary:      "",
				Value:        "",
				Aliases:      nil,
				EnvVar:       "",
				FlagValue:    flag.Value(""),
				FlagAssigned: false,
			},
		},
		{
			name: "search a valid FlagString by short name",
			args: args{
				key: "s",
				flags: []flag.Flag{
					flag.FlagInt{Name: "int", Aliases: []string{"i"}},
					flag.FlagStringSlice{Name: "slice", Aliases: []string{"c"}},
					flag.FlagBool{Name: "bool", Aliases: []string{"b"}},
					flag.FlagString{Name: "string", Aliases: []string{"s"}},
				},
			},
			wantIndex: 3,
			wantItem: flag.FlagString{
				Name:         "string",
				Summary:      "",
				Value:        "",
				Aliases:      []string{"s"},
				EnvVar:       "",
				FlagValue:    flag.Value(""),
				FlagAssigned: false,
			},
			wantFlagShort: true,
		},
		{
			name: "search a valid FlagBool by name",
			args: args{
				key: "bool",
				flags: []flag.Flag{
					flag.FlagInt{Name: "int"},
					flag.FlagStringSlice{Name: "slice"},
					flag.FlagBool{Name: "bool"},
					flag.FlagString{Name: "string"},
				},
			},
			wantIndex: 2,
			wantItem: flag.FlagBool{
				Name:         "bool",
				Summary:      "",
				Value:        false,
				Aliases:      nil,
				EnvVar:       "",
				FlagValue:    flag.Value(""),
				FlagAssigned: false,
			},
		},
		{
			name: "search a valid FlagBool by short name",
			args: args{
				key: "b",
				flags: []flag.Flag{
					flag.FlagInt{Name: "int", Aliases: []string{"i"}},
					flag.FlagStringSlice{Name: "slice", Aliases: []string{"c"}},
					flag.FlagBool{Name: "bool", Aliases: []string{"b"}},
					flag.FlagString{Name: "string", Aliases: []string{"s"}},
				},
			},
			wantIndex: 2,
			wantItem: flag.FlagBool{
				Name:         "bool",
				Summary:      "",
				Value:        false,
				Aliases:      []string{"b"},
				EnvVar:       "",
				FlagValue:    flag.Value(""),
				FlagAssigned: false,
			},
			wantFlagShort: true,
		},
		{
			name: "search a valid FlagInt by name",
			args: args{
				key: "int",
				flags: []flag.Flag{
					flag.FlagStringSlice{Name: "slice"},
					flag.FlagInt{Name: "int"},
					flag.FlagBool{Name: "bool"},
					flag.FlagString{Name: "string"},
				},
			},
			wantIndex: 1,
			wantItem: flag.FlagInt{
				Name:         "int",
				Summary:      "",
				Value:        0,
				Aliases:      nil,
				EnvVar:       "",
				FlagValue:    flag.Value(""),
				FlagAssigned: false,
			},
		},
		{
			name: "search a valid FlagInt by short name",
			args: args{
				key: "i",
				flags: []flag.Flag{
					flag.FlagStringSlice{Name: "slice", Aliases: []string{"c"}},
					flag.FlagInt{Name: "int", Aliases: []string{"i"}},
					flag.FlagBool{Name: "bool", Aliases: []string{"b"}},
					flag.FlagString{Name: "string", Aliases: []string{"s"}},
				},
			},
			wantIndex: 1,
			wantItem: flag.FlagInt{
				Name:         "int",
				Summary:      "",
				Value:        0,
				Aliases:      []string{"i"},
				EnvVar:       "",
				FlagValue:    flag.Value(""),
				FlagAssigned: false,
			},
			wantFlagShort: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualIndex, actualFlag, actualIsAlias := helpers.FindFlagByKey(tt.args.key, tt.args.flags)
			assert.Equal(t, actualIndex, tt.wantIndex, "Indexes do not match the expected value")
			assert.Equal(t, actualFlag, tt.wantItem, "Flags do not match the expected value")
			assert.Equal(t, actualIsAlias, tt.wantFlagShort, "Flag short status does not match the expected value")
		})
	}
}
