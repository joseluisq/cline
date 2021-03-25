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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			if got := v.GetProvided(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValues.GetProvided() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		name   string
		fields fields
		args   args
		want   *FlagBoolValue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			if got := v.Bool(tt.args.longFlagName); !reflect.DeepEqual(got, tt.want) {
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
		name   string
		fields fields
		args   args
		want   *FlagIntValue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			if got := v.Int(tt.args.longFlagName); !reflect.DeepEqual(got, tt.want) {
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
		name   string
		fields fields
		args   args
		want   *FlagStringValue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			if got := v.String(tt.args.longFlagName); !reflect.DeepEqual(got, tt.want) {
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
		name   string
		fields fields
		args   args
		want   *FlagStringSliceValue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &FlagValues{
				flags: tt.fields.flags,
			}
			if got := v.StringSlice(tt.args.longFlagName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValues.StringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
