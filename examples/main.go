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
				fmt.Printf("App provided flags: %s\n", ctx.AppContext.Flags.ProvidedFlags())
				fmt.Printf("Cmd provided flags: %s\n", ctx.Flags.ProvidedFlags())

				fmt.Printf("App Flag `file`: `%s`\n", ctx.AppContext.Flags.StringSlice("file"))

				i, err := ctx.Flags.Int("trace")
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Printf("Cmd Flag `trace`: `%d` (%T)\n", i, i)

				d := ctx.Flags.String("detailed")
				fmt.Printf("Cmd Flag `detailed`: `%s` (%T)\n", d, d)
				fmt.Printf("Cmd Tail arguments: %#v\n", ctx.TailArgs)
				return nil
			},
		},
	}
	app.Handler = func(ctx *cli.AppContext) error {
		fmt.Printf("App `%s` executed!\n", ctx.App.Name)
		fmt.Printf("App provided flags: %s\n", ctx.Flags.ProvidedFlags())
		fmt.Printf("App Flag `file`: `%s`\n", ctx.Flags.String("file"))
		fmt.Printf("App Flag `int`: `%v`\n", ctx.Flags.String("int"))
		b, _ := ctx.Flags.Bool("verbose")
		fmt.Printf("App Flag `verbose`: `%v`\n", b)
		fmt.Printf("App Tail arguments: %#v\n", ctx.TailArgs)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
