package helpers

import (
	"fmt"
	"strings"

	"github.com/joseluisq/cline/app"
	"github.com/joseluisq/cline/flag"
)

// It checks if a list of commands and initialize them if they are valid.
func ValidateCommands(commands []app.Cmd) (cmds []app.Cmd, err error) {
	for _, c := range commands {
		name := strings.TrimSpace(c.Name)
		if name == "" {
			err = fmt.Errorf("one command name has empty value")
			return
		}
		flags, errf := ValidateFlagsAndInit(c.Flags)
		if errf != nil {
			err = errf
			return
		}
		c.Flags = flags
		cmds = append(cmds, c)
	}
	return
}

// It checks a list of flags and initialize them if they are valid.
func ValidateFlagsAndInit(flags []flag.Flag) (vflags []flag.Flag, err error) {
	for _, v := range flags {
		switch f := v.(type) {
		case flag.FlagBool:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("bool flag name has an empty value")
				return
			}
			f.Init()
			vflags = append(vflags, f)

		case flag.FlagInt:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("int flag name has an empty value")
				return
			}
			f.Init()
			vflags = append(vflags, f)

		case flag.FlagString:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("string flag name has an empty value")
				return
			}
			f.Init()
			vflags = append(vflags, f)

		case flag.FlagStringSlice:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("string slice flag name has an empty value")
				return
			}
			f.Init()
			vflags = append(vflags, f)

		default:
			err = fmt.Errorf("one flag has invalid data type value. Use a bool, int, string, []string or nil value")
			return
		}
	}
	return
}

// It finds a flag item with its index in a given flags array by key
// then checks if every flag is a short flag or not.
func FindFlagByKey(key string, flags []flag.Flag) (index int, fl flag.Flag, isAlias bool) {
	for i, v := range flags {
		switch f := v.(type) {
		case flag.FlagBool:
			if IsFlagLong(f.Name, key) {
				return i, f, false
			}
			if IsFlagAlias(key, f.Aliases) {
				return i, f, true
			}
		case flag.FlagInt:
			if IsFlagLong(f.Name, key) {
				return i, f, false
			}
			if IsFlagAlias(key, f.Aliases) {
				return i, f, true
			}
		case flag.FlagString:
			if IsFlagLong(f.Name, key) {
				return i, f, false
			}
			if IsFlagAlias(key, f.Aliases) {
				return i, f, true
			}
		case flag.FlagStringSlice:
			if IsFlagLong(f.Name, key) {
				return i, f, false
			}
			if IsFlagAlias(key, f.Aliases) {
				return i, f, true
			}
		}
	}
	return -1, nil, false
}

// Check for long named flags.
func IsFlagLong(name string, key string) bool {
	return name == key
}

// Check for short named flags (aliases).
func IsFlagAlias(key string, aliases []string) bool {
	for _, s := range aliases {
		if s == key {
			return true
		}
	}
	return false
}
