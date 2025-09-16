// Package app provides the application structure and methods.
package app

import (
	"fmt"
	"runtime"

	"github.com/joseluisq/cline/flag"
)

// App defines the application settings.
type App struct {
	// The application name.
	Name string
	// A brief application description.
	Summary string
	// The application version.
	Version string
	// The application build time.
	BuildTime string
	// The application build commit.
	BuildCommit string
	// The application flags.
	Flags []flag.Flag
	// The application commands.
	Commands []Cmd
	// The application action handler.
	Handler AppHandler
}

// New creates a new application instance.
func New() *App {
	return &App{}
}

// PrintVersion prints the current application version information (--version).
func (ap *App) PrintVersion() {
	fmt.Printf("Version:       %s\n", ap.Version)
	fmt.Printf("Go version:    %s\n", runtime.Version())
	fmt.Printf("Built:         %s\n", ap.BuildTime)
	fmt.Printf("Commit:        %s\n", ap.BuildCommit)
}
