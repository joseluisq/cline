package cline

import (
	"strconv"
	"strings"
	"syscall"
)

// FlagInt defines a flag with `Int` type.
type FlagInt struct {
	Name                string
	Summary             string
	Value               int
	Aliases             []string
	EnvVar              string
	flagValue           AnyValue
	flagAssigned        bool
	flagProvided        bool
	flagProvidedAsAlias bool
}

// initialize sets a default flag value via associated `Value` prop or an environment variable (`EnvVar`).
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

// FlagBool defines a flag with `bool` type.
type FlagBool struct {
	Name                string
	Summary             string
	Value               bool
	Aliases             []string
	EnvVar              string
	flagValue           AnyValue
	flagAssigned        bool
	flagProvided        bool
	flagProvidedAsAlias bool
}

// initialize sets a default flag value via associated `Value` prop or an environment variable (`EnvVar`).
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

// FlagString defines a flag with `String` type.
type FlagString struct {
	Name                string
	Summary             string
	Value               string
	Aliases             []string
	EnvVar              string
	flagValue           AnyValue
	flagAssigned        bool
	flagProvided        bool
	flagProvidedAsAlias bool
}

// initialize sets a default flag value via associated `Value` prop or an environment variable (`EnvVar`).
func (fs *FlagString) initialize() {
	val := AnyValue(fs.Value)
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = AnyValue(ev)
	}
	fs.flagValue = val
}

// FlagStringSlice defines a flag with string slice type.
type FlagStringSlice struct {
	Name                string
	Summary             string
	Value               []string
	Aliases             []string
	EnvVar              string
	flagValue           AnyValue
	flagAssigned        bool
	flagProvided        bool
	flagProvidedAsAlias bool
}

// initialize sets a default flag value via associated `Value` prop or an environment variable (`EnvVar`).
func (fs *FlagStringSlice) initialize() {
	val := AnyValue(strings.Join(fs.Value, ","))
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = AnyValue(ev)
	}
	fs.flagValue = val
}
