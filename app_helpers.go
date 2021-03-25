package cline

import (
	"fmt"
	"strings"
)

// It checks if a list of commands and initialize them if they are valid.
func validateCommands(commands []Cmd) (cmds []Cmd, err error) {
	for _, c := range commands {
		name := strings.TrimSpace(c.Name)
		if name == "" {
			err = fmt.Errorf("one command name has empty value")
			return
		}
		flags, errf := validateFlagsAndInit(c.Flags)
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
func validateFlagsAndInit(flags []Flag) (vflags []Flag, err error) {
	for _, v := range flags {
		switch f := v.(type) {
		case FlagBool:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("bool flag name has an empty value")
				return
			}
			f.initialize()
			vflags = append(vflags, f)

		case FlagInt:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("int flag name has an empty value")
				return
			}
			f.initialize()
			vflags = append(vflags, f)

		case FlagString:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("string flag name has an empty value")
				return
			}
			f.initialize()
			vflags = append(vflags, f)

		case FlagStringSlice:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("string slice flag name has an empty value")
				return
			}
			f.initialize()
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
func findFlagByKey(key string, flags []Flag) (int, Flag, bool) {
	for i, v := range flags {
		switch f := v.(type) {
		case FlagBool:
			// Check for long named flags
			if f.Name == key {
				return i, f, false
			}
			// Check for short named flags
			for _, s := range f.Aliases {
				if s == key {
					return i, f, true
				}
			}
		case FlagInt:
			// Check for long named flags
			if f.Name == key {
				return i, f, false
			}
			// Check for short named flags
			for _, s := range f.Aliases {
				if s == key {
					return i, f, true
				}
			}
		case FlagString:
			// Check for long named flags
			if f.Name == key {
				return i, f, false
			}
			// Check for short named flags
			for _, s := range f.Aliases {
				if s == key {
					return i, f, true
				}
			}
		case FlagStringSlice:
			// Check for long named flags
			if f.Name == key {
				return i, f, false
			}
			// Check for short named flags
			for _, s := range f.Aliases {
				if s == key {
					return i, f, true
				}
			}
		}
	}
	return -1, nil, false
}
