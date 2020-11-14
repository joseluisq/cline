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
		cflags, err := validateAndInitFlags(c.Flags)
		if err != nil {
			return nil, err
		}
		c.Flags = cflags
		cmds = append(cmds, c)
	}
	return cmds, nil
}

// validateAndInitFlags checks for a flag and initialize it
func validateAndInitFlags(flags []Flag) ([]Flag, error) {
	var sFlags []Flag
	for _, v := range flags {
		switch f := v.(type) {
		case FlagBool:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.initialize()
			sFlags = append(sFlags, f)
		case FlagInt:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.initialize()
			sFlags = append(sFlags, f)
		case FlagString:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.initialize()
			sFlags = append(sFlags, f)
		case FlagStringSlice:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.initialize()
			sFlags = append(sFlags, f)
		default:
			return nil, fmt.Errorf("global flag has invalid data type value. Use bool, int, string, []string or nil")
		}
	}
	return sFlags, nil
}

// findFlagByKey finds a flag item in a flag's array by key.
func findFlagByKey(key string, flags []Flag) (int, Flag) {
	for i, f := range flags {
		switch fl := f.(type) {
		case FlagBool:
			// Check for long named flags
			if key == fl.Name {
				return i, fl
			}
			// Check for short named flags
			for _, s := range fl.Aliases {
				if key == s {
					return i, fl
				}
			}
		case FlagInt:
			// Check for long named flags
			if key == fl.Name {
				return i, fl
			}
			// Check for short named flags
			for _, s := range fl.Aliases {
				if key == s {
					return i, fl
				}
			}
		case FlagString:
			// Check for long named flags
			if key == fl.Name {
				return i, fl
			}
			// Check for short named flags
			for _, s := range fl.Aliases {
				if key == s {
					return i, fl
				}
			}
		case FlagStringSlice:
			// Check for long named flags
			if key == fl.Name {
				return i, fl
			}
			// Check for short named flags
			for _, s := range fl.Aliases {
				if key == s {
					return i, fl
				}
			}
		}
	}
	return -1, nil
}
