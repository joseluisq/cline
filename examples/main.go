package main

import (
	"fmt"
	"os"

	"github.com/joseluisq/cline/app"
	"github.com/joseluisq/cline/flag"
	"github.com/joseluisq/cline/handler"
)

// App version and build time values passed at compile time
// See `Makefile` > build
var (
	version     string = "devel"
	buildTime   string
	buildCommit string
)

func main() {
	ap := app.New()
	ap.Name = "enve"
	ap.Summary = "Run a program in a modified environment using .env files"
	ap.Version = version
	ap.BuildTime = buildTime
	ap.BuildCommit = buildCommit
	ap.Flags = []flag.Flag{
		flag.FlagString{
			Name:    "file",
			Summary: "Load environment variables from a file path.\nSome new line description\nAnother new line description.",
			Value:   ".env",
			Aliases: []string{"f"},
		},
		flag.FlagInt{
			Name:    "int",
			Summary: "Int value",
			Value:   5,
			Aliases: []string{"b", "z"},
		},
		flag.FlagBool{
			Name:    "verbose",
			Summary: "Enable more verbose info",
			Value:   true,
			Aliases: []string{"V"},
			EnvVar:  "ENV_VERBOSE",
		},
	}
	ap.Commands = []app.Cmd{
		{
			Name:    "info",
			Summary: "Show command information",
			Flags: []flag.Flag{
				flag.FlagInt{
					Name:    "trace",
					Summary: "Enable tracing mode",
					Value:   10,
					Aliases: []string{"t"},
				},
				flag.FlagBool{
					Name:    "detailed",
					Summary: "Enable info details",
					Value:   true,
					Aliases: []string{"d"},
				},
				flag.FlagString{
					Name:    "FF",
					Value:   ".env",
					Aliases: []string{"f"},
				},
				flag.FlagStringSlice{
					Name:    "II",
					Value:   []string{"q", "r", "s"},
					Aliases: []string{"i"},
				},
			},
			Handler: func(ctx *app.CmdContext) error {
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
	ap.Handler = func(ctx *app.AppContext) error {
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

	if err := handler.New(ap).Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
