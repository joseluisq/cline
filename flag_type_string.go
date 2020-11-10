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

// setDefaults default flag values via current the `Value` prop or an environment variable (`EnvVar`).
func (fs *FlagString) setDefaults() {
	val := FlagValue(fs.Value)
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = FlagValue(ev)
	}
	fs.zflag = val
}
