package cline

import (
	"testing"
)

func newApp(appHandler func(ctx *AppContext) error, cmdHandler func(ctx *CmdContext) error) (app *App) {
	app = New()
	app.Name = "enve"
	app.Summary = "Run a program in a modified environment using .env files"
	app.Version = "0.0.0"
	app.BuildTime = ""
	app.Flags = []Flag{
		FlagString{
			Name:    "AA",
			Value:   ".env",
			Aliases: []string{"a"},
		},
		FlagBool{
			Name:    "BB",
			Value:   false,
			Aliases: []string{"b"},
		},
		FlagInt{
			Name:    "CC",
			Aliases: []string{"c"},
		},
		FlagStringSlice{
			Name:    "DD",
			Value:   nil,
			Aliases: []string{"d"},
		},
	}
	app.Commands = []Cmd{
		{
			Name:    "info",
			Summary: "Show command information",
			Flags: []Flag{
				FlagInt{
					Name:    "GG",
					Aliases: []string{"g"},
				},
				FlagBool{
					Name:    "ZZ",
					Value:   false,
					Aliases: []string{"z"},
				},
				FlagString{
					Name:    "FF",
					Value:   ".env",
					Aliases: []string{"f"},
					EnvVar:  "FF_ENV_VAR",
				},
				FlagStringSlice{
					Name:    "II",
					Value:   []string{"q", "r", "s"},
					Aliases: []string{"i"},
				},
			},
			Handler: func(ctx *CmdContext) error {
				if cmdHandler != nil {
					return cmdHandler(ctx)
				}
				return nil
			},
		},
	}
	app.Handler = func(ctx *AppContext) error {
		if appHandler != nil {
			return appHandler(ctx)
		}
		return nil
	}
	return app
}

func Test_printHelp(t *testing.T) {
	type args struct {
		app *App
		cmd *Cmd
	}
	app := newApp(nil, nil)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "invalid command app for output",
			args:    args{},
			wantErr: true,
		},
		{
			name: "valid global output",
			args: args{
				app: app,
			},
		},
		{
			name: "valid command output",
			args: args{
				app: app,
				cmd: &app.Commands[0],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printHelp(tt.args.app, tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("printHelp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_printVersion(t *testing.T) {
	app := newApp(nil, nil)
	tests := []struct {
		name    string
		app     *App
		wantErr bool
	}{
		{
			name:    "invalid app for version output",
			app:     nil,
			wantErr: true,
		},
		{
			name: "valid app for version output",
			app:  app,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.app.printVersion(); (err != nil) != tt.wantErr {
				t.Errorf("App.printVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
