package handler

import (
	"fmt"
	"strings"

	"github.com/joseluisq/cline/app"
	"github.com/joseluisq/cline/flag"
	"github.com/joseluisq/cline/helpers"
	"github.com/joseluisq/cline/print"
)

type Handler struct {
	ap *app.App
}

// New creates a new handler for the given application.
func New(ap *app.App) *Handler {
	return &Handler{ap: ap}
}

// Run executes the current application with given arguments. Note that the first argument is always skipped.
func (h *Handler) Run(vArgs []string) error {
	// Commands and flags validation

	// 1. Check application global flags
	vflags, err := helpers.ValidateFlagsAndInit(h.ap.Flags)
	if err != nil {
		return err
	}
	h.ap.Flags = vflags

	// 2. Check commands and their flags
	vcmds, err := helpers.ValidateCommands(h.ap.Commands)
	if err != nil {
		return err
	}
	h.ap.Commands = vcmds

	// 3. Process commands and flags
	var lastCmd app.Cmd
	var lastFlag flag.Flag
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
			var flags []flag.Flag
			if hasCmd {
				flags = lastCmd.Flags
			} else {
				flags = h.ap.Flags
			}

			// Find argument key flag on flag list
			i, fl, isAlias := helpers.FindFlagByKey(flagKey, flags)
			if fl == nil {
				return fmt.Errorf("argument `%s` is not recognised", arg)
			}
			lastFlag = fl
			lastFlagIndex = i

			// Check provided incoming flags
			switch v := lastFlag.(type) {
			case flag.FlagBool:
				v.FlagProvided = true
				v.FlagProvidedAsAlias = isAlias
				lastFlag = v
			case flag.FlagInt:
				v.FlagProvided = true
				v.FlagProvidedAsAlias = isAlias
				lastFlag = v
			case flag.FlagString:
				v.FlagProvided = true
				v.FlagProvidedAsAlias = isAlias
				lastFlag = v
			case flag.FlagStringSlice:
				v.FlagProvided = true
				v.FlagProvidedAsAlias = isAlias
				lastFlag = v
			}

			// Check for bool flags and values early
			switch fl := lastFlag.(type) {
			case flag.FlagBool:
				if fl.Name != "" {
					if fl.FlagAssigned {
						tailArgs = append(tailArgs, arg)
						continue
					}

					// If bool flag is defined is assumed as `true`
					fl.FlagValue = flag.Value("1")
					// Check if we are at the last arg's item
					if argIndex == vArgsLen-1 {
						fl.FlagAssigned = true
					}
					lastFlag = fl

					if hasCmd {
						if len(lastCmd.Flags) > 0 && lastFlagIndex > -1 {
							lastCmd.Flags[lastFlagIndex] = fl
						}
					} else {
						if len(h.ap.Flags) > 0 && lastFlagIndex > -1 {
							h.ap.Flags[lastFlagIndex] = fl
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
			for _, c := range h.ap.Commands {
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
		case flag.FlagBool:
			if fl.Name != "" {
				if fl.FlagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}

				s := flag.Value(arg)
				_, err := s.ToBool()
				if err != nil {
					tailArgs = append(tailArgs, arg)
				}

				// If bool flag is defined is assumed as `true`
				if err != nil {
					s = flag.Value("1")
				}

				fl.FlagValue = s
				fl.FlagAssigned = true
				lastFlag = fl

				if hasCmd {
					if len(lastCmd.Flags) > 0 && lastFlagIndex > -1 {
						lastCmd.Flags[lastFlagIndex] = fl
					}
				} else {
					if len(h.ap.Flags) > 0 && lastFlagIndex > -1 {
						h.ap.Flags[lastFlagIndex] = fl
					}
				}

				continue
			}
		case flag.FlagInt:
			if fl.Name != "" {
				if fl.FlagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}
				s := flag.Value(arg)
				if _, err := s.ToInt(); err == nil {
					fl.FlagValue = s
					fl.FlagAssigned = true
					lastFlag = fl

					if hasCmd {
						if len(lastCmd.Flags) > 0 && lastFlagIndex > -1 {
							lastCmd.Flags[lastFlagIndex] = fl
						}
					} else {
						if len(h.ap.Flags) > 0 && lastFlagIndex > -1 {
							h.ap.Flags[lastFlagIndex] = fl
						}
					}
					continue
				} else {
					return fmt.Errorf("--%s: invalid integer value", fl.Name)
				}
			}
		case flag.FlagString:
			if fl.Name != "" {
				if fl.FlagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}
				fl.FlagValue = flag.Value(arg)
				fl.FlagAssigned = true
				lastFlag = fl

				if hasCmd {
					if len(lastCmd.Flags) > 0 && lastFlagIndex > -1 {
						lastCmd.Flags[lastFlagIndex] = fl
					}
				} else {
					if len(h.ap.Flags) > 0 && lastFlagIndex > -1 {
						h.ap.Flags[lastFlagIndex] = fl
					}
				}
				continue
			}
		case flag.FlagStringSlice:
			if fl.Name != "" {
				if fl.FlagAssigned {
					tailArgs = append(tailArgs, arg)
					continue
				}
				fl.FlagValue = flag.Value(arg)
				fl.FlagAssigned = true
				lastFlag = fl

				if hasCmd {
					if len(lastCmd.Flags) > 0 && lastFlagIndex > -1 {
						lastCmd.Flags[lastFlagIndex] = fl
					}
				} else {
					if len(h.ap.Flags) > 0 && lastFlagIndex > -1 {
						h.ap.Flags[lastFlagIndex] = fl
					}
				}
				continue
			}
		}
	}

	// Show `help` flag details
	if hasHelp {
		if hasCmd {
			return print.PrintHelp(h.ap, &lastCmd)
		}
		return print.PrintHelp(h.ap, nil)
	}

	// Show `version` flag details
	if hasVersion {
		h.ap.PrintVersion()
		return nil
	}

	// Call command handler
	if hasCmd && lastCmd.Handler != nil {
		return lastCmd.Handler(&app.CmdContext{
			Cmd: &lastCmd,
			Flags: &flag.FlagValues{
				Flags: lastCmd.Flags,
			},
			TailArgs: tailArgs,
			AppContext: &app.AppContext{
				App: h.ap,
				Flags: &flag.FlagValues{
					Flags: h.ap.Flags,
				},
			},
		})
	}

	// Call application handler
	if h.ap.Handler != nil {
		return h.ap.Handler(&app.AppContext{
			App: h.ap,
			Flags: &flag.FlagValues{
				Flags: h.ap.Flags,
			},
			TailArgs: tailArgs,
		})
	}

	return nil
}
