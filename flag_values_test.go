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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.StringSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValue.StringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValueMap_findByKey(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagValueMap{
				flags: tt.fields.flags,
			}
			if got := fm.findByKey(tt.args.flagKey); got != tt.want {
				t.Errorf("FlagValueMap.findByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValueMap_Bool(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagValueMap{
				flags: tt.fields.flags,
			}
			got, err := fm.Bool(tt.args.flagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagValueMap.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlagValueMap.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValueMap_Int(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagValueMap{
				flags: tt.fields.flags,
			}
			got, err := fm.Int(tt.args.flagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlagValueMap.Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FlagValueMap.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValueMap_String(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagValueMap{
				flags: tt.fields.flags,
			}
			if got := fm.String(tt.args.flagName); got != tt.want {
				t.Errorf("FlagValueMap.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagValueMap_StringSlice(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlagValueMap{
				flags: tt.fields.flags,
			}
			if got := fm.StringSlice(tt.args.flagName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagValueMap.StringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
