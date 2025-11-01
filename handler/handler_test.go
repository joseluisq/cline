package handler_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joseluisq/cline/app"
	"github.com/joseluisq/cline/flag"
	"github.com/joseluisq/cline/handler"
)

const maxArgsCount = 1024

func TestHandler_Run(t *testing.T) {
	tests := []struct {
		name    string
		ap      *app.App
		vArgs   []string
		wantErr bool
	}{
		{
			name:    "should handle no arguments",
			ap:      &app.App{},
			vArgs:   []string{},
			wantErr: false,
		},
		{
			name:    "should return error for unknown flag",
			ap:      &app.App{Flags: []flag.Flag{}},
			vArgs:   []string{"app", "--unknown"},
			wantErr: true,
		},
		{
			name:    "should treat unknown command as tail argument",
			ap:      &app.App{Commands: []app.Cmd{}},
			vArgs:   []string{"app", "notacmd"},
			wantErr: false,
		},
		{
			name:    "should trigger help on --help flag",
			ap:      &app.App{},
			vArgs:   []string{"app", "--help"},
			wantErr: false,
		},
		{
			name:    "should trigger version on --version flag",
			ap:      &app.App{},
			vArgs:   []string{"app", "--version"},
			wantErr: false,
		},
		{
			name: "should return error for flag with invalid int value",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagInt{Name: "num"},
				},
			},
			vArgs:   []string{"app", "--num", "notanint"},
			wantErr: true,
		},
		{
			name: "should parse valid int flag value",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagInt{Name: "num"},
				},
			},
			vArgs:   []string{"app", "--num", "42"},
			wantErr: false,
		},
		{
			name: "should handle potentially dangerous input as a string",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "input"},
				},
			},
			vArgs:   []string{"app", "--input", "; rm -rf /"},
			wantErr: false,
		},
		{
			name: "should pass tail arguments to command handler",
			ap: &app.App{
				Commands: []app.Cmd{
					{
						Name: "run",
						Handler: func(ctx *app.CmdContext) error {
							if len(ctx.TailArgs) != 2 || ctx.TailArgs[0] != "foo" || ctx.TailArgs[1] != "bar" {
								return fmt.Errorf("unexpected tail args")
							}
							return nil
						},
					},
				},
			},
			vArgs:   []string{"app", "run", "foo", "bar"},
			wantErr: false,
		},
		{
			name: "should handle SQL injection attempt as a string",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "query"},
				},
			},
			vArgs:   []string{"app", "--query", "'; DROP TABLE users; --"},
			wantErr: false,
		},
		{
			name: "should handle XSS attempt as a string",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "name"},
				},
			},
			vArgs:   []string{"app", "--name", "<script>alert('xss')</script>"},
			wantErr: false,
		},
		{
			name: "should return error when argument count exceeds limit",
			ap:   &app.App{},
			vArgs: func() []string {
				// Create a slice with more arguments than the allowed maximum
				args := make([]string, maxArgsCount+2)
				args[0] = "app"
				for i := 1; i < len(args); i++ {
					args[i] = "arg"
				}
				return args
			}(),
			wantErr: true,
		},
		{
			name:    "should return error when an argument length exceeds limit",
			ap:      &app.App{},
			vArgs:   []string{"app", "--long-arg", string(make([]byte, 4097))},
			wantErr: true,
		},
		{
			name: "should treat arguments after -- as tail arguments",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagBool{Name: "verbose"},
				},
				Handler: func(ctx *app.AppContext) error {
					assert.Equal(t, 2, len(ctx.TailArgs()), "should have two tail arguments")
					assert.Equal(t, "--another-flag", ctx.TailArgs()[0])
					assert.Equal(t, "-f", ctx.TailArgs()[1])
					return nil
				},
			},
			vArgs:   []string{"app", "--verbose", "--", "--another-flag", "-f"},
			wantErr: false,
		},
		{
			name: "should not consume next argument for boolean flag if not a bool value",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagBool{Name: "verbose"},
				},
				Handler: func(ctx *app.AppContext) error {
					v, _ := ctx.Flags().Bool("verbose")
					val, _ := v.Value()
					assert.True(t, val, "verbose flag should be true")
					assert.Equal(t, 1, len(ctx.TailArgs()), "should have one tail argument")
					assert.Equal(t, "my-command", ctx.TailArgs()[0])
					return nil
				},
			},
			vArgs:   []string{"app", "--verbose", "my-command"},
			wantErr: false,
		},
		{
			name: "should handle malformed flags as tail arguments",
			ap: &app.App{
				Handler: func(ctx *app.AppContext) error {
					assert.Equal(t, 1, len(ctx.TailArgs()), "should have one tail argument")
					assert.Equal(t, "---", ctx.TailArgs()[0])
					return nil
				},
			},
			vArgs:   []string{"app", "---"},
			wantErr: false,
		},
		{
			name: "should handle path traversal attempt as a string",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "file"},
				},
			},
			vArgs:   []string{"app", "--file", "../../../etc/passwd"},
			wantErr: false,
		},
		{
			name:    "should handle non-UTF-8 characters in flag name",
			ap:      &app.App{},
			vArgs:   []string{"app", "--\xff\xfe\xfd"},
			wantErr: true,
		},
		{
			name: "should handle non-UTF-8 characters in flag value",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "input"},
				},
			},
			vArgs:   []string{"app", "--input", "value-\xff\xfe\xfd"},
			wantErr: false,
		},
		{
			name: "should handle null byte in argument",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "input"},
				},
			},
			vArgs:   []string{"app", "--input", "some\x00value"},
			wantErr: false,
		},
		{
			name: "should handle invalid UTF-8 sequence as a tail argument",
			ap: &app.App{
				Handler: func(ctx *app.AppContext) error {
					assert.Equal(t, 1, len(ctx.TailArgs()), "should have one tail argument")
					assert.Equal(t, "arg-\xff\xfe\xfd", ctx.TailArgs()[0])
					return nil
				},
			},
			vArgs:   []string{"app", "arg-\xff\xfe\xfd"},
			wantErr: false,
		},
		{
			name: "should overwrite string slice flag on multiple assignments",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagStringSlice{Name: "items"},
				},
				Handler: func(ctx *app.AppContext) error {
					val, _ := ctx.Flags().StringSlice("items")
					assert.Equal(t, []string{"c", "d"}, val.Value(), "should only contain the last value")
					return nil
				},
			},
			vArgs:   []string{"app", "--items", "a,b", "--items", "c,d"},
			wantErr: false,
		},
		{
			name: "should handle empty value for string slice flag",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagStringSlice{Name: "items"},
				},
				Handler: func(ctx *app.AppContext) error {
					val, _ := ctx.Flags().StringSlice("items")
					assert.Equal(t, []string{""}, val.Value(), "slice should contain a single empty string")
					return nil
				},
			},
			vArgs:   []string{"app", "--items", ""},
			wantErr: false,
		},
		{
			name: "should treat subsequent value as tail arg if flag already assigned",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "input"},
				},
				Handler: func(ctx *app.AppContext) error {
					assert.Equal(t, []string{"another-value"}, ctx.TailArgs(), "should have one tail argument")
					return nil
				},
			},
			vArgs:   []string{"app", "--input", "first-value", "another-value"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler.New(tt.ap)
			if err := h.Run(tt.vArgs); tt.wantErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got one")
			}
		})
	}
}
