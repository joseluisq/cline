package flag_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joseluisq/cline/flag"
)

func TestAnyValue_ToBool(t *testing.T) {
	tests := []struct {
		name    string
		v       flag.Value
		want    bool
		wantErr bool
	}{
		{
			name:    "invalid value parsing",
			v:       flag.Value(""),
			wantErr: true,
		},
		{
			name: "valid value parsing",
			v:    flag.Value("1"),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actualVal, actualErr := tt.v.ToBool(); tt.wantErr {
				assert.Error(t, actualErr, "Expected an error but got none")
			} else {
				assert.NoError(t, actualErr, "Expected no error but got one")
				assert.Equal(t, actualVal, tt.want, "Bool value does not match the expected one")
			}
		})
	}
}

func TestAnyValue_ToInt(t *testing.T) {
	tests := []struct {
		name    string
		v       flag.Value
		want    int
		wantErr bool
	}{
		{
			name:    "invalid value parsing",
			v:       flag.Value("10.14"),
			wantErr: true,
		},
		{
			name: "valid value parsing",
			v:    flag.Value("10"),
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actualVal, actualErr := tt.v.ToInt(); tt.wantErr {
				assert.Error(t, actualErr, "Expected an error but got none")
			} else {
				assert.NoError(t, actualErr, "Expected no error but got one")
				assert.Equal(t, actualVal, tt.want, "Int value does not match the expected one")
			}
		})
	}
}

func TestAnyValue_ToString(t *testing.T) {
	tests := []struct {
		name string
		v    flag.Value
		want string
	}{
		{
			name: "empty value parsing",
			want: "",
		},
		{
			name: "no empty value parsing",
			v:    flag.Value("abc"),
			want: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualVal := tt.v.ToString()
			assert.Equal(t, actualVal, tt.want, "String value does not match the expected one")
		})
	}
}

func TestAnyValue_ToStringSlice(t *testing.T) {
	tests := []struct {
		name string
		v    flag.Value
		want []string
	}{
		{
			name: "empty value",
			v:    flag.Value(""),
			want: []string{""},
		},
		{
			name: "list of values",
			v:    flag.Value("a,b,c,d"),
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "list of values with spaces",
			v:    flag.Value("abc,   ,"),
			want: []string{"abc", "", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualValues := tt.v.ToStringSlice()
			assert.Equal(t, actualValues, tt.want, "String value does not match the expected one")
		})
	}
}

func TestFlagBoolValue_Value(t *testing.T) {
	type fields struct {
		flag flag.FlagBool
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "empty flag value",
			fields: fields{
				flag: flag.FlagBool{},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "non empty value",
			fields: fields{
				flag: flag.FlagBool{
					Name:      "version",
					Value:     true,
					FlagValue: flag.Value("true"),
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := flag.ValueBool{
				Flag: tt.fields.flag,
			}
			if actualVal, actualErr := v.Value(); tt.wantErr {
				assert.Error(t, actualErr, "Expected an error but got none")
			} else {
				assert.NoError(t, actualErr, "Expected no error but got one")
				assert.Equal(t, actualVal, tt.want, "Bool value does not match the expected one")
			}
		})
	}
}

func TestFlagBoolValue_IsProvided(t *testing.T) {
	type fields struct {
		flag flag.FlagBool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided flag",
			fields: fields{
				flag: flag.FlagBool{},
			},
		},
		{
			name: "provided flag",
			fields: fields{
				flag: flag.FlagBool{
					Name:         "version",
					Value:        true,
					FlagValue:    flag.Value("true"),
					FlagProvided: true,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := flag.ValueBool{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvided()
			assert.Equal(t, actualVal, tt.want, "IsProvided value does not match the expected one")
		})
	}
}

func TestFlagBoolValue_IsProvidedShort(t *testing.T) {
	type fields struct {
		flag flag.FlagBool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided short flag",
			fields: fields{
				flag: flag.FlagBool{},
			},
		},
		{
			name: "provided short flag",
			fields: fields{
				flag: flag.FlagBool{
					Name:                "short",
					Value:               true,
					FlagValue:           flag.Value("true"),
					FlagProvided:        true,
					FlagProvidedAsAlias: true,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := flag.ValueBool{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedShort()
			assert.Equal(t, actualVal, tt.want, "IsProvidedShort value does not match the expected one")
		})
	}
}

func TestFlagBoolValue_IsProvidedLong(t *testing.T) {
	type fields struct {
		flag flag.FlagBool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided long flag",
			fields: fields{
				flag: flag.FlagBool{},
			},
		},
		{
			name: "provided long flag",
			fields: fields{
				flag: flag.FlagBool{
					Name:                "long",
					Value:               true,
					FlagValue:           flag.Value("true"),
					FlagProvided:        true,
					FlagProvidedAsAlias: false,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := flag.ValueBool{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedLong()
			assert.Equal(t, actualVal, tt.want, "IsProvidedLong value does not match the expected one")
		})
	}
}

func TestFlagBoolValue_GetFlagType(t *testing.T) {
	type fields struct {
		flag flag.FlagBool
	}
	tests := []struct {
		name   string
		fields fields
		want   flag.FlagBool
	}{
		{
			name: "get flag type",
			fields: fields{
				flag: flag.FlagBool{},
			},
			want: flag.FlagBool{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := flag.ValueBool{
				Flag: tt.fields.flag,
			}
			actualVal := v.GetFlagType()
			assert.Equal(t, actualVal, tt.want, "GetFlagType value does not match the expected one")
		})
	}
}

func TestFlagStringValue_Value(t *testing.T) {
	type fields struct {
		flag flag.FlagString
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty flag value",
			fields: fields{
				flag: flag.FlagString{},
			},
			want: "",
		},
		{
			name: "non empty value",
			fields: fields{
				flag: flag.FlagString{
					Name:      "abc",
					FlagValue: flag.Value("xyz"),
				},
			},
			want: "xyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueString{
				Flag: tt.fields.flag,
			}
			actualVal := v.Value()
			assert.Equal(t, actualVal, tt.want, "String value does not match the expected one")
		})
	}
}

func TestFlagStringValue_IsProvided(t *testing.T) {
	type fields struct {
		flag flag.FlagString
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided flag",
			fields: fields{
				flag: flag.FlagString{},
			},
		},
		{
			name: "provided flag",
			fields: fields{
				flag: flag.FlagString{
					Name:         "alpha",
					FlagProvided: true,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueString{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvided()
			assert.Equal(t, actualVal, tt.want, "IsProvided value does not match the expected one")
		})
	}
}

func TestFlagStringValue_IsProvidedShort(t *testing.T) {
	type fields struct {
		flag flag.FlagString
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided flag",
			fields: fields{
				flag: flag.FlagString{},
			},
		},
		{
			name: "provided flag",
			fields: fields{
				flag: flag.FlagString{
					Name:                "alpha",
					FlagProvided:        true,
					FlagProvidedAsAlias: true,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueString{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedShort()
			assert.Equal(t, actualVal, tt.want, "IsProvidedShort value does not match the expected one")
		})
	}
}

func TestFlagStringValue_IsProvidedLong(t *testing.T) {
	type fields struct {
		flag flag.FlagString
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided long flag",
			fields: fields{
				flag: flag.FlagString{},
			},
		},
		{
			name: "provided long flag",
			fields: fields{
				flag: flag.FlagString{
					Name:                "long",
					Value:               "cde",
					FlagValue:           flag.Value("cde"),
					FlagProvided:        true,
					FlagProvidedAsAlias: false,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueString{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedLong()
			assert.Equal(t, actualVal, tt.want, "IsProvidedLong value does not match the expected one")
		})
	}
}

func TestFlagStringValue_GetFlagType(t *testing.T) {
	type fields struct {
		flag flag.FlagString
	}
	tests := []struct {
		name   string
		fields fields
		want   flag.FlagString
	}{
		{
			name: "get flag type",
			fields: fields{
				flag: flag.FlagString{},
			},
			want: flag.FlagString{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueString{
				Flag: tt.fields.flag,
			}
			actualVal := v.GetFlagType()
			assert.Equal(t, actualVal, tt.want, "GetFlagType value does not match the expected one")
		})
	}
}

func TestFlagStringSliceValue_Value(t *testing.T) {
	type fields struct {
		flag flag.FlagStringSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "empty flag value",
			fields: fields{
				flag: flag.FlagStringSlice{},
			},
			want: []string{""},
		},
		{
			name: "non empty value",
			fields: fields{
				flag: flag.FlagStringSlice{
					Name:      "abc",
					FlagValue: flag.Value("xyz, abc"),
				},
			},
			want: []string{"xyz", "abc"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueStringSlice{
				Flag: tt.fields.flag,
			}
			actualVal := v.Value()
			assert.Equal(t, actualVal, tt.want, "String slice value does not match the expected one")
		})
	}
}

func TestFlagStringSliceValue_IsProvided(t *testing.T) {
	type fields struct {
		flag flag.FlagStringSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided flag",
			fields: fields{
				flag: flag.FlagStringSlice{},
			},
		},
		{
			name: "provided flag",
			fields: fields{
				flag: flag.FlagStringSlice{
					Name:         "alpha",
					FlagProvided: true,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueStringSlice{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvided()
			assert.Equal(t, actualVal, tt.want, "IsProvided value does not match the expected one")
		})
	}
}

func TestFlagStringSliceValue_IsProvidedShort(t *testing.T) {
	type fields struct {
		flag flag.FlagStringSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided short flag",
			fields: fields{
				flag: flag.FlagStringSlice{},
			},
		},
		{
			name: "provided short flag false",
			fields: fields{
				flag: flag.FlagStringSlice{
					Name:         "alpha",
					FlagProvided: true,
				},
			},
		},
		{
			name: "provided short flag true",
			fields: fields{
				flag: flag.FlagStringSlice{
					Name:                "alpha",
					FlagProvided:        true,
					FlagProvidedAsAlias: true,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueStringSlice{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedShort()
			assert.Equal(t, actualVal, tt.want, "IsProvidedShort value does not match the expected one")
		})
	}
}

func TestFlagStringSliceValue_IsProvidedLong(t *testing.T) {
	type fields struct {
		flag flag.FlagStringSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided long flag",
			fields: fields{
				flag: flag.FlagStringSlice{},
			},
		},
		{
			name: "provided long flag true",
			fields: fields{
				flag: flag.FlagStringSlice{
					Name:         "alpha",
					FlagProvided: true,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueStringSlice{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedLong()
			assert.Equal(t, actualVal, tt.want, "IsProvidedLong value does not match the expected one")
		})
	}
}

func TestFlagStringSliceValue_GetFlagType(t *testing.T) {
	type fields struct {
		flag flag.FlagStringSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   flag.FlagStringSlice
	}{
		{
			name: "get flag type",
			fields: fields{
				flag: flag.FlagStringSlice{},
			},
			want: flag.FlagStringSlice{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueStringSlice{
				Flag: tt.fields.flag,
			}
			actualVal := v.GetFlagType()
			assert.Equal(t, actualVal, tt.want, "GetFlagType value does not match the expected one")
		})
	}
}

func TestFlagValues_findByKey(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFlag flag.Flag
	}{
		{
			name: "lookup on an unsupported type flag list",
			fields: fields{
				flags: []flag.Flag{
					map[string]int{"a": 1, "b": 2},
				},
			},
			args: args{
				longFlagName: "xyz",
			},
			wantFlag: nil,
		},
		{
			name: "lookup using empty flag key",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{},
					flag.FlagBool{},
					flag.FlagString{},
					flag.FlagStringSlice{},
				},
			},
			args: args{
				longFlagName: "",
			},
			wantFlag: nil,
		},
		{
			name: "lookup using valid int flag key",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int"},
					flag.FlagBool{Name: "k-bool"},
					flag.FlagString{Name: "k-string"},
					flag.FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				longFlagName: "k-int",
			},
			wantFlag: flag.FlagInt{Name: "k-int"},
		},
		{
			name: "lookup using valid bool flag key",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int"},
					flag.FlagBool{Name: "k-bool"},
					flag.FlagString{Name: "k-string"},
					flag.FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				longFlagName: "k-bool",
			},
			wantFlag: flag.FlagBool{Name: "k-bool"},
		},
		{
			name: "lookup using valid string flag key",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int"},
					flag.FlagBool{Name: "k-bool"},
					flag.FlagString{Name: "k-string"},
					flag.FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				longFlagName: "k-string",
			},
			wantFlag: flag.FlagString{Name: "k-string"},
		},
		{
			name: "lookup using valid string-slice flag key",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int"},
					flag.FlagBool{Name: "k-bool"},
					flag.FlagString{Name: "k-string"},
					flag.FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				longFlagName: "k-string-slice",
			},
			wantFlag: flag.FlagStringSlice{Name: "k-string-slice"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			actualFlag := v.FindByKey(tt.args.longFlagName)
			assert.Equal(t, actualFlag, tt.wantFlag, "FindByKey value does not match the expected one")
		})
	}
}

func TestFlagValues_getProvidedFlags(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	type args struct {
		providedOnly bool
		aliasOnly    bool
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantFlags []flag.Flag
	}{
		{
			name: "return all flags",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int"},
					flag.FlagBool{Name: "k-bool"},
					flag.FlagString{Name: "k-string"},
					flag.FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{},
			wantFlags: []flag.Flag{
				flag.FlagInt{Name: "k-int"},
				flag.FlagBool{Name: "k-bool"},
				flag.FlagString{Name: "k-string"},
				flag.FlagStringSlice{Name: "k-string-slice"},
			},
		},
		{
			name: "return empty list of flags",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int"},
					flag.FlagBool{Name: "k-bool"},
					flag.FlagString{Name: "k-string"},
					flag.FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				providedOnly: true,
			},
			wantFlags: nil,
		},
		{
			name: "return provided flags",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true},
					flag.FlagBool{Name: "k-bool"},
					flag.FlagString{Name: "k-string", FlagProvided: true},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true},
				},
			},
			args: args{
				providedOnly: true,
			},
			wantFlags: []flag.Flag{
				flag.FlagInt{Name: "k-int", FlagProvided: true},
				flag.FlagString{Name: "k-string", FlagProvided: true},
				flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true},
			},
		},
		{
			name: "return provided short flags",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true},
					flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagProvidedAsAlias: true},
					flag.FlagString{Name: "k-string", FlagProvided: true, FlagProvidedAsAlias: true},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true},
				},
			},
			args: args{
				aliasOnly: true,
			},
			wantFlags: []flag.Flag{
				flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagProvidedAsAlias: true},
				flag.FlagString{Name: "k-string", FlagProvided: true, FlagProvidedAsAlias: true},
			},
		},
		{
			name: "return provided flags only",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true},
					flag.FlagBool{Name: "k-bool", FlagProvided: true},
					flag.FlagString{Name: "k-string", FlagProvided: true},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true},
				},
			},
			args: args{
				providedOnly: true,
			},
			wantFlags: []flag.Flag{
				flag.FlagInt{Name: "k-int", FlagProvided: true},
				flag.FlagBool{Name: "k-bool", FlagProvided: true},
				flag.FlagString{Name: "k-string", FlagProvided: true},
				flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true},
			},
		},
		{
			name: "return provided short flags only",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagProvidedAsAlias: true},
					flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagProvidedAsAlias: true},
					flag.FlagString{Name: "k-string", FlagProvided: true, FlagProvidedAsAlias: true},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagProvidedAsAlias: true},
				},
			},
			args: args{
				aliasOnly: true,
			},
			wantFlags: []flag.Flag{
				flag.FlagInt{Name: "k-int", FlagProvided: true, FlagProvidedAsAlias: true},
				flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagProvidedAsAlias: true},
				flag.FlagString{Name: "k-string", FlagProvided: true, FlagProvidedAsAlias: true},
				flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagProvidedAsAlias: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			actual := v.GetProvidedFlags(tt.args.providedOnly, tt.args.aliasOnly)
			assert.Equal(t, actual, tt.wantFlags, "GetProvidedFlags value does not match the expected one")
		})
	}
}

func TestFlagValues_GetProvided(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	tests := []struct {
		name   string
		fields fields
		want   []flag.Flag
	}{
		{
			name: "get provided nil flags",
			fields: fields{
				flags: []flag.Flag{},
			},
			want: nil,
		},
		{
			name: "get provided flag list",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagProvidedAsAlias: true},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagProvidedAsAlias: true},
				},
			},
			want: []flag.Flag{
				flag.FlagInt{Name: "k-int", FlagProvided: true, FlagProvidedAsAlias: true},
				flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagProvidedAsAlias: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			actual := v.GetProvided()
			assert.Equal(t, actual, tt.want, "GetProvided value does not match the expected one")
		})
	}
}

func TestFlagValues_GetProvidedLong(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	tests := []struct {
		name   string
		fields fields
		want   []flag.Flag
	}{
		{
			name: "get provided empty flags",
			fields: fields{
				flags: []flag.Flag{},
			},
			want: []flag.Flag{},
		},
		{
			name: "get provided flag list",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true},
				},
			},
			want: []flag.Flag{
				flag.FlagInt{Name: "k-int", FlagProvided: true},
				flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			actual := v.GetProvidedLong()
			assert.Equal(t, actual, tt.want, "GetProvidedLong value does not match the expected one")
		})
	}
}

func TestFlagValues_GetProvidedShort(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	tests := []struct {
		name   string
		fields fields
		want   []flag.Flag
	}{
		{
			name: "get provided empty short flags",
			fields: fields{
				flags: []flag.Flag{},
			},
			want: nil,
		},
		{
			name: "get provided short flag list",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagProvidedAsAlias: true},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagProvidedAsAlias: true},
				},
			},
			want: []flag.Flag{
				flag.FlagInt{Name: "k-int", FlagProvided: true, FlagProvidedAsAlias: true},
				flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagProvidedAsAlias: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			actual := v.GetProvidedShort()
			assert.Equal(t, actual, tt.want, "GetProvidedShort value does not match the expected one")
		})
	}
}

func TestFlagValues_Any(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   flag.Value
	}{
		{
			name: "lookup using empty bool flag list",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int"},
					flag.FlagBool{Name: "k-bool"},
					flag.FlagString{Name: "k-string"},
					flag.FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				longFlagName: "k-bool",
			},
			want: "",
		},
		{
			name: "lookup using empty string flag list",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int"},
					flag.FlagBool{Name: "k-bool"},
					flag.FlagString{Name: "k-string"},
					flag.FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				longFlagName: "k-string",
			},
			want: "",
		},
		{
			name: "lookup using empty string-slice flag list",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int"},
					flag.FlagBool{Name: "k-bool"},
					flag.FlagString{Name: "k-string"},
					flag.FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				longFlagName: "k-string-slice",
			},
			want: "",
		},
		{
			name: "lookup using valid int flag key",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagValue: "5"},
					flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagValue: "true"},
					flag.FlagString{Name: "k-string", FlagProvided: true, FlagValue: "string_val"},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagValue: "a,b,c"},
				},
			},
			args: args{
				longFlagName: "k-int",
			},
			want: flag.Value("5"),
		},
		{
			name: "lookup using empty string-slice flag list",
			fields: fields{
				flags: []flag.Flag{
					map[string]int{"a": 1, "b": 2},
				},
			},
			args: args{
				longFlagName: "k-any",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			actual := v.Value(tt.args.longFlagName)
			assert.Equal(t, actual, tt.want, "Value does not match the expected one")
		})
	}
}

func TestFlagValues_Bool(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *flag.ValueBool
		wantErr bool
	}{
		{
			name: "get invalid bool value",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagValue: "5"},
					flag.FlagBool{Name: "k-bool", FlagProvided: false, FlagValue: "true"},
					flag.FlagString{Name: "k-string", FlagProvided: true, FlagValue: "string_val"},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagValue: "a,b,c"},
				},
			},
			args: args{
				longFlagName: "some",
			},
			wantErr: true,
		},
		{
			name: "get valid bool value",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagValue: "5"},
					flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagValue: "true"},
					flag.FlagString{Name: "k-string", FlagProvided: true, FlagValue: "string_val"},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagValue: "a,b,c"},
				},
			},
			args: args{
				longFlagName: "k-bool",
			},
			want: &flag.ValueBool{
				Flag: flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagValue: "true"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			if actual, err := v.Bool(tt.args.longFlagName); tt.wantErr {
				assert.Error(t, err, "expected error but got none")
			} else {
				assert.NoError(t, err, "did not expect error but got one")
				assert.Equal(t, actual, tt.want, "Bool value do not match")
			}
		})
	}
}

func TestFlagValues_Int(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *flag.ValueInt
		wantErr bool
	}{
		{
			name: "get invalid int value",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagValue: "5"},
					flag.FlagBool{Name: "k-bool", FlagProvided: false, FlagValue: "true"},
					flag.FlagString{Name: "k-string", FlagProvided: true, FlagValue: "string_val"},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagValue: "a,b,c"},
				},
			},
			args: args{
				longFlagName: "some",
			},
			wantErr: true,
		},
		{
			name: "get valid int value",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagValue: "256"},
					flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagValue: "true"},
					flag.FlagString{Name: "k-string", FlagProvided: true, FlagValue: "string_val"},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagValue: "a,b,c"},
				},
			},
			args: args{
				longFlagName: "k-int",
			},
			want: &flag.ValueInt{
				Flag: flag.FlagInt{Name: "k-int", FlagProvided: true, FlagValue: "256"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			if actual, err := v.Int(tt.args.longFlagName); tt.wantErr {
				assert.Error(t, err, "expected error but got none")
			} else {
				assert.NoError(t, err, "did not expect error but got one")
				assert.Equal(t, actual, tt.want, "Int value do not match")
			}
		})
	}
}

func TestFlagValues_String(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected *flag.ValueString
		wantErr  bool
	}{
		{
			name: "get invalid string value",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagValue: "5"},
					flag.FlagBool{Name: "k-bool", FlagProvided: false, FlagValue: "true"},
					flag.FlagString{Name: "k-string", FlagProvided: true, FlagValue: "string_val"},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagValue: "a,b,c"},
				},
			},
			args: args{
				longFlagName: "some",
			},
			wantErr: true,
		},
		{
			name: "get valid string value",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagValue: "256"},
					flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagValue: "true"},
					flag.FlagString{Name: "k-string", FlagProvided: true, FlagValue: "string_val"},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagValue: "a,b,c"},
				},
			},
			args: args{
				longFlagName: "k-string",
			},
			expected: &flag.ValueString{
				Flag: flag.FlagString{Name: "k-string", FlagProvided: true, FlagValue: "string_val"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			if actual, err := v.String(tt.args.longFlagName); tt.wantErr {
				assert.Error(t, err, "expected error but got none")
			} else {
				assert.NoError(t, err, "did not expect error but got one")
				assert.Equal(t, actual, tt.expected, "string value do not match")
			}
		})
	}
}

func TestFlagValues_StringSlice(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected *flag.ValueStringSlice
		wantErr  bool
	}{
		{
			name: "get invalid string-slice value",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagValue: "5"},
					flag.FlagBool{Name: "k-bool", FlagProvided: false, FlagValue: "true"},
					flag.FlagString{Name: "k-string", FlagProvided: true, FlagValue: "string_val"},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagValue: "1,2,3"},
				},
			},
			args: args{
				longFlagName: "some",
			},
			wantErr: true,
		},
		{
			name: "get valid string-slice value",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagValue: "256"},
					flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagValue: "true"},
					flag.FlagString{Name: "k-string", FlagProvided: true, FlagValue: "string_val"},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagValue: "1,2,3"},
				},
			},
			args: args{
				longFlagName: "k-string-slice",
			},
			expected: &flag.ValueStringSlice{
				Flag: flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagValue: "1,2,3"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			if actual, err := v.StringSlice(tt.args.longFlagName); tt.wantErr {
				assert.Error(t, err, "expected error but got none")
			} else {
				assert.NoError(t, err, "did not expect error but got one")
				assert.Equal(t, actual, tt.expected, "string slice value do not match")
			}
		})
	}
}

func TestFlagIntValue_Value(t *testing.T) {
	type fields struct {
		flag flag.FlagInt
	}
	tests := []struct {
		name     string
		fields   fields
		expected int
		wantErr  bool
	}{
		{
			name: "no provided value flag",
			fields: fields{
				flag: flag.FlagInt{},
			},
			wantErr: true,
		},
		{
			name: "provided value flag",
			fields: fields{
				flag: flag.FlagInt{
					Name:         "short",
					Value:        7,
					FlagValue:    flag.Value("7"),
					FlagProvided: true,
				},
			},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueInt{
				Flag: tt.fields.flag,
			}

			if actual, err := v.Value(); tt.wantErr {
				assert.Error(t, err, "expected error but got none")
			} else {
				assert.NoError(t, err, "did not expect error but got one")
				assert.Equal(t, actual, tt.expected, "Value does not match the expected one")
			}
		})
	}
}

func TestFlagIntValue_IsProvided(t *testing.T) {
	type fields struct {
		flag flag.FlagInt
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "no provided flag",
			fields: fields{
				flag: flag.FlagInt{},
			},
		},
		{
			name: "provided flag",
			fields: fields{
				flag: flag.FlagInt{
					Name:         "provided",
					Value:        32,
					FlagValue:    flag.Value("32"),
					FlagProvided: true,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueInt{
				Flag: tt.fields.flag,
			}
			actual := v.IsProvided()
			assert.Equal(t, actual, tt.expected, "IsProvided value does not match the expected one")
		})
	}
}

func TestFlagIntValue_IsProvidedShort(t *testing.T) {
	type fields struct {
		flag flag.FlagInt
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "no provided short flag",
			fields: fields{
				flag: flag.FlagInt{},
			},
		},
		{
			name: "provided short flag",
			fields: fields{
				flag: flag.FlagInt{
					Name:                "provided",
					FlagProvided:        true,
					FlagProvidedAsAlias: true,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueInt{
				Flag: tt.fields.flag,
			}
			actual := v.IsProvidedShort()
			assert.Equal(t, actual, tt.expected, "IsProvidedShort value does not match the expected one")
		})
	}
}

func TestFlagIntValue_IsProvidedLong(t *testing.T) {
	type fields struct {
		flag flag.FlagInt
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "no provided long flag",
			fields: fields{
				flag: flag.FlagInt{},
			},
		},
		{
			name: "provided long flag",
			fields: fields{
				flag: flag.FlagInt{
					Name:                "provided",
					FlagProvided:        true,
					FlagProvidedAsAlias: false,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueInt{
				Flag: tt.fields.flag,
			}
			actual := v.IsProvidedLong()
			assert.Equal(t, actual, tt.expected, "IsProvidedLong value does not match the expected one")
		})
	}
}

func TestFlagIntValue_GetFlagType(t *testing.T) {
	type fields struct {
		flag flag.FlagInt
	}
	tests := []struct {
		name     string
		fields   fields
		expected flag.FlagInt
	}{
		{
			name: "get flag type",
			fields: fields{
				flag: flag.FlagInt{},
			},
			expected: flag.FlagInt{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueInt{
				Flag: tt.fields.flag,
			}
			actual := v.GetFlagType()
			assert.Equal(t, actual, tt.expected, "GetFlagType value does not match the expected one")
		})
	}
}
