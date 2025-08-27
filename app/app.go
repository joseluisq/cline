// Package cline is a fast and lightweight CLI package for Go.
package app

import (
	"fmt"
	"runtime"

	"github.com/joseluisq/cline/flag"
)

// App defines application settings.
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

// New creates a new application instance.
func New() *App {
	return &App{}
}

// printVersion prints current application version (--version).
func (ap *App) PrintVersion() {
	fmt.Printf("Version:       %s\n", ap.Version)
	fmt.Printf("Go version:    %s\n", runtime.Version())
	fmt.Printf("Built:         %s\n", ap.BuildTime)
	fmt.Printf("Commit:        %s\n", ap.BuildCommit)
}
