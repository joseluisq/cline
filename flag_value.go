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
