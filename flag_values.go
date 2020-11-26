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

// FlagValueMap defines a hash map of command input flags with their values.
type FlagValueMap struct {
	zFlags         []Flag
	zProvidedFlags []string
}

// ProvidedFlags gets a list of flag keys with the provided input flags only.
func (fm *FlagValueMap) ProvidedFlags() []string {
	return fm.zProvidedFlags
}

// findByKey finds a `FlagValue` by a string key.
func (fm *FlagValueMap) findByKey(flagKey string) FlagValue {
	for _, v := range fm.zFlags {
		switch fl := v.(type) {
		case FlagBool:
			if flagKey == fl.Name {
				return fl.zflag
			}
		case FlagInt:
			if flagKey == fl.Name {
				return fl.zflag
			}
		case FlagString:
			if flagKey == fl.Name {
				return fl.zflag
			}
		case FlagStringSlice:
			if flagKey == fl.Name {
				return fl.zflag
			}
		}
	}
	return FlagValue("")
}

// Bool gets current flag value as `bool`.
func (fm *FlagValueMap) Bool(flagName string) (bool, error) {
	return fm.findByKey(flagName).Bool()
}

// Int gets current flag value as `int`.
func (fm *FlagValueMap) Int(flagName string) (int, error) {
	return fm.findByKey(flagName).Int()
}

// String gets current flag value as `string`.
func (fm *FlagValueMap) String(flagName string) string {
	return fm.findByKey(flagName).String()
}

// StringSlice gets current flag value as a string slice.
func (fm *FlagValueMap) StringSlice(flagName string) []string {
	return fm.findByKey(flagName).StringSlice()
}
