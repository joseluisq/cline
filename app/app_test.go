package app_test

import (
	"reflect"
	"testing"

	"github.com/joseluisq/cline/app"
	"github.com/joseluisq/cline/flag"
	"github.com/joseluisq/cline/handler"
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

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *app.App
	}{
		{
			name: "valid instance",
			want: &app.App{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := app.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_Run(t *testing.T) {
	appEmptyFlags := NewApp(nil, nil)
	for i, v := range appEmptyFlags.Flags {
		switch f := v.(type) {
		case flag.FlagBool:
			f.Name = ""
			appEmptyFlags.Flags[i] = f
		}
	}
	appEmptyCommmands := NewApp(nil, nil)
	for i, c := range appEmptyFlags.Commands {
		c.Name = ""
		appEmptyCommmands.Commands[i] = c
	}
	appInvalidHandlers := NewApp(nil, nil)
	appInvalidHandlers.Handler = nil
	appInvalidHandlers.Commands[0].Handler = nil

	type args struct {
		vArgs []string
	}
	tests := []struct {
		name    string
		app     *app.App
		args    args
		wantErr bool
	}{
		{
			name: "not recognized argument",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "--unknown", "123"},
			},
			wantErr: true,
		},
		{
			name: "run an app instance and flags",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "--AA", "sdev.env", "--BB", "--CC", "22", "-d", "xyz"},
			},
		},
		{
			name: "run an app instance and command flags",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{
					"", "--AA", "sdev.env", "--BB", "-c", "22", "-d", "2,2",
					"info", "--ZZ", "--II", "1,2,3",
				},
			},
		},
		{
			name: "run an app instance and command flags with tail args",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{
					"", "info", "--FF", ".env2", "-g", "11", "--ZZ", "-i", "a,b", "sdasdas",
				},
			},
		},
		{
			name: "run version flag",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "--version"},
			},
		},
		{
			name: "run help flag",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "--help"},
			},
		},
		{
			name: "run command help flag",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "--help"},
			},
		},
		{
			name: "run empty flag",
			app:  appEmptyFlags,
			args: args{
				vArgs: nil,
			},
			wantErr: true,
		},
		{
			name: "run valid command bool flag short",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "--ZZ"},
			},
		},
		{
			name: "run valid command bool flag long",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "--ZZ", "true"},
			},
		},
		{
			name: "run invalid command bool flag long",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "--ZZ", "1f9"},
			},
		},
		{
			name: "run command tail args",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "zzz", "yyy"},
			},
		},
		{
			name: "run empty commands",
			app:  appEmptyCommmands,
			args: args{
				vArgs: nil,
			},
			wantErr: true,
		},
		{
			name: "run invalid argument flags",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "---unknown", "123"},
			},
		},
		{
			name: "run valid bool flag short",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "--BB"},
			},
		},
		{
			name: "run valid bool flag long",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "--BB", "false"},
			},
		},
		{
			name: "run valid int flag",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "--CC", "2"},
			},
		},
		{
			name: "run invalid int flag value",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "--CC", "s1f0"},
			},
			wantErr: true,
		},
		{
			name: "run valid string slice flag",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "-d", "2,4,5"},
			},
		},
		{
			name: "run valid command string slice flag",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "--II", "2,4,5"},
			},
		},
		{
			name: "run valid command string slice and bool flags",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "--II", "2,4,5", "-z"},
			},
		},
		{
			name: "run valid command int flag",
			app:  NewApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "-g", "0", "-z"},
			},
		},
		{
			name: "run null handlers",
			app:  appInvalidHandlers,
			args: args{
				vArgs: []string{"", "info", "-g", "0", "-z"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handler.New(tt.app).Run(tt.args.vArgs); (err != nil) != tt.wantErr {
				t.Errorf("App.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
