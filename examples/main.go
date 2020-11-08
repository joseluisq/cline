package main

import (
	"fmt"
	"os"

	cli "github.com/joseluisq/cline"
)

func main() {
	app := cli.New()
	app.Name = "enve"
	app.Summary = "Run a program in a modified environment using .env files"
	app.Flags = []cli.Flag{
		cli.FlagString{
			Name:    "file",
			Summary: "load environment variables from a file path",
			Value:   ".env",
			Aliases: []string{"f"},
		},
		cli.FlagBool{
			Name:    "verbose",
			Summary: "enable more verbose info",
			Value:   false,
			Aliases: []string{"v"},
			EnvVar:  "ENV_VERBOSE",
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
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
