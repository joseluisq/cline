package cline

import "syscall"

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

// setDefaultValue default flag values.
func (fs *FlagString) setDefaultValue() {
	val := FlagValue(fs.Value)
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = FlagValue(ev)
	}
	fs.zflag = val
}
