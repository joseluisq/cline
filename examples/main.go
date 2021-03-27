package main

import (
	"fmt"
	"os"

	cli "github.com/joseluisq/cline"
)

// App version and build time values passed at compile time
// See `Makefile` > build
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
		cli.FlagInt{
			Name:    "int",
			Summary: "Int value",
			Value:   5,
			Aliases: []string{"b"},
		},
		cli.FlagBool{
			Name:    "verbose",
			Summary: "Enable more verbose info",
			Value:   true,
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
				cli.FlagString{
					Name:    "FF",
					Value:   ".env",
					Aliases: []string{"f"},
				},
				cli.FlagStringSlice{
					Name:    "II",
					Value:   []string{"q", "r", "s"},
					Aliases: []string{"i"},
				},
			},
			Handler: func(ctx *cli.CmdContext) error {
				fmt.Printf("Cmd `%s` executed!\n", ctx.Cmd.Name)

				if file, err := ctx.AppContext.Flags.String("file"); err != nil {
					fmt.Printf("App Flag `file`: `%s`\n", file.Value())
				}

				if verbose, err := ctx.AppContext.Flags.Bool("verbose"); err != nil {
					b, _ := verbose.Value()
					fmt.Printf("App Flag `verbose`: `%v`\n", b)
				}

				if trace, err := ctx.AppContext.Flags.Int("trace"); err != nil {
					i, err := trace.Value()
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					fmt.Printf("Cmd Flag `trace`: `%d` (%T)\n", i, i)
				}

				if detailed, err := ctx.AppContext.Flags.Bool("detailed"); err != nil {
					d, _ := detailed.Value()
					fmt.Printf("Cmd Flag `detailed`: `%v` (%T)\n", d, d)
				}

				fmt.Printf("Cmd Tail arguments: %#v\n", ctx.TailArgs)
				return nil
			},
		},
	}
	app.Handler = func(ctx *cli.AppContext) error {
		fmt.Printf("App `%s` executed!\n", ctx.App.Name)

		if file, err := ctx.Flags.Bool("file"); err != nil {
			if b, err := file.Value(); err != nil {
				fmt.Printf("App Flag `file`: `%t`\n", b)
			}
		}

		fmt.Printf("App Flag `int`: `%v`\n", ctx.Flags.Any("int"))

		if verbose, err := ctx.Flags.Bool("verbose"); err != nil {
			if b, err := verbose.Value(); err != nil {
				fmt.Printf("App Flag `verbose`: `%v`\n", b)
			}
		}

		fmt.Printf("App Tail arguments: %#v\n", ctx.TailArgs)
		fmt.Printf("App Provided flags: %v\n", ctx.Flags.GetProvided())
		fmt.Printf("App Provided flags (long): %v\n", ctx.Flags.GetProvidedLong())
		fmt.Printf("App Provided flags (short): %v\n", ctx.Flags.GetProvidedShort())
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
