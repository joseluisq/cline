package flag

import (
	"fmt"
	"strings"
)

// FlagValues holds all defined flags with their values.
type FlagValues struct {
	Flags []Flag
}

// NewFlagValues creates a new `FlagValues` instance.
func NewFlagValues(flags []Flag) *FlagValues {
	return &FlagValues{
		Flags: flags,
	}
}

// It finds a `Flag` by its string key in the inner list.
func (v *FlagValues) FindByKey(longFlagName string) (flag Flag) {
	longFlagName = strings.TrimSpace(longFlagName)
	if longFlagName == "" {
		return
	}
	for _, fl := range v.Flags {
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
func (v *FlagValues) GetProvidedFlags(providedOnly bool, providedAliasOnly bool) (flags []Flag) {
	if !providedOnly && !providedAliasOnly {
		flags = v.Flags
		return
	}
	for _, fl := range v.Flags {
		switch f := fl.(type) {
		case FlagBool:
			if !f.FlagProvided {
				continue
			}
			if providedOnly {
				flags = append(flags, f)
				continue
			}
			if providedAliasOnly && f.FlagProvidedAsAlias {
				flags = append(flags, f)
				continue
			}
		case FlagInt:
			if !f.FlagProvided {
				continue
			}
			if providedOnly {
				flags = append(flags, f)
				continue
			}
			if providedAliasOnly && f.FlagProvidedAsAlias {
				flags = append(flags, f)
				continue
			}
		case FlagString:
			if !f.FlagProvided {
				continue
			}
			if providedOnly {
				flags = append(flags, f)
				continue
			}
			if providedAliasOnly && f.FlagProvidedAsAlias {
				flags = append(flags, f)
				continue
			}
		case FlagStringSlice:
			if !f.FlagProvided {
				continue
			}
			if providedOnly {
				flags = append(flags, f)
				continue
			}
			if providedAliasOnly && f.FlagProvidedAsAlias {
				flags = append(flags, f)
				continue
			}
		}
	}
	return
}

// GetProvided returns all flags that were provided from stdin.
func (v *FlagValues) GetProvided() []Flag {
	return v.GetProvidedFlags(true, false)
}

// GetProvidedLong returns all flags that were provided from stdin but using long flag names only.
func (v *FlagValues) GetProvidedLong() []Flag {
	return v.GetProvidedFlags(false, false)
}

// GetProvidedShort returns all flags that were provided from stdin but using short flag names (alias) only.
func (v *FlagValues) GetProvidedShort() []Flag {
	return v.GetProvidedFlags(false, true)
}

// Value gets the current flag value but ignoring its type.
// However, the resulted value is convertible into other supported types.
// And since the `Value` type is just an alias of the built-in `string` type,
// it can be easily converted into string like `string(Value)`.
func (v *FlagValues) Value(longFlagName string) Value {
	switch f := v.FindByKey(longFlagName).(type) {
	case FlagBool:
		return f.FlagValue
	case FlagInt:
		return f.FlagValue
	case FlagString:
		return f.FlagValue
	case FlagStringSlice:
		return f.FlagValue
	default:
		return Value("")
	}
}

// Bool gets a `bool` flag value which value type should match
// with its flag definition type, otherwise it returns an error.
func (v *FlagValues) Bool(longFlagName string) (val *ValueBool, err error) {
	switch f := v.FindByKey(longFlagName).(type) {
	case FlagBool:
		val = &ValueBool{Flag: f}
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
func (v *FlagValues) Int(longFlagName string) (val *ValueInt, err error) {
	switch f := v.FindByKey(longFlagName).(type) {
	case FlagInt:
		val = &ValueInt{Flag: f}
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
func (v *FlagValues) String(longFlagName string) (val *ValueString, err error) {
	switch f := v.FindByKey(longFlagName).(type) {
	case FlagString:
		val = &ValueString{Flag: f}
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
func (v *FlagValues) StringSlice(longFlagName string) (val *ValueStringSlice, err error) {
	switch f := v.FindByKey(longFlagName).(type) {
	case FlagStringSlice:
		val = &ValueStringSlice{Flag: f}
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
