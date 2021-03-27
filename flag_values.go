package cline

import (
	"fmt"
	"strconv"
	"strings"
)

// AnyValue is an string type alias which represents
// an input value for a command flag.
type AnyValue string

// ToBool converts current flag value into `bool`.
func (v AnyValue) ToBool() (bool, error) {
	return strconv.ParseBool(v.ToString())
}

// ToInt converts current flag value into `int`.
func (v AnyValue) ToInt() (int, error) {
	return strconv.Atoi(v.ToString())
}

// ToString converts current flag value into `string`.
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

// FlagBoolValue represents a `bool` type flag value.
type FlagBoolValue struct {
	flag FlagBool
}

// Value unwraps the plain `bool` value of the current flag.
func (v *FlagBoolValue) Value() (bool, error) {
	return v.flag.flagValue.ToBool()
}

// IsProvided checks if current `bool` flag was provided from stdin.
func (v *FlagBoolValue) IsProvided() bool {
	return v.flag.flagProvided
}

// IsProvidedShort checks if current `bool` flag was provided from stdin but using its short name.
func (v *FlagBoolValue) IsProvidedShort() bool {
	return v.flag.flagProvided && v.flag.flagProvidedAsAlias
}

// IsProvidedLong checks if current `bool` flag was provided from stdin but using its long name.
func (v *FlagBoolValue) IsProvidedLong() bool {
	return v.flag.flagProvided && !v.flag.flagProvidedAsAlias
}

// GetFlagType returns the associated flag type.
func (v *FlagBoolValue) GetFlagType() FlagBool {
	return v.flag
}

// FlagIntValue represents an `int` type flag value.
type FlagIntValue struct {
	flag FlagInt
}

// Value unwraps the plain `int` value of the current flag.
func (v *FlagIntValue) Value() (int, error) {
	return v.flag.flagValue.ToInt()
}

// IsProvided checks if current `int` flag was provided from stdin.
func (v *FlagIntValue) IsProvided() bool {
	return v.flag.flagProvided
}

// IsProvidedShort checks if current `int` flag was provided from stdin but using its short name.
func (v *FlagIntValue) IsProvidedShort() bool {
	return v.flag.flagProvided && v.flag.flagProvidedAsAlias
}

// IsProvidedLong checks if current `int` flag was provided from stdin but using its long name.
func (v *FlagIntValue) IsProvidedLong() bool {
	return v.flag.flagProvided && !v.flag.flagProvidedAsAlias
}

// GetFlagType returns the associated flag type.
func (v *FlagIntValue) GetFlagType() FlagInt {
	return v.flag
}

// FlagStringValue represents a `string` type flag value.
type FlagStringValue struct {
	flag FlagString
}

// Value unwraps the plain `string` value of the current flag.
func (v *FlagStringValue) Value() string {
	return v.flag.flagValue.ToString()
}

// IsProvided checks if current `string` flag was provided from stdin.
func (v *FlagStringValue) IsProvided() bool {
	return v.flag.flagProvided
}

// IsProvidedShort checks if current `string` flag was provided from stdin but using its short name.
func (v *FlagStringValue) IsProvidedShort() bool {
	return v.flag.flagProvided && v.flag.flagProvidedAsAlias
}

// IsProvidedLong checks if current `string` flag was provided from stdin but using its long name.
func (v *FlagStringValue) IsProvidedLong() bool {
	return v.flag.flagProvided && !v.flag.flagProvidedAsAlias
}

// GetFlagType returns the associated flag type.
func (v *FlagStringValue) GetFlagType() FlagString {
	return v.flag
}

// FlagStringSliceValue represents a string slice type flag value.
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

// It finds a `Flag` by its string key in the inner list.
func (v *FlagValues) findByKey(longFlagName string) (flag Flag) {
	longFlagName = strings.TrimSpace(longFlagName)
	if longFlagName == "" {
		return
	}
	for _, fl := range v.flags {
		switch f := fl.(type) {
		case FlagBool:
			if f.Name == longFlagName {
				flag = f
				return
			}
		case FlagInt:
			if f.Name == longFlagName {
				flag = f
				return
			}
		case FlagString:
			if f.Name == longFlagName {
				flag = f
				return
			}
		case FlagStringSlice:
			if f.Name == longFlagName {
				flag = f
				return
			}
		}
	}
	return
}

// It returns provided flags by specified filters.
func (v *FlagValues) getProvidedFlags(providedOnly bool, providedAliasOnly bool) (flags []Flag) {
	if !providedOnly && !providedAliasOnly {
		flags = v.flags
		return
	}
	for _, fl := range v.flags {
		switch f := fl.(type) {
		case FlagBool:
			if !f.flagProvided {
				continue
			}
			if providedOnly {
				flags = append(flags, f)
				continue
			}
			if providedAliasOnly && f.flagProvidedAsAlias {
				flags = append(flags, f)
				continue
			}
		case FlagInt:
			if !f.flagProvided {
				continue
			}
			if providedOnly {
				flags = append(flags, f)
				continue
			}
			if providedAliasOnly && f.flagProvidedAsAlias {
				flags = append(flags, f)
				continue
			}
		case FlagString:
			if !f.flagProvided {
				continue
			}
			if providedOnly {
				flags = append(flags, f)
				continue
			}
			if providedAliasOnly && f.flagProvidedAsAlias {
				flags = append(flags, f)
				continue
			}
		case FlagStringSlice:
			if !f.flagProvided {
				continue
			}
			if providedOnly {
				flags = append(flags, f)
				continue
			}
			if providedAliasOnly && f.flagProvidedAsAlias {
				flags = append(flags, f)
				continue
			}
		}
	}
	return
}

// GetProvided returns all flags that were provided from stdin only.
func (v *FlagValues) GetProvided() []Flag {
	return v.getProvidedFlags(true, false)
}

// GetProvidedLong returns all flags that were provided from stdin but using long names only.
func (v *FlagValues) GetProvidedLong() []Flag {
	return v.getProvidedFlags(false, false)
}

// GetProvidedShort returns all flags that were provided from stdin but using short names (alias) only.
func (v *FlagValues) GetProvidedShort() []Flag {
	return v.getProvidedFlags(false, true)
}

// Any gets the current flag value but ignoring its type.
// However, the resulted value is convertible into other supported types.
// And since the `AnyValue` is just an alias of built-in `string` type
// it can be easily converted too into string like `string(AnyValue)`.
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

// Bool gets a `bool` flag value which value type should match
// with its flag definition type, otherwise it returns an error.
func (v *FlagValues) Bool(longFlagName string) (val *FlagBoolValue, err error) {
	switch f := v.findByKey(longFlagName).(type) {
	case FlagBool:
		val = &FlagBoolValue{flag: f}
		return
	default:
		t := strings.ReplaceAll(fmt.Sprintf("%T", f), "cline.", "")
		err = fmt.Errorf(
			"flag `--%s` value used as `FlagBoolValue` but declared as `%s`",
			longFlagName,
			t,
		)
		return
	}
}

// Int finds a `int` flag value which value type should match
// with its flag definition type, otherwise it returns an error.
func (v *FlagValues) Int(longFlagName string) (val *FlagIntValue, err error) {
	switch f := v.findByKey(longFlagName).(type) {
	case FlagInt:
		val = &FlagIntValue{flag: f}
		return
	default:
		t := strings.ReplaceAll(fmt.Sprintf("%T", f), "cline.", "")
		err = fmt.Errorf(
			"flag `--%s` value used as `FlagIntValue` but declared as `%s`",
			longFlagName,
			t,
		)
		return
	}
}

// String finds a `string` flag value which value type should match
// with its flag definition type, otherwise it returns an error.
func (v *FlagValues) String(longFlagName string) (val *FlagStringValue, err error) {
	switch f := v.findByKey(longFlagName).(type) {
	case FlagString:
		val = &FlagStringValue{flag: f}
		return
	default:
		t := strings.ReplaceAll(fmt.Sprintf("%T", f), "cline.", "")
		err = fmt.Errorf(
			"flag `--%s` value used as `FlagStringValue` but declared as `%s`",
			longFlagName,
			t,
		)
		return
	}
}

// StringSlice finds a string slice which value type should match
// with its flag definition type, otherwise it returns an error.
func (v *FlagValues) StringSlice(longFlagName string) (val *FlagStringSliceValue, err error) {
	switch f := v.findByKey(longFlagName).(type) {
	case FlagStringSlice:
		val = &FlagStringSliceValue{flag: f}
		return
	default:
		t := strings.ReplaceAll(fmt.Sprintf("%T", f), "cline.", "")
		err = fmt.Errorf(
			"flag `--%s` value used as `FlagStringSliceValue` but declared as `%s`",
			longFlagName,
			t,
		)
		return
	}
}
