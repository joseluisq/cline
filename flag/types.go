package flag

import (
	"strconv"
	"strings"
	"syscall"
)

// Flag defines a flag generic type.
type Flag interface{}

// FlagInt defines an `Int` type flag.
type FlagInt struct {
	// Name of the flag containing alphanumeric characters and dashes
	// but without leading dashes, spaces or any kind of special chars.
	Name string
	// An optional summary for the flag.
	Summary string
	// An optional default value for the flag.
	Value int
	// An optional list of flag aliases containing single alphanumeric characters
	// but without dashes, spaces or any special chars.
	Aliases []string
	// An optional environment variable containing uppercase alphanumeric characters
	// and underscores but without dashes, spaces or any kind of special chars.
	EnvVar string

	FlagValue           Value
	FlagAssigned        bool
	FlagProvided        bool
	FlagProvidedAsAlias bool
}

// It sets a default flag value via its associated `Value` prop
// or its environment variable (`EnvVar`) if so.
func (fi *FlagInt) Initialize() {
	val := Value(strconv.Itoa(fi.Value))
	ev, ok := syscall.Getenv(fi.EnvVar)
	if ok {
		s := Value(ev)
		if _, err := s.ToInt(); err == nil {
			val = s
		}
	}
	fi.FlagValue = val
}

// FlagBool defines a `bool` type flag.
type FlagBool struct {
	// Name of the flag containing alphanumeric characters and dashes
	// but without leading dashes, spaces or any kind of special chars.
	Name string
	// An optional summary for the flag.
	Summary string
	// An optional default value for the flag.
	Value bool
	// An optional list of flag aliases containing single alphanumeric characters
	// but without dashes, spaces or any special chars.
	Aliases []string
	// An optional environment variable containing uppercase alphanumeric characters
	// and underscores but without dashes, spaces or any kind of special chars.
	EnvVar string

	FlagValue           Value
	FlagAssigned        bool
	FlagProvided        bool
	FlagProvidedAsAlias bool
}

// It sets a default flag value via its associated `Value` prop
// or its environment variable (`EnvVar`) if so.
func (fb *FlagBool) Initialize() {
	val := Value(strconv.FormatBool(fb.Value))
	ev, ok := syscall.Getenv(fb.EnvVar)
	if ok {
		if b, err := Value(ev).ToBool(); err == nil {
			val = Value(strconv.FormatBool(b))
		}
	}
	fb.FlagValue = val
}

// FlagString defines a `String` type flag.
type FlagString struct {
	// Name of the flag containing alphanumeric characters and dashes
	// but without leading dashes, spaces or any kind of special chars.
	Name string
	// An optional summary for the flag.
	Summary string
	// An optional default value for the flag.
	Value string
	// An optional list of flag aliases containing single alphanumeric characters
	// but without dashes, spaces or any special chars.
	Aliases []string
	// An optional environment variable containing uppercase alphanumeric characters
	// and underscores but without dashes, spaces or any kind of special chars.
	EnvVar string

	FlagValue           Value
	FlagAssigned        bool
	FlagProvided        bool
	FlagProvidedAsAlias bool
}

// It sets a default flag value via its associated `Value` prop
// or its environment variable (`EnvVar`) if so.
func (fs *FlagString) Initialize() {
	val := Value(fs.Value)
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = Value(ev)
	}
	fs.FlagValue = val
}

// FlagStringSlice defines a string slice type flag.
type FlagStringSlice struct {
	// Name of the flag containing alphanumeric characters and dashes
	// but without leading dashes, spaces or any kind of special chars.
	Name string
	// An optional default value for the flag.
	Summary string
	// An optional default value for the flag.
	Value []string
	// An optional list of flag aliases containing single alphanumeric characters
	// but without dashes, spaces or any special chars.
	Aliases []string
	// An optional environment variable containing uppercase alphanumeric characters
	// and underscores but without dashes, spaces or any kind of special chars.
	EnvVar string

	FlagValue           Value
	FlagAssigned        bool
	FlagProvided        bool
	FlagProvidedAsAlias bool
}

// It sets a default flag value via its associated `Value` prop
// or its environment variable (`EnvVar`) if so.
func (fs *FlagStringSlice) Initialize() {
	val := Value(strings.Join(fs.Value, ","))
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = Value(ev)
	}
	fs.FlagValue = val
}
