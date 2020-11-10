package cline

import (
	"fmt"
	"strings"
)

// checkAndInitFlags checks for a flag and initialize it
func checkAndInitFlags(flags []Flag) ([]Flag, error) {
	var sFlags []Flag
	for _, v := range flags {
		switch f := v.(type) {
		case FlagBool:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.setDefaultValue()
			sFlags = append(sFlags, f)
		case FlagInt:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.setDefaultValue()
			sFlags = append(sFlags, f)
		case FlagString:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.setDefaultValue()
			sFlags = append(sFlags, f)
		case FlagStringSlice:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.setDefaultValue()
			sFlags = append(sFlags, f)
		default:
			return nil, fmt.Errorf("global flag has invalid data type value. Use bool, int, string, []string or nil")
		}
	}
	return sFlags, nil
}

// findFlagByKey finds a flag item in a flag's array by key.
func findFlagByKey(key string, flags []Flag) Flag {
	for _, f := range flags {
		switch fl := f.(type) {
		case FlagBool:
			// Check for long named flags
			if key == fl.Name {
				return fl
			}
			// Check for short named flags
			for _, s := range fl.Aliases {
				if key == s {
					return fl
				}
			}
		case FlagInt:
			// Check for long named flags
			if key == fl.Name {
				return fl
			}
			// Check for short named flags
			for _, s := range fl.Aliases {
				if key == s {
					return fl
				}
			}
		case FlagString:
			// Check for long named flags
			if key == fl.Name {
				return fl
			}
			// Check for short named flags
			for _, s := range fl.Aliases {
				if key == s {
					return fl
				}
			}
		case FlagStringSlice:
			// Check for long named flags
			if key == fl.Name {
				return fl
			}
			// Check for short named flags
			for _, s := range fl.Aliases {
				if key == s {
					return fl
				}
			}
		}
	}
	return nil
}
