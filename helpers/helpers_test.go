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
			want: nil,
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
		},
		{
			name: "should return error for invalid string command name",
			args: args{
				commands: []app.Cmd{{Name: "ñame"}},
			},
			wantErr: true,
		},
		{
			name: "should return error for invalid bool command name",
			args: args{
				commands: []app.Cmd{{Name: "~name"}},
			},
			wantErr: true,
		},
		{
			name: "should return error for invalid string slice command name",
			args: args{
				commands: []app.Cmd{{Name: "name-\x00"}},
			},
			wantErr: true,
		},
		{
			name: "should return error for invalid int command name",
			args: args{
				commands: []app.Cmd{{Name: "name-🚀"}},
			},
			wantErr: true,
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
			want: nil,
		},
		{
			name: "should return error for invalid FlagString name",
			args: args{
				flags: []flag.Flag{
					flag.FlagString{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error for invalid FlagBool name",
			args: args{
				flags: []flag.Flag{
					flag.FlagBool{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error for invalid FlagInt name",
			args: args{
				flags: []flag.Flag{
					flag.FlagInt{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error for invalid FlagStringSlice name",
			args: args{
				flags: []flag.Flag{
					flag.FlagStringSlice{Name: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error for invalid flag type value",
			args: args{
				flags: []flag.Flag{
					struct{ ok bool }{false},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error for nil flag in slice",
			args: args{
				flags: []flag.Flag{
					nil,
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
		},
		{
			name: "hould return error for invalid flag names",
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
		},
		{
			name: "should return error for invalid string flag name",
			args: args{
				flags: []flag.Flag{
					flag.FlagString{Name: "ñame"},
				},
			},
			wantErr: true,
		},
		{
			name: "should return error for invalid bool flag name",
			args: args{
				flags: []flag.Flag{
					flag.FlagBool{Name: "~name"},
				},
			},
			wantErr: true,
		},
		{
			name: "should return error for invalid string slice flag name",
			args: args{
				flags: []flag.Flag{
					flag.FlagStringSlice{Name: "name-\x00"},
				},
			},
			wantErr: true,
		},
		{
			name: "should return error for invalid int flag name",
			args: args{
				flags: []flag.Flag{
					flag.FlagInt{Name: "name-🚀"},
				},
			},
			wantErr: true,
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

func Test_BuildFlagMap(t *testing.T) {
	tests := []struct {
		name  string
		flags []flag.Flag
		want  map[string]helpers.FlagInfo
	}{
		{
			name:  "should handle an empty flag slice",
			flags: []flag.Flag{},
			want:  map[string]helpers.FlagInfo{},
		},
		{
			name: "should create a map with names and aliases",
			flags: []flag.Flag{
				flag.FlagBool{Name: "verbose", Aliases: []string{"v"}},
				flag.FlagString{Name: "file", Aliases: []string{"f"}},
			},
			want: map[string]helpers.FlagInfo{
				"verbose": {Flag: flag.FlagBool{Name: "verbose", Aliases: []string{"v"}}, Index: 0},
				"v":       {Flag: flag.FlagBool{Name: "verbose", Aliases: []string{"v"}}, Index: 0},
				"file":    {Flag: flag.FlagString{Name: "file", Aliases: []string{"f"}}, Index: 1},
				"f":       {Flag: flag.FlagString{Name: "file", Aliases: []string{"f"}}, Index: 1},
			},
		},
		{
			name: "should handle mixed flag types",
			flags: []flag.Flag{
				flag.FlagInt{Name: "count", Aliases: []string{"c"}},
				flag.FlagStringSlice{Name: "items", Aliases: []string{"i"}},
			},
			want: map[string]helpers.FlagInfo{
				"count": {Flag: flag.FlagInt{Name: "count", Aliases: []string{"i"}}, Index: 0},
				"c":     {Flag: flag.FlagStringSlice{Name: "count", Aliases: []string{"c"}}, Index: 0},
				"items": {Flag: flag.FlagStringSlice{Name: "items", Aliases: []string{"i"}}, Index: 1},
				"i":     {Flag: flag.FlagStringSlice{Name: "items", Aliases: []string{"i"}}, Index: 1},
			},
		},
		{
			name: "should handle alias collisions by taking the last one",
			flags: []flag.Flag{
				flag.FlagString{Name: "output"},
				flag.FlagBool{Name: "override", Aliases: []string{"output"}},
			},
			want: map[string]helpers.FlagInfo{
				"output":   {Flag: flag.FlagBool{Name: "override", Aliases: []string{"output"}}, Index: 1},
				"override": {Flag: flag.FlagBool{Name: "override", Aliases: []string{"output"}}, Index: 1},
			},
		},
		{
			name: "should skip nil flags in the slice",
			flags: []flag.Flag{
				flag.FlagString{Name: "file"},
				nil,
				flag.FlagBool{Name: "verbose"},
			},
			want: map[string]helpers.FlagInfo{
				"file":    {Flag: flag.FlagString{Name: "file"}, Index: 0},
				"verbose": {Flag: flag.FlagBool{Name: "verbose"}, Index: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := helpers.BuildFlagMap(tt.flags)
			assert.Equal(t, len(tt.want), len(got), "map length should be equal")
			for key, wantInfo := range tt.want {
				gotInfo, ok := got[key]
				assert.True(t, ok, "key %s should exist in the map", key)
				assert.Equal(t, wantInfo.Index, gotInfo.Index, "index for key %s should match", key)
				assert.ObjectsAreEqual(wantInfo.Flag, gotInfo.Flag)
			}
		})
	}
}

func TestIsValidToken(t *testing.T) {
	tests := []struct {
		name      string
		token     string
		tokenType string
		wantErr   bool
	}{
		{
			name:      "should return no error for valid alphanumeric token",
			token:     "validToken-123",
			tokenType: "flag",
		},
		{
			name:      "should return error for invalid token with special chars",
			token:     "invalid-token_123.",
			tokenType: "command",
			wantErr:   true,
		},
		{
			name:      "should return error for invalid flag token with hyphen at start",
			token:     "-invalidToken",
			tokenType: "flag",
			wantErr:   true,
		},
		{
			name:      "should return error for invalid flag token with hyphen at end",
			token:     "invalidToken-",
			tokenType: "flag",
			wantErr:   true,
		},
		{
			name:      "should return error for invalid flag token with consecutive hyphens",
			token:     "invalid--token-name",
			tokenType: "flag",
			wantErr:   true,
		},
		{
			name:      "should return error for invalid flag token with digits at start",
			token:     "123-invalid",
			tokenType: "flag",
			wantErr:   true,
		},
		{
			name:      "should return error for invalid flag token with space",
			token:     "invalid token",
			tokenType: "flag",
			wantErr:   true,
		},
		{
			name:      "should return error for invalid command token with tab",
			token:     "invalid\ttoken",
			tokenType: "command",
			wantErr:   true,
		},
		{
			name:      "should return error for invalid flag token with unicode",
			token:     "invalidñtoken",
			tokenType: "flag",
			wantErr:   true,
		},
		{
			name:      "should return error for invalid command token emoji",
			token:     "invalid🚀token",
			tokenType: "command",
			wantErr:   true,
		},
		{
			name:      "should return error for empty flag token",
			token:     "",
			tokenType: "flag",
		},
		{
			name:      "should return error for invalid flag token with null byte",
			token:     "invalid\x00token",
			tokenType: "flag",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actualErr := helpers.IsValidToken(tt.token, tt.tokenType); tt.wantErr {
				assert.Error(t, actualErr, "IsValidToken() expected to return an error but got nil")
				assert.Equal(t, actualErr.Error(), "error: "+tt.tokenType+" '"+tt.token+"' contains invalid characters")
			} else {
				assert.NoError(t, actualErr, "IsValidToken() returned an unexpected error")
			}
		})
	}
}
