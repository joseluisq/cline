package cline

import (
	"fmt"
	"strings"
)

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
			f.setDefaults()
			sFlags = append(sFlags, f)
		case FlagInt:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.setDefaults()
			sFlags = append(sFlags, f)
		case FlagString:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.setDefaults()
			sFlags = append(sFlags, f)
		case FlagStringSlice:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.setDefaults()
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
