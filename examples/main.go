package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/joseluisq/cline"
)

func main() {
	app := cli.App{
		Name:    "enve",
		Summary: "run a program in a modified environment using .env files",
		Flags: []*cli.Flag{
			{
				Name:      "file",
				Summary:   "load environment variables from a file path",
				Value:     ".env",
				Shortcuts: []string{"f"},
			},
			{
				Name:      "verbose",
				Summary:   "load environment variables from a file path",
				Value:     false,
				Shortcuts: []string{"v"},
			},
		},
		Commands: []*cli.Cmd{
			{
				Name:    "info",
				Summary: "show command information",
				Flags: []*cli.Flag{
					{
						Name:      "version",
						Summary:   "enable more verbose command information",
						Value:     10,
						Shortcuts: []string{"z"},
					},
					{
						Name:      "detailed",
						Summary:   "enable info details",
						Value:     true,
						Shortcuts: []string{"d"},
					},
				},
				Handler: func(ctx *cli.CmdContext) error {
					fmt.Printf("Cmd `%s` executed!\n", ctx.Cmd.Name)

					i, err := ctx.Flags.Int("version")

					if err != nil {
						log.Fatalln(err)
					}

					fmt.Printf("Cmd Flag `version` opted: `%d` (%T)\n", i, i)
					fmt.Printf("Cmd Tail arguments: %#v\n", ctx.TailArgs)
					return nil
				},
			},
		},
		Handler: func(ctx *cli.AppContext) error {
			fmt.Printf("Application `%s` executed!\n", ctx.App.Name)
			fmt.Printf("Application Tail arguments: %#v\n", ctx.TailArgs)
			fmt.Printf("Application Flag `file` opted: `%s`\n", ctx.Flags.String("file"))

			return nil
		},
	}

	err := app.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
