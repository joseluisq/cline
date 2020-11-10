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

// setDefaultValue default flag values.
func (fs *FlagStringSlice) setDefaultValue() {
	val := FlagValue(strings.Join(fs.Value, ","))
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = FlagValue(ev)
	}
	fs.zflag = val
}
