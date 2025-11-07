package flag_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joseluisq/cline/flag"
)

func TestAnyValue_ToBool(t *testing.T) {
	tests := []struct {
		name        string
		value       flag.Value
		expected    bool
		expectedErr error
	}{
		{
			name:        "should fail parsing when invalid empty value",
			value:       flag.Value(""),
			expectedErr: errors.New("strconv.ParseBool: parsing \"\": invalid syntax"),
		},
		{
			name:        "should fail parsing when invalid value",
			value:       flag.Value("abc"),
			expectedErr: errors.New("strconv.ParseBool: parsing \"abc\": invalid syntax"),
		},
		{
			name:     "should succeed parsing when valid value",
			value:    flag.Value("1"),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actualVal, actualErr := tt.value.ToBool(); tt.expectedErr != nil {
				assert.Error(t, actualErr, "Expected an error but got none")
				assert.Equal(t, actualErr.Error(), tt.expectedErr.Error(), "Error message does not match the expected one")
			} else {
				assert.NoError(t, actualErr, "Expected no error but got one")
				assert.Equal(t, actualVal, tt.expected, "Bool value does not match the expected one")
			}
		})
	}
}

func TestAnyValue_ToInt(t *testing.T) {
	tests := []struct {
		name        string
		value       flag.Value
		expected    int
		expectedErr error
	}{
		{
			name:        "should fail parsing when invalid empty value",
			value:       flag.Value(""),
			expectedErr: fmt.Errorf("strconv.Atoi: parsing \"\": invalid syntax"),
		},
		{
			name:        "should fail parsing when invalid value",
			value:       flag.Value("10.14"),
			expectedErr: fmt.Errorf("strconv.Atoi: parsing \"10.14\": invalid syntax"),
		},
		{
			name:     "should succeed parsing when valid value",
			value:    flag.Value("10"),
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actualVal, actualErr := tt.value.ToInt(); tt.expectedErr != nil {
				assert.Error(t, actualErr, "Expected an error but got none")
				assert.Equal(t, actualErr.Error(), tt.expectedErr.Error(), "Error message does not match the expected one")
			} else {
				assert.NoError(t, actualErr, "Expected no error but got one")
				assert.Equal(t, actualVal, tt.expected, "Int value does not match the expected one")
			}
		})
	}
}

func TestAnyValue_ToString(t *testing.T) {
	tests := []struct {
		name     string
		value    flag.Value
		expected string
	}{
		{
			name: "should return empty string when empty value",
		},
		{
			name:     "should return string value",
			value:    flag.Value("abc"),
			expected: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualVal := tt.value.ToString()
			assert.Equal(t, actualVal, tt.expected, "String value does not match the expected one")
		})
	}
}

func TestAnyValue_ToStringSlice(t *testing.T) {
	tests := []struct {
		name     string
		value    flag.Value
		expected []string
	}{
		{
			name:     "should return empty slice when empty value",
			value:    flag.Value(""),
			expected: []string{""},
		},
		{
			name:     "should return string slice value",
			value:    flag.Value("a,b,c,d"),
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "should return string slice value with empty elements",
			value:    flag.Value("abc,   ,"),
			expected: []string{"abc", "", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualValues := tt.value.ToStringSlice()
			assert.Equal(t, actualValues, tt.expected, "String value does not match the expected one")
		})
	}
}

func TestFlagBoolValue_Value(t *testing.T) {
	type fields struct {
		flag flag.FlagBool
	}
	tests := []struct {
		name        string
		fields      fields
		expected    bool
		expectedErr error
	}{
		{
			name: "should fail parsing when empty value",
			fields: fields{
				flag: flag.FlagBool{},
			},
			expectedErr: errors.New("strconv.ParseBool: parsing \"\": invalid syntax"),
		},
		{
			name: "should succeed parsing when valid value",
			fields: fields{
				flag: flag.FlagBool{
					Name:      "version",
					Value:     true,
					FlagValue: flag.Value("true"),
				},
			},
			expected: true,
		},
		{
			name: "should fail parsing when invalid value",
			fields: fields{
				flag: flag.FlagBool{
					Name:      "version",
					Value:     true,
					FlagValue: flag.Value("abc"),
				},
			},
			expectedErr: errors.New("strconv.ParseBool: parsing \"abc\": invalid syntax"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := flag.ValueBool{
				Flag: tt.fields.flag,
			}
			if actualVal, actualErr := v.Value(); tt.expectedErr != nil {
				assert.Error(t, actualErr, "Expected an error but got none")
				assert.Equal(t, actualErr.Error(), tt.expectedErr.Error(), "Error message does not match the expected one")
			} else {
				assert.NoError(t, actualErr, "Expected no error but got one")
				assert.Equal(t, actualVal, tt.expected, "Bool value does not match the expected one")
			}
		})
	}
}

func TestFlagBoolValue_IsProvided(t *testing.T) {
	type fields struct {
		flag flag.FlagBool
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "should return false when no provided flag",
			fields: fields{
				flag: flag.FlagBool{},
			},
		},
		{
			name: "should return true when provided flag",
			fields: fields{
				flag: flag.FlagBool{
					Name:         "version",
					Value:        true,
					FlagValue:    flag.Value("true"),
					FlagProvided: true,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := flag.ValueBool{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvided()
			assert.Equal(t, actualVal, tt.expected, "IsProvided value does not match the expected one")
		})
	}
}

func TestFlagBoolValue_IsProvidedShort(t *testing.T) {
	type fields struct {
		flag flag.FlagBool
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "should return false when no provided short flag",
			fields: fields{
				flag: flag.FlagBool{},
			},
		},
		{
			name: "should return true when provided short flag",
			fields: fields{
				flag: flag.FlagBool{
					Name:                "short",
					Value:               true,
					FlagValue:           flag.Value("true"),
					FlagProvided:        true,
					FlagProvidedAsAlias: true,
				},
			},
			expected: true,
		},
		{
			name: "should return false when provided value is false",
			fields: fields{
				flag: flag.FlagBool{
					Name:                "short",
					Value:               true,
					FlagValue:           flag.Value("true"),
					FlagProvided:        false,
					FlagProvidedAsAlias: true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := flag.ValueBool{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedShort()
			assert.Equal(t, actualVal, tt.expected, "IsProvidedShort value does not match the expected one")
		})
	}
}

func TestFlagBoolValue_IsProvidedLong(t *testing.T) {
	type fields struct {
		flag flag.FlagBool
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "should return false when no provided long flag",
			fields: fields{
				flag: flag.FlagBool{},
			},
		},
		{
			name: "should return true when provided long flag",
			fields: fields{
				flag: flag.FlagBool{
					Name:                "long",
					Value:               true,
					FlagValue:           flag.Value("true"),
					FlagProvided:        true,
					FlagProvidedAsAlias: false,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := flag.ValueBool{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedLong()
			assert.Equal(t, actualVal, tt.expected, "IsProvidedLong value does not match the expected one")
		})
	}
}

func TestFlagBoolValue_GetFlagType(t *testing.T) {
	type fields struct {
		flag flag.FlagBool
	}
	tests := []struct {
		name     string
		fields   fields
		expected flag.FlagBool
	}{
		{
			name: "should get flag type",
			fields: fields{
				flag: flag.FlagBool{},
			},
			expected: flag.FlagBool{},
		},
		{
			name: "should get flag type with values",
			fields: fields{
				flag: flag.FlagBool{
					Name:         "long",
					Value:        true,
					FlagValue:    flag.Value("true"),
					FlagProvided: true,
				},
			},
			expected: flag.FlagBool{
				Name:         "long",
				Value:        true,
				FlagValue:    flag.Value("true"),
				FlagProvided: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := flag.ValueBool{
				Flag: tt.fields.flag,
			}
			actualVal := v.GetFlagType()
			assert.Equal(t, actualVal, tt.expected, "GetFlagType value does not match the expected one")
		})
	}
}

func TestFlagStringValue_Value(t *testing.T) {
	type fields struct {
		flag flag.FlagString
	}
	tests := []struct {
		name     string
		fields   fields
		expected string
	}{
		{
			name: "should return empty string when empty value",
			fields: fields{
				flag: flag.FlagString{},
			},
		},
		{
			name: "should return string value",
			fields: fields{
				flag: flag.FlagString{
					Name:      "abc",
					FlagValue: flag.Value("xyz"),
				},
			},
			expected: "xyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueString{
				Flag: tt.fields.flag,
			}
			actualVal := v.Value()
			assert.Equal(t, actualVal, tt.expected, "String value does not match the expected one")
		})
	}
}

func TestFlagStringValue_IsProvided(t *testing.T) {
	type fields struct {
		flag flag.FlagString
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "should return false when no provided flag",
			fields: fields{
				flag: flag.FlagString{},
			},
		},
		{
			name: "should return true when provided flag",
			fields: fields{
				flag: flag.FlagString{
					Name:         "alpha",
					FlagProvided: true,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueString{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvided()
			assert.Equal(t, actualVal, tt.expected, "IsProvided value does not match the expected one")
		})
	}
}

func TestFlagStringValue_IsProvidedShort(t *testing.T) {
	type fields struct {
		flag flag.FlagString
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "should return false when no provided short flag",
			fields: fields{
				flag: flag.FlagString{},
			},
		},
		{
			name: "should return true when provided short flag",
			fields: fields{
				flag: flag.FlagString{
					Name:                "alpha",
					FlagProvided:        true,
					FlagProvidedAsAlias: true,
				},
			},
			expected: true,
		},
		{
			name: "should return false when provided value is false",
			fields: fields{
				flag: flag.FlagString{
					Name:                "alpha",
					FlagProvided:        false,
					FlagProvidedAsAlias: true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueString{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedShort()
			assert.Equal(t, actualVal, tt.expected, "IsProvidedShort value does not match the expected one")
		})
	}
}

func TestFlagStringValue_IsProvidedLong(t *testing.T) {
	type fields struct {
		flag flag.FlagString
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "should return false when no provided long flag",
			fields: fields{
				flag: flag.FlagString{},
			},
		},
		{
			name: "should return true when provided long flag",
			fields: fields{
				flag: flag.FlagString{
					Name:                "long",
					Value:               "cde",
					FlagValue:           flag.Value("cde"),
					FlagProvided:        true,
					FlagProvidedAsAlias: false,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueString{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedLong()
			assert.Equal(t, actualVal, tt.expected, "IsProvidedLong value does not match the expected one")
		})
	}
}

func TestFlagStringValue_GetFlagType(t *testing.T) {
	type fields struct {
		flag flag.FlagString
	}
	tests := []struct {
		name     string
		fields   fields
		expected flag.FlagString
	}{
		{
			name: "should get flag type",
			fields: fields{
				flag: flag.FlagString{},
			},
			expected: flag.FlagString{},
		},
		{
			name: "should get flag type with values",
			fields: fields{
				flag: flag.FlagString{
					Name:         "long",
					Value:        "cde",
					FlagValue:    flag.Value("cde"),
					FlagProvided: true,
				},
			},
			expected: flag.FlagString{
				Name:         "long",
				Value:        "cde",
				FlagValue:    flag.Value("cde"),
				FlagProvided: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueString{
				Flag: tt.fields.flag,
			}
			actualVal := v.GetFlagType()
			assert.Equal(t, actualVal, tt.expected, "GetFlagType value does not match the expected one")
		})
	}
}

func TestFlagStringSliceValue_Value(t *testing.T) {
	type fields struct {
		flag flag.FlagStringSlice
	}
	tests := []struct {
		name     string
		fields   fields
		expected []string
	}{
		{
			name: "should return empty slice when empty value",
			fields: fields{
				flag: flag.FlagStringSlice{},
			},
			expected: []string{""},
		},
		{
			name: "should return string slice value",
			fields: fields{
				flag: flag.FlagStringSlice{
					Name:      "abc",
					FlagValue: flag.Value("xyz, abc"),
				},
			},
			expected: []string{"xyz", "abc"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueStringSlice{
				Flag: tt.fields.flag,
			}
			actualVal := v.Value()
			assert.Equal(t, actualVal, tt.expected, "String slice value does not match the expected one")
		})
	}
}

func TestFlagStringSliceValue_IsProvided(t *testing.T) {
	type fields struct {
		flag flag.FlagStringSlice
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "should return false when no provided flag",
			fields: fields{
				flag: flag.FlagStringSlice{},
			},
		},
		{
			name: "should return true when provided flag",
			fields: fields{
				flag: flag.FlagStringSlice{
					Name:         "alpha",
					FlagProvided: true,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueStringSlice{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvided()
			assert.Equal(t, actualVal, tt.expected, "IsProvided value does not match the expected one")
		})
	}
}

func TestFlagStringSliceValue_IsProvidedShort(t *testing.T) {
	type fields struct {
		flag flag.FlagStringSlice
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "should return false when no provided short flag",
			fields: fields{
				flag: flag.FlagStringSlice{},
			},
		},
		{
			name: "should return true when provided short flag",
			fields: fields{
				flag: flag.FlagStringSlice{
					Name:         "alpha",
					FlagProvided: true,
				},
			},
		},
		{
			name: "should return false when provided value is false",
			fields: fields{
				flag: flag.FlagStringSlice{
					Name:                "alpha",
					FlagProvided:        false,
					FlagProvidedAsAlias: true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueStringSlice{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedShort()
			assert.Equal(t, actualVal, tt.expected, "IsProvidedShort value does not match the expected one")
		})
	}
}

func TestFlagStringSliceValue_IsProvidedLong(t *testing.T) {
	type fields struct {
		flag flag.FlagStringSlice
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "should return false when no provided long flag",
			fields: fields{
				flag: flag.FlagStringSlice{},
			},
		},
		{
			name: "should return true when provided long flag",
			fields: fields{
				flag: flag.FlagStringSlice{
					Name:         "alpha",
					FlagProvided: true,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueStringSlice{
				Flag: tt.fields.flag,
			}
			actualVal := v.IsProvidedLong()
			assert.Equal(t, actualVal, tt.expected, "IsProvidedLong value does not match the expected one")
		})
	}
}

func TestFlagStringSliceValue_GetFlagType(t *testing.T) {
	type fields struct {
		flag flag.FlagStringSlice
	}
	tests := []struct {
		name     string
		fields   fields
		expected flag.FlagStringSlice
	}{
		{
			name: "should get flag type",
			fields: fields{
				flag: flag.FlagStringSlice{},
			},
			expected: flag.FlagStringSlice{},
		},
		{
			name: "should get flag type with values",
			fields: fields{
				flag: flag.FlagStringSlice{
					Name:         "alpha",
					Value:        []string{"a", "b"},
					FlagValue:    flag.Value("a,b"),
					FlagProvided: true,
				},
			},
			expected: flag.FlagStringSlice{
				Name:         "alpha",
				Value:        []string{"a", "b"},
				FlagValue:    flag.Value("a,b"),
				FlagProvided: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.ValueStringSlice{
				Flag: tt.fields.flag,
			}
			actualVal := v.GetFlagType()
			assert.Equal(t, actualVal, tt.expected, "GetFlagType value does not match the expected one")
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
		name         string
		fields       fields
		args         args
		expectedFlag flag.Flag
	}{
		{
			name: "should lookup on an unsupported type flag list",
			fields: fields{
				flags: []flag.Flag{
					map[string]int{"a": 1, "b": 2},
				},
			},
			args: args{
				longFlagName: "xyz",
			},
		},
		{
			name: "should lookup using empty flag key",
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
		},
		{
			name: "should lookup using valid int flag key",
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
			expectedFlag: flag.FlagInt{Name: "k-int"},
		},
		{
			name: "should lookup using valid bool flag key",
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
			expectedFlag: flag.FlagBool{Name: "k-bool"},
		},
		{
			name: "should lookup using valid string flag key",
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
			expectedFlag: flag.FlagString{Name: "k-string"},
		},
		{
			name: "should lookup using valid string-slice flag key",
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
			expectedFlag: flag.FlagStringSlice{Name: "k-string-slice"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			actualFlag := v.FindByKey(tt.args.longFlagName)
			assert.Equal(t, actualFlag, tt.expectedFlag, "FindByKey value does not match the expected one")
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
		name          string
		fields        fields
		args          args
		expectedFlags []flag.Flag
	}{
		{
			name: "should return all flags",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int"},
					flag.FlagBool{Name: "k-bool"},
					flag.FlagString{Name: "k-string"},
					flag.FlagStringSlice{Name: "k-string-slice"},
				},
			},
			expectedFlags: []flag.Flag{
				flag.FlagInt{Name: "k-int"},
				flag.FlagBool{Name: "k-bool"},
				flag.FlagString{Name: "k-string"},
				flag.FlagStringSlice{Name: "k-string-slice"},
			},
		},
		{
			name: "should return empty list of flags",
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
		},
		{
			name: "should return only provided flags",
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
			expectedFlags: []flag.Flag{
				flag.FlagInt{Name: "k-int", FlagProvided: true},
				flag.FlagString{Name: "k-string", FlagProvided: true},
				flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true},
			},
		},
		{
			name: "should return only alias provided flags",
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
			expectedFlags: []flag.Flag{
				flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagProvidedAsAlias: true},
				flag.FlagString{Name: "k-string", FlagProvided: true, FlagProvidedAsAlias: true},
			},
		},
		{
			name: "should return provided flags only",
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
			expectedFlags: []flag.Flag{
				flag.FlagInt{Name: "k-int", FlagProvided: true},
				flag.FlagBool{Name: "k-bool", FlagProvided: true},
				flag.FlagString{Name: "k-string", FlagProvided: true},
				flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true},
			},
		},
		{
			name: "should return provided short flags only",
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
			expectedFlags: []flag.Flag{
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
			assert.Equal(t, actual, tt.expectedFlags, "GetProvidedFlags value does not match the expected one")
		})
	}
}

func TestFlagValues_GetProvided(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	tests := []struct {
		name     string
		fields   fields
		expected []flag.Flag
	}{
		{
			name: "should get provided nil flags",
			fields: fields{
				flags: []flag.Flag{},
			},
		},
		{
			name: "should get provided flag list",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagProvidedAsAlias: true},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagProvidedAsAlias: true},
				},
			},
			expected: []flag.Flag{
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
			assert.Equal(t, actual, tt.expected, "GetProvided value does not match the expected one")
		})
	}
}

func TestFlagValues_GetProvidedLong(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	tests := []struct {
		name     string
		fields   fields
		expected []flag.Flag
	}{
		{
			name: "should get provided empty flags",
			fields: fields{
				flags: []flag.Flag{},
			},
			expected: []flag.Flag{},
		},
		{
			name: "should get provided flag list",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true},
				},
			},
			expected: []flag.Flag{
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
			assert.Equal(t, actual, tt.expected, "GetProvidedLong value does not match the expected one")
		})
	}
}

func TestFlagValues_GetProvidedShort(t *testing.T) {
	type fields struct {
		flags []flag.Flag
	}
	tests := []struct {
		name     string
		fields   fields
		expected []flag.Flag
	}{
		{
			name: "should get provided empty short flags",
			fields: fields{
				flags: []flag.Flag{},
			},
			expected: nil,
		},
		{
			name: "should get provided short flag list",
			fields: fields{
				flags: []flag.Flag{
					flag.FlagInt{Name: "k-int", FlagProvided: true, FlagProvidedAsAlias: true},
					flag.FlagStringSlice{Name: "k-string-slice", FlagProvided: true, FlagProvidedAsAlias: true},
				},
			},
			expected: []flag.Flag{
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
			assert.Equal(t, actual, tt.expected, "GetProvidedShort value does not match the expected one")
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
		name     string
		fields   fields
		args     args
		expected flag.Value
	}{
		{
			name: "should lookup using empty bool flag list",
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
		},
		{
			name: "should lookup using empty string flag list",
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
		},
		{
			name: "should lookup using empty string-slice flag list",
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
		},
		{
			name: "should lookup using valid int flag key",
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
			expected: flag.Value("5"),
		},
		{
			name: "should lookup using empty string-slice flag list",
			fields: fields{
				flags: []flag.Flag{
					map[string]int{"a": 1, "b": 2},
				},
			},
			args: args{
				longFlagName: "k-any",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			actual := v.Value(tt.args.longFlagName)
			assert.Equal(t, actual, tt.expected, "Value does not match the expected one")
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
		name        string
		fields      fields
		args        args
		expected    *flag.ValueBool
		expectedErr error
	}{
		{
			name: "should get invalid bool value",
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
			expectedErr: errors.New("error: flag `--some` value used as `FlagBoolValue` but declared as `<nil>`"),
		},
		{
			name: "should get valid bool value",
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
			expected: &flag.ValueBool{
				Flag: flag.FlagBool{Name: "k-bool", FlagProvided: true, FlagValue: "true"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			if actual, err := v.Bool(tt.args.longFlagName); tt.expectedErr != nil {
				assert.Error(t, err, "expected error but got none")
				assert.Equal(t, err.Error(), tt.expectedErr.Error(), "error messages do not match")
			} else {
				assert.NoError(t, err, "did not expect error but got one")
				assert.Equal(t, actual, tt.expected, "Bool value do not match")
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
		name        string
		fields      fields
		args        args
		expected    *flag.ValueInt
		expectedErr error
	}{
		{
			name: "should get invalid int value",
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
			expectedErr: errors.New("error: flag `--some` value used as `FlagIntValue` but declared as `<nil>`"),
		},
		{
			name: "should get valid int value",
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
			expected: &flag.ValueInt{
				Flag: flag.FlagInt{Name: "k-int", FlagProvided: true, FlagValue: "256"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &flag.FlagValues{
				Flags: tt.fields.flags,
			}
			if actual, err := v.Int(tt.args.longFlagName); tt.expectedErr != nil {
				assert.Error(t, err, "expected error but got none")
				assert.Equal(t, err.Error(), tt.expectedErr.Error(), "error messages do not match")
			} else {
				assert.NoError(t, err, "did not expect error but got one")
				assert.Equal(t, actual, tt.expected, "Int value do not match")
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
		name        string
		fields      fields
		args        args
		expected    *flag.ValueString
		expectedErr error
	}{
		{
			name: "should get invalid string value",
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
			expectedErr: errors.New("error: flag `--some` value used as `FlagStringValue` but declared as `<nil>`"),
		},
		{
			name: "should get valid string value",
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
			if actual, err := v.String(tt.args.longFlagName); tt.expectedErr != nil {
				assert.Error(t, err, "expected error but got none")
				assert.Equal(t, err.Error(), tt.expectedErr.Error(), "error messages do not match")
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
		name        string
		fields      fields
		args        args
		expected    *flag.ValueStringSlice
		expectedErr error
	}{
		{
			name: "should get invalid string-slice value",
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
			expectedErr: errors.New("error: flag `--some` value used as `FlagStringSliceValue` but declared as `<nil>`"),
		},
		{
			name: "should get valid string-slice value",
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
			if actual, err := v.StringSlice(tt.args.longFlagName); tt.expectedErr != nil {
				assert.Error(t, err, "expected error but got none")
				assert.Equal(t, err.Error(), tt.expectedErr.Error(), "error messages do not match")
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
		name        string
		fields      fields
		expected    int
		expectedErr error
	}{
		{
			name: "should return error when no provided flag value",
			fields: fields{
				flag: flag.FlagInt{},
			},
			expectedErr: errors.New("strconv.Atoi: parsing \"\": invalid syntax"),
		},
		{
			name: "should return value when valid provided flag value",
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

			if actual, err := v.Value(); tt.expectedErr != nil {
				assert.Error(t, err, "expected error but got none")
				assert.Equal(t, err.Error(), tt.expectedErr.Error(), "error messages do not match")
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
			name: "should return false when no provided flag",
			fields: fields{
				flag: flag.FlagInt{},
			},
		},
		{
			name: "should return true when provided flag",
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
			name: "should return false when no provided short flag",
			fields: fields{
				flag: flag.FlagInt{},
			},
		},
		{
			name: "should return true when provided short flag",
			fields: fields{
				flag: flag.FlagInt{
					Name:                "provided",
					FlagProvided:        true,
					FlagProvidedAsAlias: true,
				},
			},
			expected: true,
		},
		{
			name: "should return false when provided value is false",
			fields: fields{
				flag: flag.FlagInt{
					Name:                "short",
					FlagValue:           flag.Value("1"),
					FlagProvided:        false,
					FlagProvidedAsAlias: true,
				},
			},
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
			name: "should return false when no provided long flag",
			fields: fields{
				flag: flag.FlagInt{},
			},
		},
		{
			name: "should return true when provided long flag",
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
			name: "should get flag type",
			fields: fields{
				flag: flag.FlagInt{},
			},
			expected: flag.FlagInt{},
		},
		{
			name: "should get flag type with values",
			fields: fields{
				flag: flag.FlagInt{
					Name:         "long",
					Value:        7,
					FlagValue:    flag.Value("true"),
					FlagProvided: true,
				},
			},
			expected: flag.FlagInt{
				Name:         "long",
				Value:        7,
				FlagValue:    flag.Value("true"),
				FlagProvided: true,
			},
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
