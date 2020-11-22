# Cline [![Build Status](https://travis-ci.com/joseluisq/cline.svg?branch=master)](https://travis-ci.com/joseluisq/cline) [![codecov](https://codecov.io/gh/joseluisq/cline/branch/master/graph/badge.svg)](https://codecov.io/gh/joseluisq/cline) [![Go Report Card](https://goreportcard.com/badge/github.com/joseluisq/cline)](https://goreportcard.com/report/github.com/joseluisq/cline) [![GoDoc](https://godoc.org/github.com/joseluisq/cline?status.svg)](https://pkg.go.dev/github.com/joseluisq/cline)

> A fast and lightweight CLI package for Go without external dependencies.

WIP project under **active** development.

## Features

- No external dependencies.
- Compact API.
- `bool`, `int`, `string` and `[]string` data types only.
- Global and local flags support.
- Flag aliases and default values support.
- Convenient context for function handlers (global and commands) with build-in types conversion API.
- Optional environment variable names for flags.
- Automatic `--help` (`-h`) and `--version` (`-v`) generation flags.
- [POSIX-compliant](https://github.com/joseluisq/cline/issues/3) flags for short and long versions (WIP)

## Usage

API definition example:

```go
package main

import (
	"fmt"
	"os"

	cli "github.com/joseluisq/cline"
)

var (
	version   string = "devel"
	buildTime string
)

func main() {
	app := cli.New()
	app.Name = "enve"
	app.Summary = "Run a program in a modified environment using .env files"
	app.Version = version
	app.BuildTime = buildTime
	app.Flags = []cli.Flag{
		cli.FlagString{
			Name:    "file",
			Summary: "Load environment variables from a file path",
			Value:   ".env",
			Aliases: []string{"f"},
		},
		cli.FlagBool{
			Name:    "verbose",
			Summary: "Enable more verbose info",
			Value:   false,
			Aliases: []string{"V"},
			EnvVar:  "ENV_VERBOSE",
		},
	}
	app.Commands = []cli.Cmd{
		{
			Name:    "info",
			Summary: "Show command information",
			Flags: []cli.Flag{
				cli.FlagInt{
					Name:    "trace",
					Summary: "Enable tracing mode",
					Value:   10,
					Aliases: []string{"t"},
				},
				cli.FlagBool{
					Name:    "detailed",
					Summary: "Enable info details",
					Value:   true,
					Aliases: []string{"d"},
				},
			},
			Handler: func(ctx *cli.CmdContext) error {
				fmt.Printf("Cmd `%s` executed!\n", ctx.Cmd.Name)
				fmt.Printf("App Flag `file` opted: `%s`\n", ctx.AppContext.Flags.StringSlice("file"))

				i, err := ctx.Flags.Int("trace")
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Printf("Cmd Flag `trace` opted: `%d` (%T)\n", i, i)

				d := ctx.Flags.String("detailed")
				fmt.Printf("Cmd Flag `detailed` opted: `%s` (%T)\n", d, d)
				fmt.Printf("Cmd Tail arguments: %#v\n", ctx.TailArgs)
				return nil
			},
		},
	}
	app.Handler = func(ctx *cli.AppContext) error {
		fmt.Printf("App `%s` executed!\n", ctx.App.Name)
		fmt.Printf("App Tail arguments: %#v\n", ctx.TailArgs)
		fmt.Printf("App Flag `file` opted: `%s`\n", ctx.Flags.StringSlice("file"))

		b, _ := ctx.Flags.Bool("verbose")

		fmt.Printf("App Flag `verbose` opted: `%v`\n", b)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

Output example:

```sh
$ go run examples/main.go -h
# NAME: enve [OPTIONS] COMMAND
#
# Run a program in a modified environment using .env files
#
# OPTIONS:
#   --file       Load environment variables from a file path
#   --verbose    Enable more verbose info
#   --version    Prints version information
#   --help       Prints help information
#
# COMMANDS:
#   info    Show command information
#
# Run 'enve COMMAND --help' for more information on a command

$ go run examples/main.go -v
# Version:       0.0.0
# Go version:    go1.15.4
# Built:         2020-11-10T21:11:36
# OS/Arch:       linux/amd64
```

More details on [examples/main.go](./examples/main.go)

## Contributions

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in current work by you, as defined in the Apache-2.0 license, shall be dual licensed as described below, without any additional terms or conditions.

Feel free to send some [Pull request](https://github.com/joseluisq/cline/pulls) or [issue](https://github.com/joseluisq/cline/issues).

## License

This work is primarily distributed under the terms of both the [MIT license](LICENSE-MIT) and the [Apache License (Version 2.0)](LICENSE-APACHE).

© 2020-present [Jose Quintana](https://git.io/joseluisq)
