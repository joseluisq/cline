package cline

import (
	"fmt"
	"strings"
)

// validateCommands checks if a command is valid and initialize it
func validateCommands(commands []Cmd) ([]Cmd, error) {
	var cmds []Cmd
	for _, c := range commands {
		name := strings.TrimSpace(c.Name)
		if name == "" {
			return nil, fmt.Errorf("command name has empty value")
		}
		vflags, err := validateFlagsAndInit(c.Flags)
		if err != nil {
			return nil, err
		}
		c.Flags = vflags
		cmds = append(cmds, c)
	}
	return cmds, nil
}

// validateFlagsAndInit validates flags and initialize them
func validateFlagsAndInit(flags []Flag) ([]Flag, error) {
	var vFlags []Flag
	for _, v := range flags {
		switch f := v.(type) {
		case FlagBool:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("flag name has empty value")
			}
			f.initialize()
			vFlags = append(vFlags, f)
		case FlagInt:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("flag name has empty value")
			}
			f.initialize()
			vFlags = append(vFlags, f)
		case FlagString:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("flag name has empty value")
			}
			f.initialize()
			vFlags = append(vFlags, f)
		case FlagStringSlice:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("flag name has empty value")
			}
			f.initialize()
			vFlags = append(vFlags, f)
		default:
			return nil, fmt.Errorf("flag has invalid data type value. Use bool, int, string, []string or nil")
		}
	}
	return vFlags, nil
}

// findFlagByKey finds a flag item with its index in a flag's array by key.
// It also checks if was a short flag or not.
func findFlagByKey(key string, flags []Flag) (int, Flag, bool) {
	var short bool = false
	for i, v := range flags {
		switch f := v.(type) {
		case FlagBool:
			// Check for long named flags
			if key == f.Name {
				return i, f, short
			}
			// Check for short named flags
			for _, s := range f.Aliases {
				if key == s {
					short = true
					return i, f, short
				}
			}
		case FlagInt:
			// Check for long named flags
			if key == f.Name {
				return i, f, short
			}
			// Check for short named flags
			for _, s := range f.Aliases {
				if key == s {
					short = true
					return i, f, short
				}
			}
		case FlagString:
			// Check for long named flags
			if key == f.Name {
				return i, f, short
			}
			// Check for short named flags
			for _, s := range f.Aliases {
				if key == s {
					short = true
					return i, f, short
				}
			}
		case FlagStringSlice:
			// Check for long named flags
			if key == f.Name {
				return i, f, short
			}
			// Check for short named flags
			for _, s := range f.Aliases {
				if key == s {
					short = true
					return i, f, short
				}
			}
		}
	}
	return -1, nil, short
}
