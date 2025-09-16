// Package flag provides all flag and flag-value types.
package flag

import (
	"strconv"
	"strings"
)

// Value is a generic string type alias which represents
// the raw input value for a command flag which can be cast into other types.
type Value string

// ToBool converts current flag value into `bool`.
func (v Value) ToBool() (bool, error) {
	return strconv.ParseBool(v.ToString())
}

// ToInt converts current flag value into `int`.
func (v Value) ToInt() (int, error) {
	return strconv.Atoi(v.ToString())
}

// ToString converts current flag value into `string`.
func (v Value) ToString() string {
	return string(v)
}

// ToStringSlice converts current flag value into a string slice.
func (v Value) ToStringSlice() []string {
	var strs []string
	for _, s := range strings.Split(string(v), ",") {
		strs = append(strs, strings.TrimSpace(s))
	}
	return strs
}

// ValueBool represents a `bool` type flag value.
type ValueBool struct {
	Flag FlagBool
}

// Value unwraps the plain `bool` value of the current flag.
func (v *ValueBool) Value() (bool, error) {
	return v.Flag.FlagValue.ToBool()
}

// IsProvided checks if current `bool` flag was provided from stdin.
func (v *ValueBool) IsProvided() bool {
	return v.Flag.FlagProvided
}

// IsProvidedShort checks if current `bool` flag was provided from stdin but using its short name.
func (v *ValueBool) IsProvidedShort() bool {
	return v.Flag.FlagProvided && v.Flag.FlagProvidedAsAlias
}

// IsProvidedLong checks if current `bool` flag was provided from stdin but using its long name.
func (v *ValueBool) IsProvidedLong() bool {
	return v.Flag.FlagProvided && !v.Flag.FlagProvidedAsAlias
}

// GetFlagType returns the associated flag type.
func (v *ValueBool) GetFlagType() FlagBool {
	return v.Flag
}

// ValueInt represents an `int` type flag value.
type ValueInt struct {
	Flag FlagInt
}

// Value unwraps the plain `int` value of the current flag.
func (v *ValueInt) Value() (int, error) {
	return v.Flag.FlagValue.ToInt()
}

// IsProvided checks if current `int` flag was provided from stdin.
func (v *ValueInt) IsProvided() bool {
	return v.Flag.FlagProvided
}

// IsProvidedShort checks if current `int` flag was provided from stdin but using its short name.
func (v *ValueInt) IsProvidedShort() bool {
	return v.Flag.FlagProvided && v.Flag.FlagProvidedAsAlias
}

// IsProvidedLong checks if current `int` flag was provided from stdin but using its long name.
func (v *ValueInt) IsProvidedLong() bool {
	return v.Flag.FlagProvided && !v.Flag.FlagProvidedAsAlias
}

// GetFlagType returns the associated flag type.
func (v *ValueInt) GetFlagType() FlagInt {
	return v.Flag
}

// ValueString represents a `string` type flag value.
type ValueString struct {
	Flag FlagString
}

// Value unwraps the plain `string` value of the current flag.
func (v *ValueString) Value() string {
	return v.Flag.FlagValue.ToString()
}

// IsProvided checks if current `string` flag was provided from stdin.
func (v *ValueString) IsProvided() bool {
	return v.Flag.FlagProvided
}

// IsProvidedShort checks if current `string` flag was provided from stdin but using its short name.
func (v *ValueString) IsProvidedShort() bool {
	return v.Flag.FlagProvided && v.Flag.FlagProvidedAsAlias
}

// IsProvidedLong checks if current `string` flag was provided from stdin but using its long name.
func (v *ValueString) IsProvidedLong() bool {
	return v.Flag.FlagProvided && !v.Flag.FlagProvidedAsAlias
}

// GetFlagType returns the associated flag type.
func (v *ValueString) GetFlagType() FlagString {
	return v.Flag
}

// ValueStringSlice represents a string slice type flag value.
type ValueStringSlice struct {
	Flag FlagStringSlice
}

// Value unwraps the plain string slice value of the current flag.
func (v *ValueStringSlice) Value() []string {
	return v.Flag.FlagValue.ToStringSlice()
}

// IsProvided checks if current string slice flag was provided from stdin.
func (v *ValueStringSlice) IsProvided() bool {
	return v.Flag.FlagProvided
}

// IsProvidedShort checks if current string slice flag was provided from stdin but using its short name.
func (v *ValueStringSlice) IsProvidedShort() bool {
	return v.Flag.FlagProvided && v.Flag.FlagProvidedAsAlias
}

// IsProvidedLong checks if current string slice flag was provided from stdin but using its long name.
func (v *ValueStringSlice) IsProvidedLong() bool {
	return v.Flag.FlagProvided && !v.Flag.FlagProvidedAsAlias
}

// GetFlagType returns the associated flag type.
func (v *ValueStringSlice) GetFlagType() FlagStringSlice {
	return v.Flag
}
