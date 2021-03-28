# CLIne ![CI](https://github.com/joseluisq/cline/workflows/CI/badge.svg) [![codecov](https://codecov.io/gh/joseluisq/cline/branch/master/graph/badge.svg)](https://codecov.io/gh/joseluisq/cline) [![Go Report Card](https://goreportcard.com/badge/github.com/joseluisq/cline)](https://goreportcard.com/report/github.com/joseluisq/cline) [![PkgGoDev](https://pkg.go.dev/badge/github.com/joseluisq/cline)](https://pkg.go.dev/github.com/joseluisq/cline)

> A fast and lightweight CLI package for Go without external dependencies.

WIP project under **active** development.

## Features

- No external dependencies more than few [Go's stdlib](https://golang.org/pkg/#stdlib) ones.
- Compact but concise API.
- Global and command flags support.
- `bool`, `int`, `string` and `[]string` flag's data types.
- Flag aliases and default values support.
- Convenient contexts for function handlers (global and command flags)
- Context built-in types conversion API for `bool`, `int`, `string` and `[]string` flag values.
- Convenient API to detect provided (passed) flags only with thier props.
- Optional environment variable names for flags.
- Automatic `--help` (`-h`) flag for global flags and commands.
- Automatic `--version` (`-v`) flag with relevant info like app version, Go version, build datetime and OS/Arch.
- POSIX-compliant flags (partial but [WIP](https://github.com/joseluisq/cline/issues/3))

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
				fmt.Printf("App Flag `file` opted: `%s`\n", ctx.AppContext.Flags.Any("file"))

				trace, err := ctx.Flags.Int("trace")
				if err != nil {
					return err
				}
				i, err := trace.Value()
				if err != nil {
					return err
				}

				fmt.Printf("Cmd Flag `trace` opted: `%d` (%T)\n", i, i)

				detailed, err := ctx.Flags.Bool("detailed")
				if err != nil {
					return err
				}
				d, err := detailed.Value()
				if err != nil {
					return err
				}

				fmt.Printf("Cmd Flag `detailed` opted: `%v` (%T)\n", d, d)

				fmt.Printf("Cmd Tail arguments: %#v\n", ctx.TailArgs)
				return nil
			},
		},
	}
	app.Handler = func(ctx *cli.AppContext) error {
		fmt.Printf("App `%s` executed!\n", ctx.App.Name)
		fmt.Printf("App Tail arguments: %#v\n", ctx.TailArgs)

		if f, err := ctx.Flags.StringSlice("file"); err == nil {
			fmt.Printf("App Flag `file` opted: `%v`\n", f.Value())
		}

		if v, err := ctx.Flags.Bool("verbose"); err == nil {
			b, _ := v.Value()
			fmt.Printf("App Flag `verbose` opted: `%v`\n", b)
		}

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
# enve 0.0.0
# Run a program in a modified environment using .env files

# USAGE:
#    enve [OPTIONS] COMMAND
#
# OPTIONS:
#       -f --file      Load environment variables from a file path.
#                      Some new line description
#                      Another new line description. [default: .env]
#    -b,-z --int       Int value [default: 5]
#       -V --verbose   Enable more verbose info [default: true] [env: ENV_VERBOSE]
#       -h --help      Prints help information
#       -v --version   Prints version information
#
# COMMANDS:
#    info   Show command information
#
# Run 'enve COMMAND --help' for more information on a command

$ go run examples/main.go -v
# Version:       0.0.0
# Go version:    go1.16.2
# Built:         2021-03-28T20:03:04
# Commit:        1885ad1
# OS/Arch:       linux/amd64
```

More details on [examples/main.go](./examples/main.go)

## Contributions

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in current work by you, as defined in the Apache-2.0 license, shall be dual licensed as described below, without any additional terms or conditions.

Feel free to send some [Pull request](https://github.com/joseluisq/cline/pulls) or [issue](https://github.com/joseluisq/cline/issues).

## License

This work is primarily distributed under the terms of both the [MIT license](LICENSE-MIT) and the [Apache License (Version 2.0)](LICENSE-APACHE).

© 2020-present [Jose Quintana](https://git.io/joseluisq)
