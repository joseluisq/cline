package main

import (
	"fmt"
	"os"

	cli "github.com/joseluisq/cline"
)

// App version and build time values passed at compile time
// See `Makefile` > build
var (
	version     string = "devel"
	buildTime   string
	buildCommit string
)

func main() {
	app := cli.New()
	app.Name = "enve"
	app.Summary = "Run a program in a modified environment using .env files"
	app.Version = version
	app.BuildTime = buildTime
	app.BuildCommit = buildCommit
	app.Flags = []cli.Flag{
		cli.FlagString{
			Name:    "file",
			Summary: "Load environment variables from a file path.\nSome new line description\nAnother new line description.",
			Value:   ".env",
			Aliases: []string{"f"},
		},
		cli.FlagInt{
			Name:    "int",
			Summary: "Int value",
			Value:   5,
			Aliases: []string{"b", "z"},
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

				file, err := ctx.AppContext.Flags.String("file")
				if err != nil {
					return err
				}
				fmt.Printf("App Flag `file`: `%s`\n", file.Value())

				verbose, err := ctx.AppContext.Flags.Bool("verbose")
				if err != nil {
					return err
				}
				b, _ := verbose.Value()
				fmt.Printf("App Flag `verbose`: `%v`\n", b)

				trace, err := ctx.AppContext.Flags.Int("trace")
				if err != nil {
					return err
				}
				i, err := trace.Value()
				if err != nil {
					return err
				}
				fmt.Printf("Cmd Flag `trace`: `%d` (%T)\n", i, i)

				detailed, err := ctx.AppContext.Flags.Bool("detailed")
				if err != nil {
					return err
				}
				d, _ := detailed.Value()
				fmt.Printf("Cmd Flag `detailed`: `%v` (%T)\n", d, d)

				fmt.Printf("Cmd Tail arguments: %#v\n", ctx.TailArgs)
				return nil
			},
		},
	}
	app.Handler = func(ctx *cli.AppContext) error {
		fmt.Printf("App `%s` executed!\n", ctx.App.Name)

		file, err := ctx.Flags.String("file")
		if err != nil {
			return err
		}
		b := file.Value()
		fmt.Printf("App Flag `file`: `%s`\n", b)

		fmt.Printf("App Flag `int`: `%v`\n", ctx.Flags.Any("int"))

		verbose, err := ctx.Flags.Bool("verbose")
		if err != nil {
			return err
		}
		v, err := verbose.Value()
		if err != nil {
			return err
		}
		fmt.Printf("App Flag `verbose`: `%v`\n", v)

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
