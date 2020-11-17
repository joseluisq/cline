package cline

import (
	"fmt"
	"runtime"
	"strings"
)

// printHelp prints current application flags and commands info (--help).
func printHelp(app *App, cmd *Cmd) error {
	if app == nil {
		return fmt.Errorf("application instance not found")
	}

	var summary string
	var flags []Flag

	if cmd == nil {
		summary = app.Summary
		flags = app.Flags
	} else {
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

	var vflags [][]string
	var flagLen int = 0
	for _, fl := range flags {
		fname := ""
		switch f := fl.(type) {
		case FlagBool:
			fname = f.Name
			vflags = append(vflags, []string{f.Name, f.Summary})
		case FlagInt:
			fname = f.Name
			vflags = append(vflags, []string{f.Name, f.Summary})
		case FlagString:
			fname = f.Name
			vflags = append(vflags, []string{f.Name, f.Summary})
		case FlagStringSlice:
			fname = f.Name
			vflags = append(vflags, []string{f.Name, f.Summary})
		}
		if len([]rune(fname)) > flagLen {
			flagLen = len([]rune(fname))
		}
	}
	if len([]rune("version")) > flagLen {
		flagLen = len([]rune("version"))
	}
	if len([]rune("help")) > flagLen {
		flagLen = len([]rune("help"))
	}
	for _, f := range vflags {
		fmt.Printf(
			"  --%s%s    %s\n",
			f[0],
			strings.Repeat(" ", flagLen-len([]rune(f[0]))),
			f[1],
		)
	}

	// App with commands
	if cmd == nil {
		// Print special flags
		fmt.Printf(
			"  --%s%s    %s\n",
			"version",
			strings.Repeat(" ", flagLen-len([]rune("version"))),
			"Prints version information",
		)
		fmt.Printf(
			"  --%s%s    %s\n",
			"help",
			strings.Repeat(" ", flagLen-len([]rune("help"))),
			"Prints help information",
		)

		// Print commands
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
			fmt.Printf("  %s%s    %s\n", c[0], strings.Repeat(
				" ", cmdLen-len([]rune(c[0]))), c[1],
			)
		}

		fmt.Printf("\n")
		fmt.Printf("Run '%s COMMAND --help' for more information on a command\n", app.Name)
	} else {
		fmt.Printf("\n")
		fmt.Printf("Run '%s %s COMMAND --help' for more information on a command\n", app.Name, cmd.Name)
	}
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
	fmt.Printf("OS/Arch:       %s/%s\n", runtime.GOOS, runtime.GOARCH)
	return nil
}
