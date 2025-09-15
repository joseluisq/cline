package testsuite

import (
	"github.com/joseluisq/cline/app"
	"github.com/joseluisq/cline/flag"
)

func NewApp(appHandler app.AppHandler, cmdHandler app.CmdHandler) *app.App {
	ap := app.New()
	ap.Name = "enve"
	ap.Summary = "Run a program in a modified environment using .env files"
	ap.Version = "0.0.0"
	ap.BuildTime = ""
	ap.Flags = []flag.Flag{
		flag.FlagString{
			Name:    "AA",
			Value:   ".env",
			Aliases: []string{"a"},
		},
		flag.FlagBool{
			Name:    "BB",
			Value:   false,
			Aliases: []string{"b"},
		},
		flag.FlagInt{
			Name:    "CC",
			Aliases: []string{"c"},
		},
		flag.FlagStringSlice{
			Name:    "DD",
			Value:   nil,
			Aliases: []string{"d"},
		},
	}
	ap.Commands = []app.Cmd{
		{
			Name:    "info",
			Summary: "Show command information",
			Flags: []flag.Flag{
				flag.FlagInt{
					Name:    "GG",
					Aliases: []string{"g"},
				},
				flag.FlagBool{
					Name:    "ZZ",
					Value:   false,
					Aliases: []string{"z"},
				},
				flag.FlagString{
					Name:    "FF",
					Summary: "abcde",
					Value:   ".env",
					Aliases: []string{"f"},
					EnvVar:  "FF_ENV_VAR",
				},
				flag.FlagStringSlice{
					Name:    "II",
					Value:   []string{"q", "r", "s"},
					Aliases: []string{"i"},
				},
			},
			Handler: func(ctx *app.CmdContext) error {
				if cmdHandler != nil {
					return cmdHandler(ctx)
				}
				return nil
			},
		},
	}
	ap.Handler = func(ctx *app.AppContext) error {
		if appHandler != nil {
			return appHandler(ctx)
		}
		return nil
	}
	return ap
}
