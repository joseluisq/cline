package cline

import (
	"testing"
)

func newApp() (app *App) {
	app = New()
	app.Name = "enve"
	app.Summary = "Run a program in a modified environment using .env files"
	app.Version = "0.0.0"
	app.BuildTime = ""
	app.Flags = []Flag{
		FlagString{
			Name:    "A",
			Value:   ".env",
			Aliases: []string{"a"},
		},
		FlagBool{
			Name:    "B",
			Value:   false,
			Aliases: []string{"b"},
		},
		FlagInt{
			Name:    "C",
			Value:   32,
			Aliases: []string{"c"},
		},
		FlagStringSlice{
			Name:    "D",
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
					Name:    "E",
					Value:   10,
					Aliases: []string{"e"},
				},
				FlagString{
					Name:    "F",
					Aliases: []string{"f"},
				},
			},
		},
	}
	return app
}

func Test_printHelp(t *testing.T) {
	type args struct {
		app *App
		cmd *Cmd
	}
	app := newApp()
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
	app := newApp()
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
