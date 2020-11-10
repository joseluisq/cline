package cline

import (
	"strconv"
	"syscall"
)

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

// setDefaults default flag values via current the `Value` prop or an environment variable (`EnvVar`).
func (fb *FlagBool) setDefaults() {
	val := FlagValue(strconv.FormatBool(fb.Value))
	ev, ok := syscall.Getenv(fb.EnvVar)
	if ok {
		if b, err := FlagValue(ev).Bool(); err == nil {
			val = FlagValue(strconv.FormatBool(b))
		}
	}
	fb.zflag = val
}
