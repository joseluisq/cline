// Package cline is a fast and lightweight CLI package for Go.
//
package cline

import (
	"fmt"
	"strings"
)

// Flag defines a flag generic type.
type Flag interface{}

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
	Name      string
	Summary   string
	Version   string
	BuildTime string
	Flags     []Flag
	Commands  []Cmd
	Handler   AppHandler
}

// AppContext defines an application context.
type AppContext struct {
	App      *App
	Flags    *FlagMap
	TailArgs *[]string
}

// AppHandler responds to an application action.
type AppHandler func(*AppContext) error

// New creates a new application instance.
func New() *App {
	return &App{}
}

// Run executes the application.
func (app *App) Run(vArgs []string) error {
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
	var hasHelp = false
	var hasVersion = false

	for i := 1; i < len(vArgs); i++ {
		arg := strings.ToLower(strings.TrimSpace(vArgs[i]))

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

			// Process special flags (help and version)
			switch flagKey {
			case "help", "h":
				hasHelp = true
			case "version", "v":
				hasVersion = true
			}
			if hasHelp || hasVersion {
				break
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

	// Show `help` flag details
	if hasHelp {
		return app.printHelp()
	}

	// Show `version` flag details
	if hasVersion {
		app.printVersion()
		return nil
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
