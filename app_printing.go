package cline

import (
	"fmt"
	"runtime"
	"strings"
)

type flagStruct struct {
	name     string
	aliases  []string
	summary  string
	defaults string
	envVar   string
}

// printHelp prints current application flags and commands info (--help).
func printHelp(app *App, cmd *Cmd) error {
	if app == nil {
		return fmt.Errorf("application instance not found")
	}

	summary := app.Summary
	flags := app.Flags
	if cmd != nil {
		summary = cmd.Summary
		flags = cmd.Flags
	}

	// TODO: subcommands support
	if cmd == nil {
		fmt.Printf("NAME: %s [OPTIONS] COMMAND\n\n", app.Name)
	} else {
		fmt.Printf("NAME: %s %s [OPTIONS] COMMAND\n\n", app.Name, cmd.Name)
	}
	fmt.Printf("%s\n\n", summary)

	// Print options
	fmt.Printf("OPTIONS:\n")

	var vflags []flagStruct
	var fLen int = 0
	var aliasMaxLen = 0

	// Append help and version flags
	flags = append(flags, FlagString{
		Name: "help", Aliases: []string{"h"}, Summary: "Prints help information",
	})
	if cmd == nil {
		flags = append(flags, FlagString{
			Name: "version", Aliases: []string{"v"}, Summary: "Prints version information",
		})
	}

	// Calculate app or command flags positions
	for _, fl := range flags {
		var vFlag flagStruct

		fname := ""
		switch f := fl.(type) {
		case FlagBool:
			fname = f.Name
			vFlag = flagStruct{name: f.Name, aliases: f.Aliases, summary: f.Summary, defaults: f.flagValue.ToString(), envVar: f.EnvVar}
		case FlagInt:
			fname = f.Name
			vFlag = flagStruct{name: f.Name, aliases: f.Aliases, summary: f.Summary, defaults: f.flagValue.ToString(), envVar: f.EnvVar}
		case FlagString:
			fname = f.Name
			vFlag = flagStruct{name: f.Name, aliases: f.Aliases, summary: f.Summary, defaults: f.flagValue.ToString(), envVar: f.EnvVar}
		case FlagStringSlice:
			fname = f.Name
			vFlag = flagStruct{name: f.Name, aliases: f.Aliases, summary: f.Summary, defaults: f.flagValue.ToString(), envVar: f.EnvVar}
		}
		if len([]rune(fname)) > fLen {
			fLen = len([]rune(fname))
		}

		var aliases []string
		for _, a := range vFlag.aliases {
			s := "-" + a
			aliases = append(aliases, s)
		}
		aliasesLen := len([]rune(strings.Join(aliases, ",")))
		if aliasesLen > aliasMaxLen {
			aliasMaxLen = aliasesLen
		}

		vFlag.aliases = aliases
		vflags = append(vflags, vFlag)
	}

	// Print app or command flags
	for _, v := range vflags {
		shorts := strings.Join(v.aliases, ",")

		repeatLeft := aliasMaxLen - len([]rune(shorts))
		marginLeftRepeat := strings.Repeat(" ", repeatLeft)

		repeatRight := fLen - len([]rune(v.name))
		marginRightRepeat := strings.Repeat(" ", repeatRight)

		defaultVal := strings.TrimSpace(v.defaults)
		if defaultVal != "" {
			defaultVal = " [default: " + defaultVal + "]"
		}
		envVar := strings.TrimSpace(v.envVar)
		if envVar != "" {
			envVar = " [env: " + envVar + "]"
		}

		line := fmt.Sprintf(
			"  %s%s --%s%s    %s%s%s\n",
			marginLeftRepeat,
			shorts,
			v.name,
			marginRightRepeat,
			v.summary,
			defaultVal,
			envVar,
		)

		fmt.Println(line)
	}

	// Print app commands
	if cmd == nil {
		if len(app.Commands) > 0 {
			fmt.Printf("\n")
			fmt.Printf("COMMANDS:\n")

			var vcmds [][]string
			var cmdLen int = 0
			for _, c := range app.Commands {
				vcmds = append(vcmds, []string{c.Name, c.Summary})
				if len([]rune(c.Name)) > cmdLen {
					cmdLen = len([]rune(c.Name))
				}
			}
			for _, c := range vcmds {
				fmt.Printf(
					"  %s%s%s    %s\n",
					"",
					c[0],
					strings.Repeat(
						" ", cmdLen-len([]rune(c[0])),
					),
					c[1],
				)
			}

			fmt.Printf("\n")
			fmt.Printf("Run '%s COMMAND --help' for more information on a command\n", app.Name)
		}
	} else {
		fmt.Printf("\n")
		fmt.Printf("Run '%s %s COMMAND --help' for more information on a command\n", app.Name, cmd.Name)
	}

	fmt.Printf("\n")
	return nil
}

// printVersion prints current application version (--version).
func (app *App) printVersion() error {
	if app == nil {
		return fmt.Errorf("application instance not found")
	}
	fmt.Printf("Version:       %s\n", app.Version)
	fmt.Printf("Go version:    %s\n", runtime.Version())
	fmt.Printf("Built:         %s\n", app.BuildTime)
	fmt.Printf("Commit:        %s\n", app.BuildCommit)
	fmt.Printf("OS/Arch:       %s/%s\n", runtime.GOOS, runtime.GOARCH)
	return nil
}
