package cline

import (
	"strconv"
	"strings"
	"syscall"
)

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

	flagValue           AnyValue
	flagAssigned        bool
	flagProvided        bool
	flagProvidedAsAlias bool
}

// It sets a default flag value via its associated `Value` prop
// or its environment variable (`EnvVar`) if so.
func (fi *FlagInt) initialize() {
	val := AnyValue(strconv.Itoa(fi.Value))
	ev, ok := syscall.Getenv(fi.EnvVar)
	if ok {
		s := AnyValue(ev)
		if _, err := s.ToInt(); err == nil {
			val = s
		}
	}
	fi.flagValue = val
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

	flagValue           AnyValue
	flagAssigned        bool
	flagProvided        bool
	flagProvidedAsAlias bool
}

// It sets a default flag value via its associated `Value` prop
// or its environment variable (`EnvVar`) if so.
func (fb *FlagBool) initialize() {
	val := AnyValue(strconv.FormatBool(fb.Value))
	ev, ok := syscall.Getenv(fb.EnvVar)
	if ok {
		if b, err := AnyValue(ev).ToBool(); err == nil {
			val = AnyValue(strconv.FormatBool(b))
		}
	}
	fb.flagValue = val
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

	flagValue           AnyValue
	flagAssigned        bool
	flagProvided        bool
	flagProvidedAsAlias bool
}

// It sets a default flag value via its associated `Value` prop
// or its environment variable (`EnvVar`) if so.
func (fs *FlagString) initialize() {
	val := AnyValue(fs.Value)
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = AnyValue(ev)
	}
	fs.flagValue = val
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

	flagValue           AnyValue
	flagAssigned        bool
	flagProvided        bool
	flagProvidedAsAlias bool
}

// It sets a default flag value via its associated `Value` prop
// or its environment variable (`EnvVar`) if so.
func (fs *FlagStringSlice) initialize() {
	val := AnyValue(strings.Join(fs.Value, ","))
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = AnyValue(ev)
	}
	fs.flagValue = val
}
