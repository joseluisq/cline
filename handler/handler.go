// Package handler provides a handler for the CLI application.
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

const (
	maxArgLen    = 4096 // 4KB
	maxArgsCount = 1024
)

// New creates a new handler for the given application.
func New(ap *app.App) *Handler {
	return &Handler{ap: ap}
}

// Run executes the current application with given arguments.
// Note that the first argument is always skipped.
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

	// Optimization: Create maps for faster flag lookups
	appFlagMap := helpers.BuildFlagMap(h.ap.Flags)
	cmdFlagMaps := make(map[string]map[string]helpers.FlagInfo, len(h.ap.Commands))
	for _, cmd := range h.ap.Commands {
		cmdFlagMaps[cmd.Name] = helpers.BuildFlagMap(cmd.Flags)
	}

	// 3. Process commands and flags
	var lastCmd *app.Cmd
	var lastFlag flag.Flag
	var lastFlagIndex = -1
	var tailArgs = make([]string, 0, 4)
	var hasCmd = false
	var hasHelp = false
	var hasVersion = false
	var vArgsLen = len(vArgs)

	if vArgsLen > maxArgsCount {
		return fmt.Errorf("number of arguments exceeds the limit of %d", maxArgsCount)
	}

	for idx := 1; idx < vArgsLen; idx++ {
		arg := strings.TrimSpace(vArgs[idx])

		if len(arg) > maxArgLen {
			return fmt.Errorf("argument exceeds maximum length of %d characters", maxArgLen)
		}

		// Check for no supported arguments (remaining)
		if arg == "--" {
			if idx+1 < vArgsLen {
				tailArgs = append(tailArgs, vArgs[idx+1:]...)
			}
			break // Stop processing further arguments as flags
		}

		if len(tailArgs) > 0 {
			tailArgs = append(tailArgs, arg)
			continue
		}

		// 3.1. Flags (options)
		if strings.HasPrefix(arg, "-") {
			var flagKey string
			isAlias := !strings.HasPrefix(arg, "--")
			if isAlias {
				flagKey = arg[1:]
			} else {
				flagKey = arg[2:]
			}

			// Skip unsupported flags
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

			var flagMap map[string]helpers.FlagInfo
			if hasCmd {
				flagMap = cmdFlagMaps[lastCmd.Name]
			} else {
				flagMap = appFlagMap
			}

			flagInfo, ok := flagMap[flagKey]
			if !ok {
				return fmt.Errorf("unknown argument: %s", arg)
			}
			lastFlag = flagInfo.Flag
			lastFlagIndex = flagInfo.Index

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

					// A boolean flag is considered true on its own
					fl.FlagValue = flag.Value("true")
					fl.FlagAssigned = true

					// Check if the next argument could be a value for this bool flag
					if idx+1 < vArgsLen {
						nextArg := vArgs[idx+1]
						if _, err := flag.Value(nextArg).ToBool(); err == nil {
							fl.FlagValue = flag.Value(nextArg)
							// Skip next argument
							idx++
						}
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
			for i, c := range h.ap.Commands {
				if c.Name == arg {
					hasCmd = true
					lastCmd = &h.ap.Commands[i]
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

				// This case is now handled when the flag is first seen.
				// If we are here, it means the argument isn't a value for the bool flag.
				tailArgs = append(tailArgs, arg)
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
					return fmt.Errorf("invalid integer value for flag --%s", fl.Name)
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
			return print.PrintHelp(h.ap, lastCmd)
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
			Cmd:      lastCmd,
			Flags:    flag.NewFlagValues(lastCmd.Flags),
			TailArgs: tailArgs,
			AppContext: app.NewContext(
				h.ap,
				flag.NewFlagValues(h.ap.Flags),
				[]string{},
			),
		})
	}

	// Call application handler
	if h.ap.Handler != nil {
		ctx := app.NewContext(
			h.ap,
			flag.NewFlagValues(h.ap.Flags),
			tailArgs,
		)
		return h.ap.Handler(ctx)
	}

	return nil
}
