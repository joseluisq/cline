package cline

import (
	"reflect"
	"testing"
)

func TestFlagValue_Bool(t *testing.T) {
	tests := []struct {
		name    string
		v       FlagValue
		want    bool
		wantErr bool
	}{
		{
			name:    "valid bool value as string",
			v:       FlagValue("true"),
			want:    true,
			wantErr: false,
		},
		{
			name:    "valid bool value as int",
			v:       FlagValue("0"),
			want:    false,
			wantErr: false,
		},
		{
			name:    "invalid bool value",
			v:       FlagValue(""),
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Bool()
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagValue.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlagValue.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValue_Int(t *testing.T) {
	tests := []struct {
		name    string
		v       FlagValue
		want    int
		wantErr bool
	}{
		{
			name:    "valid int value as string",
			v:       FlagValue("64"),
			want:    64,
			wantErr: false,
		},
		{
			name:    "invalid bool value",
			v:       FlagValue("z"),
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Int()
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagValue.Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlagValue.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValue_String(t *testing.T) {
	tests := []struct {
		name string
		v    FlagValue
		want string
	}{
		{
			name: "valid string value",
			v:    FlagValue("go"),
			want: "go",
		},
		{
			name: "empty string value",
			v:    FlagValue(""),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("FlagValue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValue_StringSlice(t *testing.T) {
	tests := []struct {
		name string
		v    FlagValue
		want []string
	}{
		{
			name: "valid string value",
			v:    FlagValue("go,rust,zig"),
			want: []string{"go", "rust", "zig"},
		},
		{
			name: "empty string value",
			v:    FlagValue(""),
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.StringSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValue.StringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagMapping_findByKey(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		flagKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   FlagValue
	}{
		{
			name: "find by invalid flag key in null flags",
			fields: fields{
				flags: nil,
			},
			args: args{
				flagKey: "xyz",
			},
			want: FlagValue(""),
		},
		{
			name: "find flag value by unknown key",
			fields: fields{
				flags: []Flag{
					FlagString{
						Name:    "str",
						Summary: "string",
						Value:   "z",
						zflag:   FlagValue("10"),
					},
					FlagInt{
						Name:    "trace",
						Summary: "tracing",
						Value:   10,
						zflag:   FlagValue("10"),
					},
					FlagBool{
						Name:    "verbose",
						Summary: "info details",
						zflag:   FlagValue("true"),
					},
					FlagStringSlice{
						Name:    "coords",
						Summary: "xyz coordinate axis",
						zflag:   FlagValue("x,y,z"),
					},
				},
			},
			args: args{
				flagKey: "xyz",
			},
			want: FlagValue(""),
		},
		{
			name: "find flag bool value by valid key",
			fields: fields{
				flags: []Flag{
					FlagInt{
						Name:    "trace",
						Summary: "tracing",
						zflag:   FlagValue("10"),
					},
					FlagBool{
						Name:          "verbose",
						Summary:       "info details",
						zflag:         FlagValue("true"),
						zflagAssigned: false,
					},
				},
			},
			args: args{
				flagKey: "verbose",
			},
			want: FlagValue("true"),
		},
		{
			name: "find flag int value by valid key",
			fields: fields{
				flags: []Flag{
					FlagBool{
						Name:    "verbose",
						Summary: "info details",
						zflag:   FlagValue("true"),
					},
					FlagInt{
						Name:    "trace",
						Summary: "tracing",
						zflag:   FlagValue("64"),
					},
				},
			},
			args: args{
				flagKey: "trace",
			},
			want: FlagValue("64"),
		},
		{
			name: "find flag string value by valid key",
			fields: fields{
				flags: []Flag{
					FlagBool{
						Name:    "verbose",
						Summary: "info details",
						zflag:   FlagValue("true"),
					},
					FlagString{
						Name:    "trace",
						Summary: "tracing",
						zflag:   FlagValue("xyz"),
					},
				},
			},
			args: args{
				flagKey: "trace",
			},
			want: FlagValue("xyz"),
		},
		{
			name: "find flag string slice value by valid key",
			fields: fields{
				flags: []Flag{
					FlagBool{
						Name:    "verbose",
						Summary: "info details",
						zflag:   FlagValue("true"),
					},
					FlagStringSlice{
						Name:    "output",
						Summary: "format supported",
						zflag:   FlagValue("json,xml,txt"),
					},
				},
			},
			args: args{
				flagKey: "output",
			},
			want: FlagValue("json,xml,txt"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagMapping{
				zFlags: tt.fields.flags,
			}
			if got := fm.findByKey(tt.args.flagKey); got != tt.want {
				t.Errorf("FlagMapping.findByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagMapping_Bool(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		flagName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "find bool flag value by valid key",
			fields: fields{
				flags: []Flag{
					FlagString{
						Name:    "str",
						Summary: "string",
						zflag:   FlagValue("10"),
					},
					FlagInt{
						Name:    "trace",
						Summary: "tracing",
						zflag:   FlagValue("10"),
					},
					FlagBool{
						Name:    "verbose",
						Summary: "info details",
						zflag:   FlagValue("true"),
					},
					FlagStringSlice{
						Name:    "coords",
						Summary: "xyz coordinate axis",
						zflag:   FlagValue("x,y,z"),
					},
				},
			},
			args: args{
				flagName: "verbose",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagMapping{
				zFlags: tt.fields.flags,
			}
			got, err := fm.Bool(tt.args.flagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagMapping.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlagMapping.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagMapping_Int(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		flagName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "find int flag value by valid key",
			fields: fields{
				flags: []Flag{
					FlagString{
						Name:    "str",
						Summary: "string",
						zflag:   FlagValue("10"),
					},
					FlagInt{
						Name:    "level",
						Summary: "level code",
						zflag:   FlagValue("2"),
					},
					FlagBool{
						Name:    "verbose",
						Summary: "info details",
						zflag:   FlagValue("false"),
					},
					FlagStringSlice{
						Name:    "coords",
						Summary: "xyz coordinate axis",
						zflag:   FlagValue("x,y,z"),
					},
				},
			},
			args: args{
				flagName: "level",
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagMapping{
				zFlags: tt.fields.flags,
			}
			got, err := fm.Int(tt.args.flagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagMapping.Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlagMapping.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagMapping_String(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		flagName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "find string flag value by valid key",
			fields: fields{
				flags: []Flag{
					FlagInt{
						Name:    "level",
						Summary: "level code",
						zflag:   FlagValue("2"),
					},
					FlagBool{
						Name:    "verbose",
						Summary: "info details",
						zflag:   FlagValue("false"),
					},
					FlagString{
						Name:    "str",
						Summary: "string",
						zflag:   FlagValue("something"),
					},
					FlagStringSlice{
						Name:    "coords",
						Summary: "xyz coordinate axis",
						zflag:   FlagValue("x,y,z"),
					},
				},
			},
			args: args{
				flagName: "str",
			},
			want: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagMapping{
				zFlags: tt.fields.flags,
			}
			if got := fm.String(tt.args.flagName); got != tt.want {
				t.Errorf("FlagMapping.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagMapping_StringSlice(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		flagName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "find string slice flag value by valid key",
			fields: fields{
				flags: []Flag{
					FlagInt{
						Name:    "level",
						Summary: "level code",
						zflag:   FlagValue("2"),
					},
					FlagBool{
						Name:    "verbose",
						Summary: "info details",
						zflag:   FlagValue("false"),
					},
					FlagStringSlice{
						Name:    "coords",
						Summary: "xyz coordinate axis",
						zflag:   FlagValue("x,y,z"),
					},
					FlagString{
						Name:    "str",
						Summary: "string",
						zflag:   FlagValue("str"),
					},
				},
			},
			args: args{
				flagName: "coords",
			},
			want: []string{"x", "y", "z"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagMapping{
				zFlags: tt.fields.flags,
			}
			if got := fm.StringSlice(tt.args.flagName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagMapping.StringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagMapInput_GetProvidedFlags(t *testing.T) {
	tests := []struct {
		name    string
		vArgs   []string
		want    []FlagProvided
		wantErr bool
	}{
		{
			name:  "valid app provided flags",
			vArgs: []string{"", "--AA", "abc", "-b"},
			want:  []FlagProvided{{Name: "AA", IsAlias: false}, {Name: "BB", IsAlias: true}},
		},
		{
			name:  "valid command provided flags",
			vArgs: []string{"", "info", "--FF", "abc", "-z"},
			want:  []FlagProvided{{Name: "FF", IsAlias: false}, {Name: "ZZ", IsAlias: true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := newApp(func(ctx *AppContext) error {
				if got := ctx.Flags.GetProvidedFlags(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FlagMapping.ProvidedFlags() = %v, want %v", got, tt.want)
				}
				return nil
			}, func(ctx *CmdContext) error {
				if got := ctx.Flags.GetProvidedFlags(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FlagMapping.ProvidedFlags() = %v, want %v", got, tt.want)
				}
				return nil
			})
			if err := app.Run(tt.vArgs); (err != nil) != tt.wantErr {
				t.Errorf("App.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFlagMapping_IsProvidedFlag(t *testing.T) {
	type fields struct {
		zFlags         []Flag
		zFlagsProvided []FlagProvided
	}
	type args struct {
		flagName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "not provided flag",
			fields: fields{
				zFlags: []Flag{
					FlagInt{Name: "int", Aliases: []string{"i"}},
				},
			},
			args: args{
				flagName: "int",
			},
		},
		{
			name: "provided flag",
			fields: fields{
				zFlags: []Flag{
					FlagBool{Name: "bool", Aliases: []string{"b"}},
				},
				zFlagsProvided: []FlagProvided{
					{Name: "bool"},
				},
			},
			args: args{
				flagName: "bool",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagMapping{
				zFlags:         tt.fields.zFlags,
				zFlagsProvided: tt.fields.zFlagsProvided,
			}
			if got := fm.IsProvidedFlag(tt.args.flagName); got != tt.want {
				t.Errorf("FlagMapping.IsProvidedFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagMapping_IsLongProvidedFlag(t *testing.T) {
	type fields struct {
		zFlags         []Flag
		zFlagsProvided []FlagProvided
	}
	type args struct {
		flagName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "not provided flag",
			fields: fields{
				zFlags: []Flag{
					FlagInt{Name: "int", Aliases: []string{"i"}},
				},
			},
			args: args{
				flagName: "int",
			},
		},
		{
			name: "provided long name flag",
			fields: fields{
				zFlags: []Flag{
					FlagBool{Name: "bool", Aliases: []string{"b"}},
				},
				zFlagsProvided: []FlagProvided{
					{Name: "bool"},
				},
			},
			args: args{
				flagName: "bool",
			},
			want: true,
		},
		{
			name: "provided short name flag (alias)",
			fields: fields{
				zFlags: []Flag{
					FlagBool{Name: "int", Aliases: []string{"b"}},
				},
				zFlagsProvided: []FlagProvided{
					{Name: "int", IsAlias: true},
				},
			},
			args: args{
				flagName: "int",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagMapping{
				zFlags:         tt.fields.zFlags,
				zFlagsProvided: tt.fields.zFlagsProvided,
			}
			if got := fm.IsLongProvidedFlag(tt.args.flagName); got != tt.want {
				t.Errorf("FlagMapping.IsLongProvidedFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagMapping_IsShortProvidedFlag(t *testing.T) {
	type fields struct {
		zFlags         []Flag
		zFlagsProvided []FlagProvided
	}
	type args struct {
		flagName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "not provided flag",
			fields: fields{
				zFlags: []Flag{
					FlagStringSlice{Name: "strs", Aliases: []string{"s"}},
				},
			},
			args: args{
				flagName: "strs",
			},
		},
		{
			name: "provided long name flag",
			fields: fields{
				zFlags: []Flag{
					FlagString{Name: "str", Aliases: []string{"s"}},
				},
				zFlagsProvided: []FlagProvided{
					{Name: "str"},
				},
			},
			args: args{
				flagName: "str",
			},
		},
		{
			name: "provided short name flag (alias)",
			fields: fields{
				zFlags: []Flag{
					FlagBool{Name: "int", Aliases: []string{"i"}},
				},
				zFlagsProvided: []FlagProvided{
					{Name: "int", IsAlias: true},
				},
			},
			args: args{
				flagName: "int",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagMapping{
				zFlags:         tt.fields.zFlags,
				zFlagsProvided: tt.fields.zFlagsProvided,
			}
			if got := fm.IsShortProvidedFlag(tt.args.flagName); got != tt.want {
				t.Errorf("FlagMapping.IsShortProvidedFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
