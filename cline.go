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
	// Cmd references to current application command.
	Cmd *Cmd
	// Flags references to flag input values of current command.
	Flags *FlagValues
	// TailArgs contains current tail input arguments.
	TailArgs []string
	// AppContext references to current application context.
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
	// App references to current application instance.
	App *App
	// Flags references to flag input values of current application (global flags).
	Flags *FlagValues
	// TailArgs contains current tail input arguments.
	TailArgs []string
}

// AppHandler responds to an application action.
type AppHandler func(*AppContext) error

// New creates a new application instance.
func New() *App {
	return &App{}
}

// Run executes the current application with given arguments. Note that the first argument is always skipped.
func (app *App) Run(vArgs []string) error {
	// Commands and flags validation

	// 1. Check application global flags
	vflags, err := validateFlagsAndInit(app.Flags)
	if err != nil {
		return err
	}
	app.Flags = vflags

	// 2. Check commands and their flags
	vcmds, err := validateCommands(app.Commands)
	if err != nil {
		return err
	}
	app.Commands = vcmds

	// 3. Process commands and flags
	var lastCmd Cmd
	var lastFlag Flag
	var lastFlagIndex int = -1
	var tailArgs []string
	var hasCmd = false
	var hasHelp = false
	var hasVersion = false
	var vArgsLen = len(vArgs)

	for argIndex := 1; argIndex < vArgsLen; argIndex++ {
		arg := strings.TrimSpace(vArgs[argIndex])

		// Check for no supported arguments (remaining)
		if len(tailArgs) > 0 {
			tailArgs = append(tailArgs, arg)
			continue
		}

		// 3.1. Flags (options)
		if strings.HasPrefix(arg, "-") {
			flagKey := strings.TrimPrefix(strings.TrimPrefix(arg, "-"), "-")
			// Skip unsupported fags
			if flagKey == "" || strings.HasPrefix(flagKey, "-") {
				tailArgs = append(tailArgs, arg)
				continue
			}

			// Process special flags (help and version)
			switch flagKey {
			case "help", "h":
				hasHelp = true
			case "version", "v":
				if !hasCmd {
					hasVersion = true
				}
			}
			if hasHelp || hasVersion {
				break
			}

			// Assign flag default values
			var flags []Flag
			if hasCmd {
				flags = lastCmd.Flags
			} else {
				flags = app.Flags
			}

			// Find argument key flag on flag list
			i, flag, isAlias := findFlagByKey(flagKey, flags)
			if flag == nil {
				return fmt.Errorf("argument `%s` is not recognised", arg)
			}
			lastFlag = flag
			lastFlagIndex = i

			// Check provided incoming flags
			switch v := lastFlag.(type) {
			case FlagBool:
				v.flagProvided = true
				v.flagProvidedAsAlias = isAlias
				lastFlag = v
			case FlagInt:
				v.flagProvided = true
				v.flagProvidedAsAlias = isAlias
				lastFlag = v
			case FlagString:
				v.flagProvided = true
				v.flagProvidedAsAlias = isAlias
				lastFlag = v
			case FlagStringSlice:
				v.flagProvided = true
				v.flagProvidedAsAlias = isAlias
				lastFlag = v
			}

			// Check for bool flags and values early
			switch fl := lastFlag.(type) {
			case FlagBool:
				if fl.Name != "" {
					if fl.flagAssigned {
						tailArgs = append(tailArgs, arg)
						continue
					}

					// If bool flag is defined is assumed as `true`
					fl.flagValue = AnyValue("1")
					// Check if we are at the last arg's item
					if argIndex == vArgsLen-1 {
						fl.flagAssigned = true
					}
					lastFlag = fl

					if hasCmd {
						if len(lastCmd.Flags) > 0 && lastFlagIndex > -1 {
							lastCmd.Flags[lastFlagIndex] = fl
						}
					} else {
						if len(app.Flags) > 0 && lastFlagIndex > -1 {
							app.Flags[lastFlagIndex] = fl
						}
					}

					continue
				}
			}

			continue
		}

		// 3.2. Commands
		// 3.2.1 Check for a valid command (first time)
		if !hasCmd {
			for _, c := range app.Commands {
				if c.Name == arg {
					hasCmd = true
					lastCmd = c
					break
				}
			}
			if hasCmd {
				continue
			}
		}

		// 4. If there is no command found assume it as a tail arg
		if lastFlag == nil {
			tailArgs = append(tailArgs, arg)
			continue
		}

		// 5. Process app or command flag values
		switch fl := lastFlag.(type) {
		case FlagBool:
			if fl.Name != "" {
				if fl.flagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}

				s := AnyValue(arg)
				_, err := s.ToBool()
				if err != nil {
					tailArgs = append(tailArgs, arg)
				}

				// If bool flag is defined is assumed as `true`
				if err != nil {
					s = AnyValue("1")
				}

				fl.flagValue = s
				fl.flagAssigned = true
				lastFlag = fl

				if hasCmd {
					if len(lastCmd.Flags) > 0 && lastFlagIndex > -1 {
						lastCmd.Flags[lastFlagIndex] = fl
					}
				} else {
					if len(app.Flags) > 0 && lastFlagIndex > -1 {
						app.Flags[lastFlagIndex] = fl
					}
				}

				continue
			}
		case FlagInt:
			if fl.Name != "" {
				if fl.flagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}
				s := AnyValue(arg)
				if _, err := s.ToInt(); err == nil {
					fl.flagValue = s
					fl.flagAssigned = true
					lastFlag = fl

					if hasCmd {
						if len(lastCmd.Flags) > 0 && lastFlagIndex > -1 {
							lastCmd.Flags[lastFlagIndex] = fl
						}
					} else {
						if len(app.Flags) > 0 && lastFlagIndex > -1 {
							app.Flags[lastFlagIndex] = fl
						}
					}
					continue
				} else {
					return fmt.Errorf("--%s: invalid integer value", fl.Name)
				}
			}
		case FlagString:
			if fl.Name != "" {
				if fl.flagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}
				fl.flagValue = AnyValue(arg)
				fl.flagAssigned = true
				lastFlag = fl

				if hasCmd {
					if len(lastCmd.Flags) > 0 && lastFlagIndex > -1 {
						lastCmd.Flags[lastFlagIndex] = fl
					}
				} else {
					if len(app.Flags) > 0 && lastFlagIndex > -1 {
						app.Flags[lastFlagIndex] = fl
					}
				}
				continue
			}
		case FlagStringSlice:
			if fl.Name != "" {
				if fl.flagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}
				fl.flagValue = AnyValue(arg)
				fl.flagAssigned = true
				lastFlag = fl

				if hasCmd {
					if len(lastCmd.Flags) > 0 && lastFlagIndex > -1 {
						lastCmd.Flags[lastFlagIndex] = fl
					}
				} else {
					if len(app.Flags) > 0 && lastFlagIndex > -1 {
						app.Flags[lastFlagIndex] = fl
					}
				}
				continue
			}
		}
	}

	// Show `help` flag details
	if hasHelp {
		if hasCmd {
			return printHelp(app, &lastCmd)
		}
		return printHelp(app, nil)
	}

	// Show `version` flag details
	if hasVersion {
		app.printVersion()
		return nil
	}

	// Call command handler
	if hasCmd && lastCmd.Handler != nil {
		return lastCmd.Handler(&CmdContext{
			Cmd: &lastCmd,
			Flags: &FlagValues{
				flags: lastCmd.Flags,
			},
			TailArgs: tailArgs,
			AppContext: &AppContext{
				App: app,
				Flags: &FlagValues{
					flags: app.Flags,
				},
			},
		})
	}

	// Call application handler
	if app.Handler != nil {
		return app.Handler(&AppContext{
			App: app,
			Flags: &FlagValues{
				flags: app.Flags,
			},
			TailArgs: tailArgs,
		})
	}

	return nil
}
