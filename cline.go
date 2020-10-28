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

// Flag defines a flag generic type.
type Flag interface{}

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

// Cmd defines an application command.
type Cmd struct {
	Name    string
	Summary string
	Flags   []Flag
	Handler CmdHandler
}

// CmdContext defines command context.
type CmdContext struct {
	Cmd        *Cmd
	Flags      *FlagMap
	TailArgs   *[]string
	AppContext *AppContext
}

// CmdHandler responds to a command action.
type CmdHandler func(*CmdContext) error

// App defines application settings
type App struct {
	Name     string
	Summary  string
	Flags    []Flag
	Commands []Cmd
	Handler  AppHandler
}

// AppContext defines an application context.
type AppContext struct {
	App      *App
	Flags    *FlagMap
	TailArgs *[]string
}

// AppHandler responds to an application action.
type AppHandler func(*AppContext) error

// FlagMap defines a hash map of command flags.
type FlagMap struct {
	flags []Flag
}

func (fm *FlagMap) findByKey(flagKey string) FlagValue {
	for _, v := range fm.flags {
		switch fl := v.(type) {
		case FlagBool:
			if flagKey == fl.Name {
				return fl.zflag
			}
			break
		case FlagInt:
			if flagKey == fl.Name {
				return fl.zflag
			}
			break
		case FlagString:
			if flagKey == fl.Name {
				return fl.zflag
			}
			break
		case FlagStringSlice:
			if flagKey == fl.Name {
				return fl.zflag
			}
			break
		}
	}
	return FlagValue("")
}

// Bool gets current flag value as `bool`.
func (fm *FlagMap) Bool(flagName string) (bool, error) {
	return fm.findByKey(flagName).Bool()
}

// Int gets current flag value as `int`.
func (fm *FlagMap) Int(flagName string) (int, error) {
	return fm.findByKey(flagName).Int()
}

// String gets current flag value as `string`.
func (fm *FlagMap) String(flagName string) string {
	return fm.findByKey(flagName).String()
}

// StringSlice gets current flag value as a string slice.
func (fm *FlagMap) StringSlice(flagName string) []string {
	return fm.findByKey(flagName).StringSlice()
}

// StringSlice converts current flag value to a string slice.
func (v FlagValue) StringSlice() []string {
	var strs []string
	for _, s := range strings.Split(string(v), ",") {
		strs = append(strs, strings.TrimSpace(s))
	}
	return strs
}

// FlagBool defines a flag with `bool` type.
type FlagBool struct {
	Name          string
	Summary       string
	Value         bool
	Aliases       []string
	EnvVar        string
	zflag         FlagValue
	zflagAssigned bool
}

// setDefaultValue default flag values.
func (fb *FlagBool) setDefaultValue() {
	val := FlagValue(strconv.FormatBool(fb.Value))
	ev, ok := syscall.Getenv(fb.EnvVar)
	if ok {
		if b, err := FlagValue(ev).Bool(); err == nil {
			val = FlagValue(strconv.FormatBool(b))
		}
	}
	fb.zflag = val
}

// FlagInt defines a flag with `Int` type.
type FlagInt struct {
	Name          string
	Summary       string
	Value         int
	Aliases       []string
	EnvVar        string
	zflag         FlagValue
	zflagAssigned bool
}

// setDefaultValue default flag values.
func (fi *FlagInt) setDefaultValue() {
	val := FlagValue(strconv.Itoa(fi.Value))
	ev, ok := syscall.Getenv(fi.EnvVar)
	if ok {
		s := FlagValue(ev)
		if _, err := s.Int(); err == nil {
			val = s
		}
	}
	fi.zflag = val
}

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

// FlagStringSlice defines a flag with string slice type.
type FlagStringSlice struct {
	Name          string
	Summary       string
	Value         []string
	Aliases       []string
	EnvVar        string
	zflag         FlagValue
	zflagAssigned bool
}

// setDefaultValue default flag values.
func (fs *FlagStringSlice) setDefaultValue() {
	val := FlagValue(strings.Join(fs.Value, ","))
	ev, ok := syscall.Getenv(fs.EnvVar)
	if ok {
		val = FlagValue(ev)
	}
	fs.zflag = val
}

// New creates a new application instance.
func New() *App {
	return &App{}
}

// Run executes the application.
func (app *App) Run() error {
	// Commands and flags validation
	// 1. Check application flags
	aflags, err := checkAndInitFlags(app.Flags)
	if err != nil {
		return err
	}
	app.Flags = aflags

	// 2. Check commands and their flags
	var cmds []Cmd
	for _, c := range app.Commands {
		name := strings.ToLower(strings.TrimSpace(c.Name))
		if name == "" {
			return fmt.Errorf("command name has empty value")
		}
		cflags, err := checkAndInitFlags(c.Flags)
		if err != nil {
			return err
		}
		c.Flags = cflags
		cmds = append(cmds, c)
	}
	app.Commands = cmds

	// 3. Process commands and flags
	var lastCmd Cmd
	var lastFlag Flag
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
			var flags []Flag
			if hasCommand {
				flags = lastCmd.Flags
			} else {
				flags = app.Flags
			}
			// Find argument key flag on flag list
			flag := findFlagByKey(flagKey, flags)
			if flag == nil {
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
		switch fl := lastFlag.(type) {
		case FlagBool:
			if fl.Name != "" {
				if fl.zflagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}
				s := FlagValue(arg)
				if _, err := s.Bool(); err == nil {
					fl.zflag = s
					fl.zflagAssigned = true
					lastFlag = fl
				} else {
					tailArgs = append(tailArgs, arg)
				}
				continue
			}
		case FlagInt:
			if fl.Name != "" {
				if fl.zflagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}
				s := FlagValue(arg)
				if _, err := s.Int(); err == nil {
					fl.zflag = s
					fl.zflagAssigned = true
					lastFlag = fl
				} else {
					tailArgs = append(tailArgs, arg)
				}
				continue
			}
		case FlagString:
			if fl.Name != "" {
				if fl.zflagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}
				fl.zflag = FlagValue(arg)
				fl.zflagAssigned = true
				lastFlag = fl
				continue
			}
		case FlagStringSlice:
			if fl.Name != "" {
				if fl.zflagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}
				fl.zflag = FlagValue(arg)
				fl.zflagAssigned = true
				lastFlag = fl
				continue
			}
		default:
			tailArgs = append(tailArgs, arg)
			continue
		}
	}

	// Call command handler
	if hasCommand && lastCmd.Handler != nil {
		return lastCmd.Handler(&CmdContext{
			Cmd: &lastCmd,
			Flags: &FlagMap{
				flags: lastCmd.Flags,
			},
			TailArgs: &tailArgs,
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
			TailArgs: &tailArgs,
		})
	}

	return nil
}

func checkAndInitFlags(flags []Flag) ([]Flag, error) {
	var rflags []Flag
	for _, v := range flags {
		switch f := v.(type) {
		case FlagBool:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.setDefaultValue()
			rflags = append(rflags, f)
		case FlagInt:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.setDefaultValue()
			rflags = append(rflags, f)
		case FlagString:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.setDefaultValue()
			rflags = append(rflags, f)
		case FlagStringSlice:
			name := strings.ToLower(strings.TrimSpace(f.Name))
			if name == "" {
				return nil, fmt.Errorf("global flag name has empty value")
			}
			f.setDefaultValue()
			rflags = append(rflags, f)
		default:
			return nil, fmt.Errorf("global flag has invalid data type value. Use bool, int, string, []string or nil")
		}
	}
	return rflags, nil
}

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
