// Package helpers provides utility functions for working with flag and command types.
package helpers

import (
	"fmt"
	"slices"
	"strings"

	"github.com/joseluisq/cline/app"
	"github.com/joseluisq/cline/flag"
)

// ValidateCommands checks if a list of commands and initialize them if they are valid.
func ValidateCommands(commands []app.Cmd) (cmds []app.Cmd, err error) {
	for _, c := range commands {
		name := strings.TrimSpace(c.Name)
		if name == "" {
			err = fmt.Errorf("error: command name cannot be empty")
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

// ValidateFlagsAndInit checks a list of flags and initialize them if they are valid.
func ValidateFlagsAndInit(flags []flag.Flag) (vflags []flag.Flag, err error) {
	for _, v := range flags {
		if v == nil {
			err = fmt.Errorf("error: flag list contains a nil value")
			return
		}
		switch f := v.(type) {
		case flag.FlagBool:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("error: bool flag name cannot be empty")
				return
			}
			f.Init()
			vflags = append(vflags, f)

		case flag.FlagInt:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("error: int flag name cannot be empty")
				return
			}
			f.Init()
			vflags = append(vflags, f)

		case flag.FlagString:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("error: string flag name cannot be empty")
				return
			}
			f.Init()
			vflags = append(vflags, f)

		case flag.FlagStringSlice:
			if name := strings.ToLower(strings.TrimSpace(f.Name)); name == "" {
				err = fmt.Errorf("error: string slice flag name cannot be empty")
				return
			}
			f.Init()
			vflags = append(vflags, f)

		default:
			err = fmt.Errorf("error: invalid data type for flag or flag pointer (%T). Use a FlagBool, FlagInt, FlagString, FlagStringSlice or nil value instead", v)
			return
		}
	}
	return
}

// FindFlagByKey finds a flag item with its index in a given flags array by key
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

// IsFlagLong checks for long named flags.
func IsFlagLong(name string, key string) bool {
	return name == key
}

// IsFlagAlias checks for short named flags (aliases).
func IsFlagAlias(key string, aliases []string) bool {
	return slices.Contains(aliases, key)
}

type FlagInfo struct {
	Flag  flag.Flag
	Index int
}

// BuildFlagMap creates a map of flags for quick lookups.
func BuildFlagMap(flags []flag.Flag) map[string]FlagInfo {
	// Pre-allocate for names and aliases
	flagMap := make(map[string]FlagInfo, len(flags)*2)
	for i, f := range flags {
		info := FlagInfo{Flag: f, Index: i}
		switch ft := f.(type) {
		case flag.FlagBool:
			flagMap[ft.Name] = info
			for _, alias := range ft.Aliases {
				flagMap[alias] = info
			}
		case flag.FlagInt:
			flagMap[ft.Name] = info
			for _, alias := range ft.Aliases {
				flagMap[alias] = info
			}
		case flag.FlagString:
			flagMap[ft.Name] = info
			for _, alias := range ft.Aliases {
				flagMap[alias] = info
			}
		case flag.FlagStringSlice:
			flagMap[ft.Name] = info
			for _, alias := range ft.Aliases {
				flagMap[alias] = info
			}
		}
	}
	return flagMap
}
