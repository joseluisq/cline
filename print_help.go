package cline

import (
	"fmt"
	"strings"
)

// printHelp prints current application flags and commands (--help).
func (app *App) printHelp() error {
	fmt.Printf("NAME: %s [OPTIONS] COMMAND\n\n", app.Name)
	fmt.Printf("%s\n\n", app.Summary)

	// Print options
	fmt.Printf("OPTIONS:\n")

	var flags [][]string
	var flagLen int = 0
	for _, fl := range app.Flags {
		fname := ""
		switch f := fl.(type) {
		case FlagBool:
			fname = f.Name
			flags = append(flags, []string{f.Name, f.Summary})
		case FlagInt:
			fname = f.Name
			flags = append(flags, []string{f.Name, f.Summary})
		case FlagString:
			fname = f.Name
			flags = append(flags, []string{f.Name, f.Summary})
		case FlagStringSlice:
			fname = f.Name
			flags = append(flags, []string{f.Name, f.Summary})
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
	for _, f := range flags {
		fmt.Printf("  --%s%s    %s\n", f[0], strings.Repeat(" ", flagLen-len([]rune(f[0]))), f[1])
	}

	// Print special flags
	fmt.Printf("  --%s%s    %s\n", "version", strings.Repeat(" ", flagLen-len([]rune("version"))), "Prints version information")
	fmt.Printf("  --%s%s    %s\n", "help", strings.Repeat(" ", flagLen-len([]rune("help"))), "Prints help information")

	// Print commands
	fmt.Printf("\n")
	fmt.Printf("COMMANDS:\n")

	var cmds [][]string
	var cmdLen int = 0
	for _, c := range app.Commands {
		cmds = append(cmds, []string{c.Name, c.Summary})
		if len([]rune(c.Name)) > cmdLen {
			cmdLen = len([]rune(c.Name))
		}
	}
	for _, c := range cmds {
		fmt.Printf("  %s%s    %s\n", c[0], strings.Repeat(" ", cmdLen-len([]rune(c[0]))), c[1])
	}

	fmt.Printf("\n")
	fmt.Printf("Run '%s COMMAND --help' for more information on a command\n", app.Name)
	return nil
}
