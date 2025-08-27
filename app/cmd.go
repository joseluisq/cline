package app

import (
	"github.com/joseluisq/cline/flag"
)

// Cmd defines an application command.
type Cmd struct {
	// The command name in alphanumeric format without special characters or spaces.
	Name string
	// A brief command description.
	Summary string
	// The command flags.
	Flags []flag.Flag
	// The command action handler.
	Handler CmdHandler
}

// CmdContext defines command context.
type CmdContext struct {
	// It references to current application command.
	Cmd *Cmd
	// It references to flag input values of the current command.
	Flags *flag.FlagValues
	// It contains current tail input arguments for the current command.
	TailArgs []string
	// It references to current application context.
	AppContext *AppContext
}

// CmdHandler responds to a command action.
type CmdHandler func(*CmdContext) error
