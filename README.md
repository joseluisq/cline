# Cline [![Build Status](https://travis-ci.com/joseluisq/cline.svg?branch=master)](https://travis-ci.com/joseluisq/cline) [![codecov](https://codecov.io/gh/joseluisq/cline/branch/master/graph/badge.svg)](https://codecov.io/gh/joseluisq/cline) [![Go Report Card](https://goreportcard.com/badge/github.com/joseluisq/cline)](https://goreportcard.com/report/github.com/joseluisq/cline) [![GoDoc](https://godoc.org/github.com/joseluisq/cline?status.svg)](https://pkg.go.dev/github.com/joseluisq/cline)

> A fast and lightweight CLI package for Go with no dependencies.

WIP project under **active** development.

## Features

- No external dependencies.
- Compact API.
- `bool`, `int`, `string` and `[]string` data types only.
- Global and local flags support.
- Flag aliases and default values support.
- Convenient context for function handlers (global and commands) with build-in types conversion API.
- Optional environment variable names for flags.
- [POSIX-compliant](https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html) flags for short and long versions (WIP)
- Automatic `--help` (`-h`) and `--version` (`-v`) generation flags. (WIP)

## Usage

```go
package main

import (
	"fmt"
	"os"

	cli "github.com/joseluisq/cline"
)

func main() {
	app := cli.New()
	app.Name = "enve"
	app.Summary = "run a program in a modified environment using .env files"
	app.Flags = []cli.Flag{
		cli.FlagString{
			Name:    "file",
			Summary: "load environment variables from a file path",
			Value:   ".env",
			Aliases: []string{"f"},
		},
		cli.FlagBool{
			Name:    "verbose",
			Summary: "load environment variables from a file path",
			Value:   false,
			Aliases: []string{"v"},
			EnvVar:  "ENV_FILE",
		},
	}
	app.Commands = []cli.Cmd{
		{
			Name:    "info",
			Summary: "show command information",
			Flags: []cli.Flag{
				cli.FlagInt{
					Name:    "version",
					Summary: "enable more verbose command information",
					Value:   10,
					Aliases: []string{"z"},
				},
				cli.FlagBool{
					Name:    "detailed",
					Summary: "enable info details",
					Value:   true,
					Aliases: []string{"d"},
				},
			},
			Handler: func(ctx *cli.CmdContext) error {
				fmt.Printf("Cmd `%s` executed!\n", ctx.Cmd.Name)
				i, err := ctx.Flags.Int("version")
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Printf("Cmd Flag `version` opted: `%d` (%T)\n", i, i)
				fmt.Printf("Cmd Tail arguments: %#v\n", ctx.TailArgs)
				return nil
			},
		},
	}
	app.Handler = func(ctx *cli.AppContext) error {
		fmt.Printf("App `%s` executed!\n", ctx.App.Name)
		fmt.Printf("App Tail arguments: %#v\n", ctx.TailArgs)
		fmt.Printf("App Flag `file` opted: `%s`\n", ctx.Flags.StringSlice("file"))
		fmt.Printf("App Flag `verbose` opted: `%s`\n", ctx.Flags.StringSlice("verbose"))
		return nil
	}
	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

## Contributions

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in current work by you, as defined in the Apache-2.0 license, shall be dual licensed as described below, without any additional terms or conditions.

Feel free to send some [Pull request](https://github.com/joseluisq/cline/pulls) or [issue](https://github.com/joseluisq/cline/issues).

## License

This work is primarily distributed under the terms of both the [MIT license](LICENSE-MIT) and the [Apache License (Version 2.0)](LICENSE-APACHE).

© 2020-present [Jose Quintana](https://git.io/joseluisq)
