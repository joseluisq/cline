package cline

import (
	"strconv"
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

// setDefaultValue default flag values.
func (fi *FlagInt) setDefaultValue() {
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
