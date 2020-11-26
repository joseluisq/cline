package cline

import (
	"strconv"
	"strings"
)

// FlagValue represents a `bool`, `int`, `string` or `string slice` input value for a command flag.
type FlagValue string

// Bool converts current flag value into a `bool`.
func (v FlagValue) Bool() (bool, error) {
	return strconv.ParseBool(v.String())
}

// Int converts current flag value into an `int`.
func (v FlagValue) Int() (int, error) {
	return strconv.Atoi(v.String())
}

// String converts current flag value into a `string`.
func (v FlagValue) String() string {
	return string(v)
}

// StringSlice converts current flag value into a string slice.
func (v FlagValue) StringSlice() []string {
	var strs []string
	for _, s := range strings.Split(string(v), ",") {
		strs = append(strs, strings.TrimSpace(s))
	}
	return strs
}

// FlagProvided defines a provided input flag (passed from stdin only) and if it was an alias or not.
type FlagProvided struct {
	Name    string
	IsAlias bool
}

// isFlagProvided checks for a provided flag and check alias named ones as well.
func (fm *FlagMapping) isFlagProvided(flagName string, checkAlias bool, aliasValue bool) bool {
	for _, p := range fm.zFlagsProvided {
		if flagName == p.Name {
			if checkAlias {
				return aliasValue == p.IsAlias
			}
			return true
		}
	}
	return false
}

// GetProvidedFlags gets a list of provided input flags (only those passed from stdin).
func (fm *FlagMapping) GetProvidedFlags() []FlagProvided {
	return fm.zFlagsProvided
}

// IsProvidedFlag checks if a flag was provided (only those passed from stdin).
// `flagName` specifies the long flag name.
func (fm *FlagMapping) IsProvidedFlag(flagName string) bool {
	return fm.isFlagProvided(flagName, false, false)
}

// IsLongProvidedFlag checks if a flag was provided (only those passed from stdin) using its long name.
// `flagName` specifies the long flag name.
func (fm *FlagMapping) IsLongProvidedFlag(flagName string) bool {
	return fm.isFlagProvided(flagName, true, false)
}

// IsShortProvidedFlag checks if a flag was provided (only those passed from stdin) using its short name.
// `flagName` specifies the long flag name.
func (fm *FlagMapping) IsShortProvidedFlag(flagName string) bool {
	return fm.isFlagProvided(flagName, true, true)
}

// findByKey finds a `FlagValue` by a string key.
// `flagName` specifies the long flag name.
func (fm *FlagMapping) findByKey(flagName string) FlagValue {
	for _, v := range fm.zFlags {
		switch fl := v.(type) {
		case FlagBool:
			if flagName == fl.Name {
				return fl.zflag
			}
		case FlagInt:
			if flagName == fl.Name {
				return fl.zflag
			}
		case FlagString:
			if flagName == fl.Name {
				return fl.zflag
			}
		case FlagStringSlice:
			if flagName == fl.Name {
				return fl.zflag
			}
		}
	}
	return FlagValue("")
}

// FlagMapping defines a hash map of command input flags with their values.
type FlagMapping struct {
	zFlags         []Flag
	zFlagsProvided []FlagProvided
}

// Bool gets current flag value as `bool`.
// `flagName` specifies the long flag name.
func (fm *FlagMapping) Bool(flagName string) (bool, error) {
	return fm.findByKey(flagName).Bool()
}

// Int gets current flag value as `int`.
// `flagName` specifies the long flag name.
func (fm *FlagMapping) Int(flagName string) (int, error) {
	return fm.findByKey(flagName).Int()
}

// String gets current flag value as `string`.
// `flagName` specifies the long flag name.
func (fm *FlagMapping) String(flagName string) string {
	return fm.findByKey(flagName).String()
}

// StringSlice gets current flag value as a string slice.
// `flagName` specifies the long flag name.
func (fm *FlagMapping) StringSlice(flagName string) []string {
	return fm.findByKey(flagName).StringSlice()
}
