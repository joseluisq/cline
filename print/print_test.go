package print_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joseluisq/cline/app"
	"github.com/joseluisq/cline/flag"
	"github.com/joseluisq/cline/print"
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

func Test_printHelp(t *testing.T) {
	type args struct {
		app *app.App
		cmd *app.Cmd
	}
	ap := NewApp(nil, nil)
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
				app: ap,
			},
		},
		{
			name: "valid command output",
			args: args{
				app: ap,
				cmd: &ap.Commands[0],
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := print.PrintHelp(tt.args.app, tt.args.cmd); tt.wantErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got one")
			}
		})
	}
}

func TestApp_printVersion(t *testing.T) {
	tests := []struct {
		name    string
		app     *app.App
		wantErr bool
	}{
		{
			name: "valid app for version output",
			app:  NewApp(nil, nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				assert.Fail(t, "Failed to create pipe for stdout capture", err)
			}
			os.Stdout = w

			defer func() { os.Stdout = oldStdout }()
			defer w.Close()

			tt.app.PrintVersion()

			out := make([]byte, 1024)
			defer r.Close()
			if _, err = r.Read(out); err != nil {
				assert.Fail(t, "Failed to read pipe for stdout capture", err)
			}

			str := string(out)
			assert.Contains(t, str, "Version:")
			assert.Contains(t, str, "Go version:")
			assert.Contains(t, str, "Built:")
			assert.Contains(t, str, "Commit:")
		})
	}
}
