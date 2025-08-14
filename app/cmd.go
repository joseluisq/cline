package app

import (
	"github.com/joseluisq/cline/flag"
)

// Cmd defines an application command.
type Cmd struct {
	Name    string
	Summary string
	Flags   []flag.Flag
	Handler CmdHandler
}

// CmdContext defines command context.
type CmdContext struct {
	// Cmd references to current application command.
	Cmd *Cmd
	// Flags references to flag input values of current command.
	Flags *flag.FlagValues
	// TailArgs contains current tail input arguments.
	TailArgs []string
	// // AppContext references to current application context.
	AppContext *AppContext
}

// CmdHandler responds to a command action.
type CmdHandler func(*CmdContext) error
