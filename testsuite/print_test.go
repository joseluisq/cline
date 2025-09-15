package testsuite

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joseluisq/cline/app"
	"github.com/joseluisq/cline/print"
)

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
