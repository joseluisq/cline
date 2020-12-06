package cline

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// AnyValue is an alias of string type which represents an input value for a command flag.
type AnyValue string

// ToBool converts current flag value into a `bool`.
func (v AnyValue) ToBool() (bool, error) {
	return strconv.ParseBool(v.ToString())
}

// ToInt converts current flag value into an `int`.
func (v AnyValue) ToInt() (int, error) {
	return strconv.Atoi(v.ToString())
}

// ToString converts current flag value into a `string`.
func (v AnyValue) ToString() string {
	return string(v)
}

// ToStringSlice converts current flag value into a string slice.
func (v AnyValue) ToStringSlice() []string {
	var strs []string
	for _, s := range strings.Split(string(v), ",") {
		strs = append(strs, strings.TrimSpace(s))
	}
	return strs
}

// FlagBoolValue represents a flag value bool type.
type FlagBoolValue struct {
	flag FlagBool
}

// Value unwraps the plain bool value of the current flag.
func (v *FlagBoolValue) Value() (bool, error) {
	return v.flag.flagValue.ToBool()
}

// IsProvided checks if current bool flag was provided from stdin.
func (v *FlagBoolValue) IsProvided() bool {
	return v.flag.flagProvided
}

// IsProvidedShort checks if current bool flag was provided from stdin but using its short name.
func (v *FlagBoolValue) IsProvidedShort() bool {
	return v.flag.flagProvided && v.flag.flagProvidedAsAlias
}

// IsProvidedLong checks if current bool flag was provided from stdin but using its long name.
func (v *FlagBoolValue) IsProvidedLong() bool {
	return v.flag.flagProvided && !v.flag.flagProvidedAsAlias
}

// GetFlagType returns the associated flag type.
func (v *FlagBoolValue) GetFlagType() FlagBool {
	return v.flag
}

// FlagIntValue represents a flag value int type.
type FlagIntValue struct {
	flag FlagInt
}

// Value unwraps the plain int value of the current flag.
func (v *FlagIntValue) Value() (int, error) {
	return v.flag.flagValue.ToInt()
}

// IsProvided checks if current int flag was provided from stdin.
func (v *FlagIntValue) IsProvided() bool {
	return v.flag.flagProvided
}

// IsProvidedShort checks if current int flag was provided from stdin but using its short name.
func (v *FlagIntValue) IsProvidedShort() bool {
	return v.flag.flagProvided && v.flag.flagProvidedAsAlias
}

// IsProvidedLong checks if current int flag was provided from stdin but using its long name.
func (v *FlagIntValue) IsProvidedLong() bool {
	return v.flag.flagProvided && !v.flag.flagProvidedAsAlias
}

// GetFlagType returns the associated flag type.
func (v *FlagIntValue) GetFlagType() FlagInt {
	return v.flag
}

// FlagStringValue represents a flag value string type.
type FlagStringValue struct {
	flag FlagString
}

// Value unwraps the plain string value of the current flag.
func (v *FlagStringValue) Value() string {
	return v.flag.flagValue.ToString()
}

// IsProvided checks if current string flag was provided from stdin.
func (v *FlagStringValue) IsProvided() bool {
	return v.flag.flagProvided
}

// IsProvidedShort checks if current string flag was provided from stdin but using its short name.
func (v *FlagStringValue) IsProvidedShort() bool {
	return v.flag.flagProvided && v.flag.flagProvidedAsAlias
}

// IsProvidedLong checks if current string flag was provided from stdin but using its long name.
func (v *FlagStringValue) IsProvidedLong() bool {
	return v.flag.flagProvided && !v.flag.flagProvidedAsAlias
}

// GetFlagType returns the associated flag type.
func (v *FlagStringValue) GetFlagType() FlagString {
	return v.flag
}

// FlagStringSliceValue represents a flag value string slice type.
type FlagStringSliceValue struct {
	flag FlagStringSlice
}

// Value unwraps the plain string slice value of the current flag.
func (v *FlagStringSliceValue) Value() []string {
	return v.flag.flagValue.ToStringSlice()
}

// IsProvided checks if current string slice flag was provided from stdin.
func (v *FlagStringSliceValue) IsProvided() bool {
	return v.flag.flagProvided
}

// IsProvidedShort checks if current string slice flag was provided from stdin but using its short name.
func (v *FlagStringSliceValue) IsProvidedShort() bool {
	return v.flag.flagProvided && v.flag.flagProvidedAsAlias
}

// IsProvidedLong checks if current string slice flag was provided from stdin but using its long name.
func (v *FlagStringSliceValue) IsProvidedLong() bool {
	return v.flag.flagProvided && !v.flag.flagProvidedAsAlias
}

// GetFlagType returns the associated flag type.
func (v *FlagStringSliceValue) GetFlagType() FlagStringSlice {
	return v.flag
}

// FlagValues defines list of flag values.
type FlagValues struct {
	flags []Flag
}

// findByKey finds a `Flag` by its string key.
func (v *FlagValues) findByKey(longFlagName string) Flag {
	for _, f := range v.flags {
		switch fl := f.(type) {
		case FlagBool:
			if longFlagName == fl.Name {
				return fl
			}
		case FlagInt:
			if longFlagName == fl.Name {
				return fl
			}
		case FlagString:
			if longFlagName == fl.Name {
				return fl
			}
		case FlagStringSlice:
			if longFlagName == fl.Name {
				return fl
			}
		}
	}
	return nil
}

// GetProvided returns all flags that were provided from stdin only.
func (v *FlagValues) GetProvided() []Flag {
	var flags []Flag
	for _, e := range v.flags {
		switch f := e.(type) {
		case FlagBool:
			if f.flagProvided {
				flags = append(flags, f)
			}
		case FlagInt:
			if f.flagProvided {
				flags = append(flags, f)
			}
		case FlagString:
			if f.flagProvided {
				flags = append(flags, f)
			}
		case FlagStringSlice:
			if f.flagProvided {
				flags = append(flags, f)
			}
		}
	}
	return flags
}

// GetProvidedLong returns all flags that were provided from stdin but using long names only.
func (v *FlagValues) GetProvidedLong() []Flag {
	var flags []Flag
	for _, e := range v.flags {
		switch f := e.(type) {
		case FlagBool:
			if f.flagProvided && !f.flagProvidedAsAlias {
				flags = append(flags, f)
			}
		case FlagInt:
			if f.flagProvided && !f.flagProvidedAsAlias {
				flags = append(flags, f)
			}
		case FlagString:
			if f.flagProvided && !f.flagProvidedAsAlias {
				flags = append(flags, f)
			}
		case FlagStringSlice:
			if f.flagProvided && !f.flagProvidedAsAlias {
				flags = append(flags, f)
			}
		}
	}
	return flags
}

// Any finds a flag value but ignoring its type. The result value is convertible to other supported types.
// Since `AnyValue` is just a `string` alias type, it can be converted easily with `string(AnyValue)`.
func (v *FlagValues) Any(longFlagName string) AnyValue {
	switch f := v.findByKey(longFlagName).(type) {
	case FlagBool:
		return f.flagValue
	case FlagInt:
		return f.flagValue
	case FlagString:
		return f.flagValue
	case FlagStringSlice:
		return f.flagValue
	}
	return AnyValue("")
}

// Bool finds a `bool` flag value. It's type should match with its flag definition type.
func (v *FlagValues) Bool(longFlagName string) *FlagBoolValue {
	switch f := v.findByKey(longFlagName).(type) {
	case FlagBool:
		return &FlagBoolValue{flag: f}
	default:
		t := strings.ReplaceAll(fmt.Sprintf("%T", f), "cline.", "")
		fmt.Printf("error: flag `--%s` value used as `FlagBoolValue` but declared as `%s`.\n", longFlagName, t)
		os.Exit(1)
	}
	return nil
}

// Int finds a `int` flag value. It's type should match with its flag definition type.
func (v *FlagValues) Int(longFlagName string) *FlagIntValue {
	switch f := v.findByKey(longFlagName).(type) {
	case FlagInt:
		return &FlagIntValue{flag: f}
	default:
		t := strings.ReplaceAll(fmt.Sprintf("%T", f), "cline.", "")
		fmt.Printf("error: flag `--%s` value used as `FlagIntValue` but declared as `%s`.\n", longFlagName, t)
		os.Exit(1)
	}
	return nil
}

// String finds a `string` flag value. It's type should match with its flag definition type.
func (v *FlagValues) String(longFlagName string) *FlagStringValue {
	switch f := v.findByKey(longFlagName).(type) {
	case FlagString:
		return &FlagStringValue{flag: f}
	default:
		t := strings.ReplaceAll(fmt.Sprintf("%T", f), "cline.", "")
		fmt.Printf("error: flag `--%s` value used as `FlagStringValue` but declared as `%s`.\n", longFlagName, t)
		os.Exit(1)
	}
	return nil
}

// StringSlice finds a string slice. It's type should match with its flag definition type.
func (v *FlagValues) StringSlice(longFlagName string) *FlagStringSliceValue {
	switch f := v.findByKey(longFlagName).(type) {
	case FlagStringSlice:
		return &FlagStringSliceValue{flag: f}
	default:
		t := strings.ReplaceAll(fmt.Sprintf("%T", f), "cline.", "")
		fmt.Printf("error: flag `--%s` value used as `FlagStringSliceValue` but declared as `%s`.\n", longFlagName, t)
		os.Exit(1)
	}
	return nil
}
