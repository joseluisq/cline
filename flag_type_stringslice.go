package cline

import (
	"strings"
	"syscall"
)

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

// setDefaults default flag values via current the `Value` prop or an environment variable (`EnvVar`).
func (fs *FlagStringSlice) setDefaults() {
	val := FlagValue(strings.Join(fs.Value, ","))
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = FlagValue(ev)
	}
	fs.zflag = val
}
