// Package cline is a fast and lightweight CLI package for Go.
package app

import (
	"fmt"
	"runtime"

	"github.com/joseluisq/cline/flag"
)

// App defines application settings
type App struct {
	Name        string
	Summary     string
	Version     string
	BuildTime   string
	BuildCommit string
	Flags       []flag.Flag
	Commands    []Cmd
	Handler     AppHandler
}

// AppContext defines an application context.
type AppContext struct {
	// App references to current application instance.
	App *App
	// Flags references to flag input values of current application (global flags).
	Flags *flag.FlagValues
	// TailArgs contains current tail input arguments.
	TailArgs []string
}

// AppHandler responds to an application action.
type AppHandler func(*AppContext) error

// New creates a new application instance.
func New() *App {
	return &App{}
}

// printVersion prints current application version (--version).
func (ap *App) PrintVersion() error {
	if ap == nil {
		return fmt.Errorf("application instance not found")
	}
	fmt.Printf("Version:       %s\n", ap.Version)
	fmt.Printf("Go version:    %s\n", runtime.Version())
	fmt.Printf("Built:         %s\n", ap.BuildTime)
	fmt.Printf("Commit:        %s\n", ap.BuildCommit)
	fmt.Printf("OS/Arch:       %s/%s\n", runtime.GOOS, runtime.GOARCH)
	return nil
}

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
