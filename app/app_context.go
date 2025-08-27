package app

import (
	"github.com/joseluisq/cline/flag"
)

// AppContext defines an application context.
type AppContext struct {
	app      *App
	flags    *flag.FlagValues
	tailArgs []string
}

func NewContext(app *App, flagValues *flag.FlagValues, tailArgs []string) *AppContext {
	return &AppContext{
		app:      app,
		flags:    flagValues,
		tailArgs: tailArgs,
	}
}

// App gets a reference of the current application instance.
func (c *AppContext) App() *App {
	return c.app
}

// Flags gets the global flag values for the current application.
func (c *AppContext) Flags() *flag.FlagValues {
	return c.flags
}

// TailArgs gets the current tail input arguments.
func (c *AppContext) TailArgs() []string {
	return c.tailArgs
}
