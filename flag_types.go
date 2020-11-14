package cline

import (
	"strconv"
	"strings"
	"syscall"
)

// FlagInt defines a flag with `Int` type.
type FlagInt struct {
	Name          string
	Summary       string
	Value         int
	Aliases       []string
	EnvVar        string
	zflag         FlagValue
	zflagAssigned bool
}

// initialize sets a default flag value via associated `Value` prop or an environment variable (`EnvVar`).
func (fi *FlagInt) initialize() {
	val := FlagValue(strconv.Itoa(fi.Value))
	ev, ok := syscall.Getenv(fi.EnvVar)
	if ok {
		s := FlagValue(ev)
		if _, err := s.Int(); err == nil {
			val = s
		}
	}
	fi.zflag = val
}

// FlagBool defines a flag with `bool` type.
type FlagBool struct {
	Name          string
	Summary       string
	Value         bool
	Aliases       []string
	EnvVar        string
	zflag         FlagValue
	zflagAssigned bool
}

// initialize sets a default flag value via associated `Value` prop or an environment variable (`EnvVar`).
func (fb *FlagBool) initialize() {
	val := FlagValue(strconv.FormatBool(fb.Value))
	ev, ok := syscall.Getenv(fb.EnvVar)
	if ok {
		if b, err := FlagValue(ev).Bool(); err == nil {
			val = FlagValue(strconv.FormatBool(b))
		}
	}
	fb.zflag = val
}

// FlagString defines a flag with `String` type.
type FlagString struct {
	Name          string
	Summary       string
	Value         string
	Aliases       []string
	EnvVar        string
	zflag         FlagValue
	zflagAssigned bool
}

// initialize sets a default flag value via associated `Value` prop or an environment variable (`EnvVar`).
func (fs *FlagString) initialize() {
	val := FlagValue(fs.Value)
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = FlagValue(ev)
	}
	fs.zflag = val
}

// FlagStringSlice defines a flag with string slice type.
type FlagStringSlice struct {
	Name          string
	Summary       string
	Value         []string
	Aliases       []string
	EnvVar        string
	zflag         FlagValue
	zflagAssigned bool
}

// initialize sets a default flag value via associated `Value` prop or an environment variable (`EnvVar`).
func (fs *FlagStringSlice) initialize() {
	val := FlagValue(strings.Join(fs.Value, ","))
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = FlagValue(ev)
	}
	fs.zflag = val
}
