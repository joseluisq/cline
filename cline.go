// Package cline is a fast and lightweight CLI package for Go.
//
package cline

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

// FlagValue represents a `bool`, `int` or `string` input value for a command flag.
type FlagValue string

// Bool converts current flag value to `bool`.
func (v FlagValue) Bool() (bool, error) {
	return strconv.ParseBool(v.String())
}

// Int converts current flag value to `int`.
func (v FlagValue) Int() (int, error) {
	return strconv.Atoi(v.String())
}

// String converts current flag value to `string`.
func (v FlagValue) String() string {
	return string(v)
}

// StringSlice converts current flag value to a string slice.
func (v FlagValue) StringSlice() []string {
	var strs []string
	for _, s := range strings.Split(string(v), ",") {
		strs = append(strs, strings.TrimSpace(s))
	}
	return strs
}

// Flag defines a command flag.
type Flag struct {
	Name             string
	Summary          string
	Value            interface{}
	Shortcuts        []string
	EnvVar           string
	vflag            FlagValue
	vflagAssigned    bool
	vflagEnvAssigned bool
}

// FlagMap defines a hash map of command flags.
type FlagMap struct {
	flags []*Flag
}

func (b *FlagMap) byKey(flagKey string) FlagValue {
	for _, v := range b.flags {
		if flagKey == v.Name {
			return v.vflag
		}
	}
	return FlagValue("")
}

// Bool gets current flag value as `bool`.
func (b *FlagMap) Bool(flagName string) (bool, error) {
	return b.byKey(flagName).Bool()
}

// Int gets current flag value as `int`.
func (b *FlagMap) Int(flagName string) (int, error) {
	return b.byKey(flagName).Int()
}

// String gets current flag value as `string`.
func (b *FlagMap) String(flagName string) string {
	return b.byKey(flagName).String()
}

// StringSlice gets current flag value as a string slice.
func (b *FlagMap) StringSlice(flagName string) []string {
	return b.byKey(flagName).StringSlice()
}

// Cmd defines an application command.
type Cmd struct {
	Name    string
	Summary string
	Flags   []*Flag
	Handler CmdHandler
}

// App defines application settings
type App struct {
	Name     string
	Summary  string
	Flags    []*Flag
	Commands []*Cmd
	Handler  AppHandler
}

// AppContext defines an application context.
type AppContext struct {
	App      *App
	Flags    *FlagMap
	TailArgs []string
}

// AppHandler responds to an application action.
type AppHandler func(*AppContext) error

// CmdContext defines command context.
type CmdContext struct {
	Cmd        *Cmd
	Flags      *FlagMap
	TailArgs   []string
	AppContext *AppContext
}

// CmdHandler responds to a command action.
type CmdHandler func(*CmdContext) error

func findFlagByKey(key string, opts []*Flag) (*Flag, bool) {
	for _, f := range opts {
		// Check for long named flags
		if key == f.Name {
			return f, true
		}

		// Check for short named flags
		for _, s := range f.Shortcuts {
			if key == s {
				return f, true
			}
		}
	}
	return &Flag{}, false
}

// Assign default flag values
func setDefaultFlagValue(flag *Flag) bool {
	switch v := flag.Value.(type) {
	case bool:
		val := FlagValue(strconv.FormatBool(v))
		ev, ok := syscall.Getenv(flag.EnvVar)
		if ok {
			if b, err := FlagValue(ev).Bool(); err == nil {
				val = FlagValue(strconv.FormatBool(b))
			}
		}
		flag.vflagEnvAssigned = ok
		flag.vflag = val
		return true
	case int:
		val := FlagValue(strconv.Itoa(v))
		ev, ok := syscall.Getenv(flag.EnvVar)
		if ok {
			s := FlagValue(ev)
			if _, err := s.Int(); err == nil {
				val = s
			}
		}
		flag.vflagEnvAssigned = ok
		flag.vflag = val
		return true
	case string:
		val := FlagValue(v)
		ev, ok := syscall.Getenv(flag.EnvVar)
		if ok {
			val = FlagValue(ev)
		}
		flag.vflagEnvAssigned = ok
		flag.vflag = val
		return true
	case []string:
		val := FlagValue(strings.Join(v, ","))
		ev, ok := syscall.Getenv(flag.EnvVar)
		if ok {
			val = FlagValue(ev)
		}
		flag.vflagEnvAssigned = ok
		flag.vflag = val
		return true
	case nil:
		flag.vflagEnvAssigned = false
		flag.vflag = FlagValue("")
		return true
	}
	return false
}

// Run executes the application.
func (app *App) Run() error {
	// Commands and flags validation

	// 1. Validate the application flags
	for _, f := range app.Flags {
		fname := strings.ToLower(strings.TrimSpace(f.Name))

		if fname == "" {
			return fmt.Errorf("global flag name defined with an empty value")
		}

		assigned := setDefaultFlagValue(f)

		if !assigned {
			return fmt.Errorf("global flag `%s` has invalid data type value. Use bool, int, string, []string or nil", fname)
		}
	}

	// 2. Validate commands and their flags
	for _, c := range app.Commands {
		cname := strings.ToLower(strings.TrimSpace(c.Name))

		if cname == "" {
			return fmt.Errorf("one command name defined with an empty value")
		}

		for _, f := range c.Flags {
			fname := strings.ToLower(strings.TrimSpace(f.Name))

			if fname == "" {
				return fmt.Errorf("one flag name for command `%s` defined with an empty value", cname)
			}

			assigned := setDefaultFlagValue(f)

			if !assigned {
				return fmt.Errorf("command flag `%s` has invalid data type value. Use bool, int, string, []string or nil", fname)
			}
		}
	}

	// 3. Process commands and flags

	var lastCmd *Cmd
	var lastFlag *Flag
	var tailArgs []string
	var hasCommand = false

	for i := 1; i < len(os.Args); i++ {
		arg := strings.ToLower(strings.TrimSpace(os.Args[i]))

		// Check for no supported arguments (remaining)
		if len(tailArgs) > 0 {
			tailArgs = append(tailArgs, arg)
			continue
		}

		// 3.1. Flags (options)
		if strings.HasPrefix(arg, "-") {
			flagKey := strings.TrimPrefix(strings.TrimPrefix(arg, "-"), "-")

			// Skip unsupported fags
			if strings.HasPrefix(flagKey, "-") {
				tailArgs = append(tailArgs, arg)
				continue
			}

			// Assign flag default values
			var flags []*Flag

			if hasCommand {
				flags = lastCmd.Flags
			} else {
				flags = app.Flags
			}

			// Check for no supported flag only
			flag, found := findFlagByKey(flagKey, flags)

			if !found {
				return fmt.Errorf("argument `%s` is not recognised", arg)
			}

			lastFlag = flag
			continue
		}

		// 3.2. Commands
		// 3.2.1 Check for a valid command (first time)
		if !hasCommand {
			for _, c := range app.Commands {
				if c.Name == arg {
					hasCommand = true
					lastCmd = c
					break
				}
			}

			if hasCommand {
				continue
			}
		}

		// 4. If there is no command found assume it as a tail arg
		if lastFlag == nil {
			tailArgs = append(tailArgs, arg)
			continue
		}

		// 5. Process command flag Values
		if lastFlag.Name != "" {
			if lastFlag.vflagAssigned {
				tailArgs = append(tailArgs, arg)
				continue
			}

			s := FlagValue(arg)
			switch lastFlag.Value.(type) {
			case bool:
				if _, err := s.Bool(); err == nil {
					lastFlag.vflag = s
					lastFlag.vflagAssigned = true
				} else {
					tailArgs = append(tailArgs, arg)
				}
			case int:
				if _, err := s.Int(); err == nil {
					lastFlag.vflag = s
					lastFlag.vflagAssigned = true
				} else {
					tailArgs = append(tailArgs, arg)
				}
			case string, []string:
				lastFlag.vflag = s
				lastFlag.vflagAssigned = true
			default:
				tailArgs = append(tailArgs, arg)
			}
			continue
		}
	}

	// Call command handler
	if hasCommand && lastCmd.Handler != nil {
		return lastCmd.Handler(&CmdContext{
			Cmd: lastCmd,
			Flags: &FlagMap{
				flags: lastCmd.Flags,
			},
			TailArgs: tailArgs,
			AppContext: &AppContext{
				App: app,
				Flags: &FlagMap{
					flags: app.Flags,
				},
			},
		})
	}

	// Call application handler
	if app.Handler != nil {
		return app.Handler(&AppContext{
			App: app,
			Flags: &FlagMap{
				flags: app.Flags,
			},
			TailArgs: tailArgs,
		})
	}

	return nil
}
