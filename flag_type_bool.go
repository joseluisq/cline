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

// setDefaultValue default flag values.
func (fb *FlagBool) setDefaultValue() {
	val := FlagValue(strconv.FormatBool(fb.Value))
	ev, ok := syscall.Getenv(fb.EnvVar)
	if ok {
		if b, err := FlagValue(ev).Bool(); err == nil {
			val = FlagValue(strconv.FormatBool(b))
		}
	}
	fb.zflag = val
}
