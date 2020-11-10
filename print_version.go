package cline

import (
	"fmt"
	"runtime"
)

// printVersion prints current application version (--version).
func (app *App) printVersion() {
	fmt.Printf("Version:       %s\n", app.Version)
	fmt.Printf("Go version:    %s\n", runtime.Version())
	fmt.Printf("Built:         %s\n", app.BuildTime)
	fmt.Printf("OS/Arch:       %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
