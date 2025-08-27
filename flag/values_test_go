package cline

import (
	"reflect"
	"testing"
)

func TestAnyValue_ToBool(t *testing.T) {
	tests := []struct {
		name    string
		v       AnyValue
		want    bool
		wantErr bool
	}{
		{
			name:    "invalid value parsing",
			v:       AnyValue(""),
			wantErr: true,
		},
		{
			name: "valid value parsing",
			v:    AnyValue("1"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.ToBool()
			if (err != nil) != tt.wantErr {
				t.Errorf("AnyValue.ToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AnyValue.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyValue_ToInt(t *testing.T) {
	tests := []struct {
		name    string
		v       AnyValue
		want    int
		wantErr bool
	}{
		{
			name:    "invalid value parsing",
			v:       AnyValue("10.14"),
			wantErr: true,
		},
		{
			name: "valid value parsing",
			v:    AnyValue("10"),
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.ToInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("AnyValue.ToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AnyValue.ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyValue_ToString(t *testing.T) {
	tests := []struct {
		name string
		v    AnyValue
		want string
	}{
		{
			name: "empty value parsing",
			want: "",
		},
		{
			name: "no empty value parsing",
			v:    AnyValue("abc"),
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.ToString(); got != tt.want {
				t.Errorf("AnyValue.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyValue_ToStringSlice(t *testing.T) {
	tests := []struct {
		name string
		v    AnyValue
		want []string
	}{
		{
			name: "empty value",
			v:    AnyValue(""),
			want: []string{""},
		},
		{
			name: "list of values",
			v:    AnyValue("a,b,c,d"),
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "list of values with spaces",
			v:    AnyValue("abc,   ,"),
			want: []string{"abc", "", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.ToStringSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnyValue.ToStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagBoolValue_Value(t *testing.T) {
	type fields struct {
		flag FlagBool
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
				flag: FlagBool{},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "non empty value",
			fields: fields{
				flag: FlagBool{
					Name:      "version",
					Value:     true,
					flagValue: AnyValue("true"),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagBoolValue{
				flag: tt.fields.flag,
			}
			got, err := v.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagBoolValue.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlagBoolValue.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagBoolValue_IsProvided(t *testing.T) {
	type fields struct {
		flag FlagBool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided flag",
			fields: fields{
				flag: FlagBool{},
			},
		},
		{
			name: "provided flag",
			fields: fields{
				flag: FlagBool{
					Name:         "version",
					Value:        true,
					flagValue:    AnyValue("true"),
					flagProvided: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagBoolValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvided(); got != tt.want {
				t.Errorf("FlagBoolValue.IsProvided() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagBoolValue_IsProvidedShort(t *testing.T) {
	type fields struct {
		flag FlagBool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided short flag",
			fields: fields{
				flag: FlagBool{},
			},
		},
		{
			name: "provided short flag",
			fields: fields{
				flag: FlagBool{
					Name:                "short",
					Value:               true,
					flagValue:           AnyValue("true"),
					flagProvided:        true,
					flagProvidedAsAlias: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagBoolValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvidedShort(); got != tt.want {
				t.Errorf("FlagBoolValue.IsProvidedShort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagBoolValue_IsProvidedLong(t *testing.T) {
	type fields struct {
		flag FlagBool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided long flag",
			fields: fields{
				flag: FlagBool{},
			},
		},
		{
			name: "provided long flag",
			fields: fields{
				flag: FlagBool{
					Name:                "long",
					Value:               true,
					flagValue:           AnyValue("true"),
					flagProvided:        true,
					flagProvidedAsAlias: false,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagBoolValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvidedLong(); got != tt.want {
				t.Errorf("FlagBoolValue.IsProvidedLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagBoolValue_GetFlagType(t *testing.T) {
	type fields struct {
		flag FlagBool
	}
	tests := []struct {
		name   string
		fields fields
		want   FlagBool
	}{
		{
			name: "get flag type",
			fields: fields{
				flag: FlagBool{},
			},
			want: FlagBool{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagBoolValue{
				flag: tt.fields.flag,
			}
			if got := v.GetFlagType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagBoolValue.GetFlagType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagStringValue_Value(t *testing.T) {
	type fields struct {
		flag FlagString
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty flag value",
			fields: fields{
				flag: FlagString{},
			},
			want: "",
		},
		{
			name: "non empty value",
			fields: fields{
				flag: FlagString{
					Name:      "abc",
					flagValue: AnyValue("xyz"),
				},
			},
			want: "xyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagStringValue{
				flag: tt.fields.flag,
			}
			if got := v.Value(); got != tt.want {
				t.Errorf("FlagStringValue.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagStringValue_IsProvided(t *testing.T) {
	type fields struct {
		flag FlagString
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided flag",
			fields: fields{
				flag: FlagString{},
			},
		},
		{
			name: "provided flag",
			fields: fields{
				flag: FlagString{
					Name:         "alpha",
					flagProvided: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagStringValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvided(); got != tt.want {
				t.Errorf("FlagStringValue.IsProvided() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagStringValue_IsProvidedShort(t *testing.T) {
	type fields struct {
		flag FlagString
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided flag",
			fields: fields{
				flag: FlagString{},
			},
		},
		{
			name: "provided flag",
			fields: fields{
				flag: FlagString{
					Name:                "alpha",
					flagProvided:        true,
					flagProvidedAsAlias: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagStringValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvidedShort(); got != tt.want {
				t.Errorf("FlagStringValue.IsProvidedShort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagStringValue_IsProvidedLong(t *testing.T) {
	type fields struct {
		flag FlagString
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided long flag",
			fields: fields{
				flag: FlagString{},
			},
		},
		{
			name: "provided long flag",
			fields: fields{
				flag: FlagString{
					Name:                "long",
					Value:               "cde",
					flagValue:           AnyValue("cde"),
					flagProvided:        true,
					flagProvidedAsAlias: false,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagStringValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvidedLong(); got != tt.want {
				t.Errorf("FlagStringValue.IsProvidedLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagStringValue_GetFlagType(t *testing.T) {
	type fields struct {
		flag FlagString
	}
	tests := []struct {
		name   string
		fields fields
		want   FlagString
	}{
		{
			name: "get flag type",
			fields: fields{
				flag: FlagString{},
			},
			want: FlagString{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagStringValue{
				flag: tt.fields.flag,
			}
			if got := v.GetFlagType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagStringValue.GetFlagType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagStringSliceValue_Value(t *testing.T) {
	type fields struct {
		flag FlagStringSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "empty flag value",
			fields: fields{
				flag: FlagStringSlice{},
			},
			want: []string{""},
		},
		{
			name: "non empty value",
			fields: fields{
				flag: FlagStringSlice{
					Name:      "abc",
					flagValue: AnyValue("xyz, abc"),
				},
			},
			want: []string{"xyz", "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagStringSliceValue{
				flag: tt.fields.flag,
			}
			if got := v.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagStringSliceValue.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagStringSliceValue_IsProvided(t *testing.T) {
	type fields struct {
		flag FlagStringSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided flag",
			fields: fields{
				flag: FlagStringSlice{},
			},
		},
		{
			name: "provided flag",
			fields: fields{
				flag: FlagStringSlice{
					Name:         "alpha",
					flagProvided: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagStringSliceValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvided(); got != tt.want {
				t.Errorf("FlagStringSliceValue.IsProvided() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagStringSliceValue_IsProvidedShort(t *testing.T) {
	type fields struct {
		flag FlagStringSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided short flag",
			fields: fields{
				flag: FlagStringSlice{},
			},
		},
		{
			name: "provided short flag false",
			fields: fields{
				flag: FlagStringSlice{
					Name:         "alpha",
					flagProvided: true,
				},
			},
		},
		{
			name: "provided short flag true",
			fields: fields{
				flag: FlagStringSlice{
					Name:                "alpha",
					flagProvided:        true,
					flagProvidedAsAlias: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagStringSliceValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvidedShort(); got != tt.want {
				t.Errorf("FlagStringSliceValue.IsProvidedShort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagStringSliceValue_IsProvidedLong(t *testing.T) {
	type fields struct {
		flag FlagStringSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided long flag",
			fields: fields{
				flag: FlagStringSlice{},
			},
		},
		{
			name: "provided long flag true",
			fields: fields{
				flag: FlagStringSlice{
					Name:         "alpha",
					flagProvided: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagStringSliceValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvidedLong(); got != tt.want {
				t.Errorf("FlagStringSliceValue.IsProvidedLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagStringSliceValue_GetFlagType(t *testing.T) {
	type fields struct {
		flag FlagStringSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   FlagStringSlice
	}{
		{
			name: "get flag type",
			fields: fields{
				flag: FlagStringSlice{},
			},
			want: FlagStringSlice{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagStringSliceValue{
				flag: tt.fields.flag,
			}
			if got := v.GetFlagType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagStringSliceValue.GetFlagType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValues_findByKey(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFlag Flag
	}{
		{
			name: "lookup on an unsupported type flag list",
			fields: fields{
				flags: []Flag{
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
				flags: []Flag{
					FlagInt{},
					FlagBool{},
					FlagString{},
					FlagStringSlice{},
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
				flags: []Flag{
					FlagInt{Name: "k-int"},
					FlagBool{Name: "k-bool"},
					FlagString{Name: "k-string"},
					FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				longFlagName: "k-int",
			},
			wantFlag: FlagInt{Name: "k-int"},
		},
		{
			name: "lookup using valid bool flag key",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int"},
					FlagBool{Name: "k-bool"},
					FlagString{Name: "k-string"},
					FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				longFlagName: "k-bool",
			},
			wantFlag: FlagBool{Name: "k-bool"},
		},
		{
			name: "lookup using valid string flag key",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int"},
					FlagBool{Name: "k-bool"},
					FlagString{Name: "k-string"},
					FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				longFlagName: "k-string",
			},
			wantFlag: FlagString{Name: "k-string"},
		},
		{
			name: "lookup using valid string-slice flag key",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int"},
					FlagBool{Name: "k-bool"},
					FlagString{Name: "k-string"},
					FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{
				longFlagName: "k-string-slice",
			},
			wantFlag: FlagStringSlice{Name: "k-string-slice"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			if gotFlag := v.findByKey(tt.args.longFlagName); !reflect.DeepEqual(gotFlag, tt.wantFlag) {
				t.Errorf("FlagValues.findByKey() = %v, want %v", gotFlag, tt.wantFlag)
			}
		})
	}
}

func TestFlagValues_getProvidedFlags(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		providedOnly bool
		aliasOnly    bool
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantFlags []Flag
	}{
		{
			name: "return all flags",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int"},
					FlagBool{Name: "k-bool"},
					FlagString{Name: "k-string"},
					FlagStringSlice{Name: "k-string-slice"},
				},
			},
			args: args{},
			wantFlags: []Flag{
				FlagInt{Name: "k-int"},
				FlagBool{Name: "k-bool"},
				FlagString{Name: "k-string"},
				FlagStringSlice{Name: "k-string-slice"},
			},
		},
		{
			name: "return empty list of flags",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int"},
					FlagBool{Name: "k-bool"},
					FlagString{Name: "k-string"},
					FlagStringSlice{Name: "k-string-slice"},
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
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true},
					FlagBool{Name: "k-bool"},
					FlagString{Name: "k-string", flagProvided: true},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true},
				},
			},
			args: args{
				providedOnly: true,
			},
			wantFlags: []Flag{
				FlagInt{Name: "k-int", flagProvided: true},
				FlagString{Name: "k-string", flagProvided: true},
				FlagStringSlice{Name: "k-string-slice", flagProvided: true},
			},
		},
		{
			name: "return provided short flags",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true},
					FlagBool{Name: "k-bool", flagProvided: true, flagProvidedAsAlias: true},
					FlagString{Name: "k-string", flagProvided: true, flagProvidedAsAlias: true},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true},
				},
			},
			args: args{
				aliasOnly: true,
			},
			wantFlags: []Flag{
				FlagBool{Name: "k-bool", flagProvided: true, flagProvidedAsAlias: true},
				FlagString{Name: "k-string", flagProvided: true, flagProvidedAsAlias: true},
			},
		},
		{
			name: "return provided flags only",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true},
					FlagBool{Name: "k-bool", flagProvided: true},
					FlagString{Name: "k-string", flagProvided: true},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true},
				},
			},
			args: args{
				providedOnly: true,
			},
			wantFlags: []Flag{
				FlagInt{Name: "k-int", flagProvided: true},
				FlagBool{Name: "k-bool", flagProvided: true},
				FlagString{Name: "k-string", flagProvided: true},
				FlagStringSlice{Name: "k-string-slice", flagProvided: true},
			},
		},
		{
			name: "return provided short flags only",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagProvidedAsAlias: true},
					FlagBool{Name: "k-bool", flagProvided: true, flagProvidedAsAlias: true},
					FlagString{Name: "k-string", flagProvided: true, flagProvidedAsAlias: true},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagProvidedAsAlias: true},
				},
			},
			args: args{
				aliasOnly: true,
			},
			wantFlags: []Flag{
				FlagInt{Name: "k-int", flagProvided: true, flagProvidedAsAlias: true},
				FlagBool{Name: "k-bool", flagProvided: true, flagProvidedAsAlias: true},
				FlagString{Name: "k-string", flagProvided: true, flagProvidedAsAlias: true},
				FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagProvidedAsAlias: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			if gotFlags := v.getProvidedFlags(tt.args.providedOnly, tt.args.aliasOnly); !reflect.DeepEqual(gotFlags, tt.wantFlags) {
				t.Errorf("FlagValues.getProvidedFlags() = %v, want %v", gotFlags, tt.wantFlags)
			}
		})
	}
}

func TestFlagValues_GetProvided(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	tests := []struct {
		name   string
		fields fields
		want   []Flag
	}{
		{
			name: "get provided nil flags",
			fields: fields{
				flags: []Flag{},
			},
			want: nil,
		},
		{
			name: "get provided flag list",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagProvidedAsAlias: true},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagProvidedAsAlias: true},
				},
			},
			want: []Flag{
				FlagInt{Name: "k-int", flagProvided: true, flagProvidedAsAlias: true},
				FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagProvidedAsAlias: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			if got := v.GetProvided(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValues.GetProvided() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestFlagValues_GetProvidedLong(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	tests := []struct {
		name   string
		fields fields
		want   []Flag
	}{
		{
			name: "get provided empty flags",
			fields: fields{
				flags: []Flag{},
			},
			want: []Flag{},
		},
		{
			name: "get provided flag list",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true},
				},
			},
			want: []Flag{
				FlagInt{Name: "k-int", flagProvided: true},
				FlagStringSlice{Name: "k-string-slice", flagProvided: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			if got := v.GetProvidedLong(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValues.GetProvidedLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValues_GetProvidedShort(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	tests := []struct {
		name   string
		fields fields
		want   []Flag
	}{
		{
			name: "get provided empty short flags",
			fields: fields{
				flags: []Flag{},
			},
			want: nil,
		},
		{
			name: "get provided short flag list",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagProvidedAsAlias: true},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagProvidedAsAlias: true},
				},
			},
			want: []Flag{
				FlagInt{Name: "k-int", flagProvided: true, flagProvidedAsAlias: true},
				FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagProvidedAsAlias: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			if got := v.GetProvidedShort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValues.GetProvidedShort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValues_Any(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   AnyValue
	}{
		{
			name: "lookup using empty bool flag list",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int"},
					FlagBool{Name: "k-bool"},
					FlagString{Name: "k-string"},
					FlagStringSlice{Name: "k-string-slice"},
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
				flags: []Flag{
					FlagInt{Name: "k-int"},
					FlagBool{Name: "k-bool"},
					FlagString{Name: "k-string"},
					FlagStringSlice{Name: "k-string-slice"},
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
				flags: []Flag{
					FlagInt{Name: "k-int"},
					FlagBool{Name: "k-bool"},
					FlagString{Name: "k-string"},
					FlagStringSlice{Name: "k-string-slice"},
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
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagValue: "5"},
					FlagBool{Name: "k-bool", flagProvided: true, flagValue: "true"},
					FlagString{Name: "k-string", flagProvided: true, flagValue: "string_val"},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagValue: "a,b,c"},
				},
			},
			args: args{
				longFlagName: "k-int",
			},
			want: AnyValue("5"),
		},
		{
			name: "lookup using empty string-slice flag list",
			fields: fields{
				flags: []Flag{
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
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			if got := v.Any(tt.args.longFlagName); got != tt.want {
				t.Errorf("FlagValues.Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValues_Bool(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FlagBoolValue
		wantErr bool
	}{
		{
			name: "get invalid bool value",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagValue: "5"},
					FlagBool{Name: "k-bool", flagProvided: false, flagValue: "true"},
					FlagString{Name: "k-string", flagProvided: true, flagValue: "string_val"},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagValue: "a,b,c"},
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
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagValue: "5"},
					FlagBool{Name: "k-bool", flagProvided: true, flagValue: "true"},
					FlagString{Name: "k-string", flagProvided: true, flagValue: "string_val"},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagValue: "a,b,c"},
				},
			},
			args: args{
				longFlagName: "k-bool",
			},
			want: &FlagBoolValue{
				flag: FlagBool{Name: "k-bool", flagProvided: true, flagValue: "true"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			got, err := v.Bool(tt.args.longFlagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagValues.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValues.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValues_Int(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FlagIntValue
		wantErr bool
	}{
		{
			name: "get invalid int value",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagValue: "5"},
					FlagBool{Name: "k-bool", flagProvided: false, flagValue: "true"},
					FlagString{Name: "k-string", flagProvided: true, flagValue: "string_val"},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagValue: "a,b,c"},
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
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagValue: "256"},
					FlagBool{Name: "k-bool", flagProvided: true, flagValue: "true"},
					FlagString{Name: "k-string", flagProvided: true, flagValue: "string_val"},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagValue: "a,b,c"},
				},
			},
			args: args{
				longFlagName: "k-int",
			},
			want: &FlagIntValue{
				flag: FlagInt{Name: "k-int", flagProvided: true, flagValue: "256"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			got, err := v.Int(tt.args.longFlagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagValues.Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValues.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValues_String(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FlagStringValue
		wantErr bool
	}{
		{
			name: "get invalid string value",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagValue: "5"},
					FlagBool{Name: "k-bool", flagProvided: false, flagValue: "true"},
					FlagString{Name: "k-string", flagProvided: true, flagValue: "string_val"},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagValue: "a,b,c"},
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
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagValue: "256"},
					FlagBool{Name: "k-bool", flagProvided: true, flagValue: "true"},
					FlagString{Name: "k-string", flagProvided: true, flagValue: "string_val"},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagValue: "a,b,c"},
				},
			},
			args: args{
				longFlagName: "k-string",
			},
			want: &FlagStringValue{
				flag: FlagString{Name: "k-string", flagProvided: true, flagValue: "string_val"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			got, err := v.String(tt.args.longFlagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagValues.String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValues.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValues_StringSlice(t *testing.T) {
	type fields struct {
		flags []Flag
	}
	type args struct {
		longFlagName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FlagStringSliceValue
		wantErr bool
	}{
		{
			name: "get invalid string-slice value",
			fields: fields{
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagValue: "5"},
					FlagBool{Name: "k-bool", flagProvided: false, flagValue: "true"},
					FlagString{Name: "k-string", flagProvided: true, flagValue: "string_val"},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagValue: "1,2,3"},
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
				flags: []Flag{
					FlagInt{Name: "k-int", flagProvided: true, flagValue: "256"},
					FlagBool{Name: "k-bool", flagProvided: true, flagValue: "true"},
					FlagString{Name: "k-string", flagProvided: true, flagValue: "string_val"},
					FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagValue: "1,2,3"},
				},
			},
			args: args{
				longFlagName: "k-string-slice",
			},
			want: &FlagStringSliceValue{
				flag: FlagStringSlice{Name: "k-string-slice", flagProvided: true, flagValue: "1,2,3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			got, err := v.StringSlice(tt.args.longFlagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagValues.StringSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValues.StringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagIntValue_Value(t *testing.T) {
	type fields struct {
		flag FlagInt
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "no provided value flag",
			fields: fields{
				flag: FlagInt{},
			},
			wantErr: true,
		},
		{
			name: "provided value flag",
			fields: fields{
				flag: FlagInt{
					Name:         "short",
					Value:        7,
					flagValue:    AnyValue("7"),
					flagProvided: true,
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagIntValue{
				flag: tt.fields.flag,
			}
			got, err := v.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagIntValue.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlagIntValue.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagIntValue_IsProvided(t *testing.T) {
	type fields struct {
		flag FlagInt
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided flag",
			fields: fields{
				flag: FlagInt{},
			},
		},
		{
			name: "provided flag",
			fields: fields{
				flag: FlagInt{
					Name:         "provided",
					Value:        32,
					flagValue:    AnyValue("32"),
					flagProvided: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagIntValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvided(); got != tt.want {
				t.Errorf("FlagIntValue.IsProvided() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagIntValue_IsProvidedShort(t *testing.T) {
	type fields struct {
		flag FlagInt
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided short flag",
			fields: fields{
				flag: FlagInt{},
			},
		},
		{
			name: "provided short flag",
			fields: fields{
				flag: FlagInt{
					Name:                "provided",
					flagProvided:        true,
					flagProvidedAsAlias: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagIntValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvidedShort(); got != tt.want {
				t.Errorf("FlagIntValue.IsProvidedShort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagIntValue_IsProvidedLong(t *testing.T) {
	type fields struct {
		flag FlagInt
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "no provided long flag",
			fields: fields{
				flag: FlagInt{},
			},
		},
		{
			name: "provided long flag",
			fields: fields{
				flag: FlagInt{
					Name:                "provided",
					flagProvided:        true,
					flagProvidedAsAlias: false,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagIntValue{
				flag: tt.fields.flag,
			}
			if got := v.IsProvidedLong(); got != tt.want {
				t.Errorf("FlagIntValue.IsProvidedLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagIntValue_GetFlagType(t *testing.T) {
	type fields struct {
		flag FlagInt
	}
	tests := []struct {
		name   string
		fields fields
		want   FlagInt
	}{
		{
			name: "get flag type",
			fields: fields{
				flag: FlagInt{},
			},
			want: FlagInt{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagIntValue{
				flag: tt.fields.flag,
			}
			if got := v.GetFlagType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagIntValue.GetFlagType() = %v, want %v", got, tt.want)
			}
		})
	}
}
