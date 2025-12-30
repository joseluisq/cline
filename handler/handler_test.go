package handler

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joseluisq/cline/app"
	"github.com/joseluisq/cline/flag"
)

func TestHandler_Run(t *testing.T) {
	tests := []struct {
		name    string
		ap      *app.App
		vargs   []string
		wantErr error
	}{
		{
			name:  "should handle no arguments",
			ap:    &app.App{},
			vargs: []string{},
		},
		{
			name:    "should return error for unknown flag",
			ap:      &app.App{Flags: []flag.Flag{}},
			vargs:   []string{"app", "--unknown"},
			wantErr: errors.New("error: unknown flag '--unknown' argument"),
		},
		{
			name:  "should treat unknown command as tail argument",
			ap:    &app.App{Commands: []app.Cmd{}},
			vargs: []string{"app", "notacmd"},
		},
		{
			name:  "should trigger help on --help flag",
			ap:    &app.App{},
			vargs: []string{"app", "--help"},
		},
		{
			name: "should trigger command help when --help flag is used with a command",
			ap: &app.App{
				Commands: []app.Cmd{
					{
						Name:    "info",
						Summary: "Show command information",
					},
				},
			},
			vargs: []string{"app", "info", "--help"},
		},
		{
			name:  "should trigger version on --version flag",
			ap:    &app.App{},
			vargs: []string{"app", "--version"},
		},
		{
			name: "should return error for flag with invalid int value",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagInt{Name: "num"},
				},
			},
			vargs:   []string{"app", "--num", "notanint"},
			wantErr: errors.New("error: invalid integer value for flag '--num'"),
		},
		{
			name: "should parse valid int flag value",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagInt{Name: "num"},
				},
			},
			vargs: []string{"app", "--num", "42"},
		},
		{
			name: "should handle potentially dangerous input as a string",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "input"},
				},
			},
			vargs: []string{"app", "--input", "; rm -rf /"},
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
			vargs: []string{"app", "run", "foo", "bar"},
		},
		{
			name: "should handle SQL injection attempt as a string",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "query"},
				},
			},
			vargs: []string{"app", "--query", "'; DROP TABLE users; --"},
		},
		{
			name: "should handle XSS attempt as a string",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "name"},
				},
			},
			vargs: []string{"app", "--name", "<script>alert('xss')</script>"},
		},
		{
			name: "should return error when argument count exceeds limit",
			ap:   &app.App{},
			vargs: func() []string {
				// Create a slice with more arguments than the allowed maximum
				args := make([]string, maxArgsCount+2)
				args[0] = "app"
				for i := 1; i < len(args); i++ {
					args[i] = "arg"
				}
				return args
			}(),
			wantErr: fmt.Errorf("error: number of arguments exceeds the limit of %d", maxArgsCount),
		},
		{
			name: "should return error when an argument length exceeds limit",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "long-arg"},
				},
			},
			vargs:   []string{"app", "--long-arg", string(make([]byte, 4097))},
			wantErr: fmt.Errorf("error: argument exceeds maximum length of %d characters", maxArgLen),
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
			vargs: []string{"app", "--verbose", "--", "--another-flag", "-f"},
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
			vargs: []string{"app", "--verbose", "my-command"},
		},
		{
			name: "should treat non-boolean argument after bool flag as tail arg with a command",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagBool{Name: "verbose"},
				},
				Commands: []app.Cmd{
					{Name: "start"},
				},
			},
			vargs: []string{"app", "--verbose", "start"},
		},
		{
			name: "should treat non-boolean numeric value after bool flag as tail arg",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagBool{Name: "enabled"},
				},
				Handler: func(ctx *app.AppContext) error {
					enabledFlag, _ := ctx.Flags().Bool("enabled")
					val, _ := enabledFlag.Value()
					// The flag itself should be true because it was present.
					assert.True(t, val, "enabled flag should be true")
					// The non-boolean value "2" should become a tail argument.
					assert.Equal(t, 1, len(ctx.TailArgs()), "should have one tail argument")
					assert.Equal(t, "2", ctx.TailArgs()[0])
					return nil
				},
			},
			vargs: []string{"app", "--enabled", "2"},
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
			vargs:   []string{"app", "---"},
			wantErr: fmt.Errorf("error: flag '-' contains invalid characters"),
		},
		{
			name: "should handle path traversal attempt as a string",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "file"},
				},
			},
			vargs: []string{"app", "--file", "../../../etc/passwd"},
		},
		{
			name:    "should return error when non-UTF-8 characters in flag name",
			ap:      &app.App{},
			vargs:   []string{"app", "--\xff\xfe\xfd"},
			wantErr: fmt.Errorf("error: argument contains invalid UTF-8 characters"),
		},
		{
			name:    "should return error when non-ASCII characters in flag name",
			ap:      &app.App{},
			vargs:   []string{"app", "--ñame"},
			wantErr: fmt.Errorf("error: flag 'ñame' contains invalid characters"),
		},
		{
			name:    "should return error when non-ASCII characters in flag name with single dash",
			ap:      &app.App{},
			vargs:   []string{"app", "-–file"},
			wantErr: fmt.Errorf("error: flag '–file' contains invalid characters"),
		},
		{
			name: "should return error when non-UTF-8 characters in flag value",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "input"},
				},
			},
			vargs:   []string{"app", "--input", "value-\xff\xfe\xfd"},
			wantErr: fmt.Errorf("error: argument contains invalid UTF-8 characters"),
		},
		{
			name: "should handle null byte in argument",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "input"},
				},
			},
			vargs: []string{"app", "--input", "some\x00value"},
		},
		{
			name: "should return error when invalid UTF-8 sequence as a tail argument",
			ap: &app.App{
				Handler: func(ctx *app.AppContext) error {
					assert.Equal(t, 1, len(ctx.TailArgs()), "should have one tail argument")
					assert.Equal(t, "arg-\xff\xfe\xfd", ctx.TailArgs()[0])
					return nil
				},
			},
			vargs:   []string{"app", "arg-\xff\xfe\xfd"},
			wantErr: fmt.Errorf("error: argument contains invalid UTF-8 characters"),
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
			vargs: []string{"app", "--items", "a,b", "--items", "c,d"},
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
			vargs: []string{"app", "--items", ""},
		},
		{
			name: "should return error if flag requiring value is followed by another flag",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "input"},
					flag.FlagBool{Name: "verbose"},
				},
			},
			vargs:   []string{"app", "--input", "--verbose"},
			wantErr: fmt.Errorf("error: flag '--input' requires a value"),
		},
		{
			name: "should return error if string flag is last argument",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "input"},
				},
			},
			vargs:   []string{"app", "--input"},
			wantErr: fmt.Errorf("error: flag '--input' requires a value"),
		},
		{
			name: "should return error if int flag is last argument",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagInt{Name: "num"},
				},
			},
			vargs:   []string{"app", "--num"},
			wantErr: fmt.Errorf("error: flag '--num' requires a value"),
		},
		{
			name: "should return error if string slice flag is last argument",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagStringSlice{Name: "items"},
				},
			},
			vargs:   []string{"app", "--items"},
			wantErr: fmt.Errorf("error: flag '--items' requires a value"),
		},
		{
			name: "should return error for inconsistent special flag version alias",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagStringSlice{Name: "items"},
				},
			},
			vargs:   []string{"app", "--v"},
			wantErr: fmt.Errorf("error: unknown flag '--v' argument"),
		},
		{
			name: "should execute when consistent special flag version alias",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagStringSlice{Name: "items"},
				},
			},
			vargs: []string{"app", "-v"},
		},
		{
			name: "should return error for inconsistent special flag help alias",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagStringSlice{Name: "items"},
				},
			},
			vargs:   []string{"app", "--h"},
			wantErr: fmt.Errorf("error: unknown flag '--h' argument"),
		},
		{
			name: "should execute when consistent special flag help alias",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagStringSlice{Name: "items"},
				},
			},
			vargs: []string{"app", "-h"},
		},
		{
			name: "should skip when unsupported argument is provided",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagStringSlice{Name: "items"},
				},
			},
			vargs: []string{"app", "-"},
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
			vargs: []string{"app", "--input", "first-value", "another-value"},
		},
		{
			name: "should treat subsequent value as tail arg if string slice flag already assigned",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagStringSlice{Name: "items"},
				},
				Handler: func(ctx *app.AppContext) error {
					itemsFlag, _ := ctx.Flags().StringSlice("items")
					val := itemsFlag.Value()
					assert.Equal(t, []string{"a", "b"}, val)
					assert.Equal(t, []string{"another-value"}, ctx.TailArgs(), "should have one tail argument")
					return nil
				},
			},
			vargs: []string{"app", "--items", "a,b", "another-value"},
		},
		{
			name: "should parse command-specific bool flag when command is present",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagBool{Name: "global"},
				},
				Commands: []app.Cmd{
					{
						Name: "run",
						Flags: []flag.Flag{
							flag.FlagBool{Name: "local"},
						},
						Handler: func(ctx *app.CmdContext) error {
							localFlag, err := ctx.Flags.Bool("local")
							assert.NoError(t, err)
							assert.True(t, localFlag.IsProvided(), "local flag should be provided")
							return nil
						},
					},
				},
			},
			vargs: []string{"app", "run", "--local"},
		},
		{
			name: "should return error for command with invalid ASCII characters",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagBool{Name: "global"},
				},
				Commands: []app.Cmd{
					{
						Name: "ràn",
						Flags: []flag.Flag{
							flag.FlagBool{Name: "local"},
						},
					},
				},
			},
			vargs:   []string{"app", "run", "--local"},
			wantErr: fmt.Errorf("error: command 'ràn' contains invalid characters"),
		},
		{
			name: "should parse command-specific string flag when command is present",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{Name: "global"},
				},
				Commands: []app.Cmd{
					{
						Name: "run",
						Flags: []flag.Flag{
							flag.FlagString{Name: "local"},
						},
						Handler: func(ctx *app.CmdContext) error {
							localFlag, err := ctx.Flags.String("local")
							assert.NoError(t, err)
							assert.True(t, localFlag.IsProvided(), "local flag should be provided")
							return nil
						},
					},
				},
			},
			vargs: []string{"app", "run", "--local", "value"},
		},
		{
			name: "should parse command-specific string slice flag when command is present",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagStringSlice{Name: "global"},
				},
				Commands: []app.Cmd{
					{
						Name: "run",
						Flags: []flag.Flag{
							flag.FlagStringSlice{Name: "local"},
						},
						Handler: func(ctx *app.CmdContext) error {
							localFlag, err := ctx.Flags.StringSlice("local")
							assert.NoError(t, err)
							assert.True(t, localFlag.IsProvided(), "local flag should be provided")
							return nil
						},
					},
				},
			},
			vargs: []string{"app", "run", "--local", "value1,value2"},
		},
		{
			name: "should parse command-specific int flag when command is present",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagInt{Name: "global"},
				},
				Commands: []app.Cmd{
					{
						Name: "run",
						Flags: []flag.Flag{
							flag.FlagInt{Name: "local"},
						},
						Handler: func(ctx *app.CmdContext) error {
							localFlag, err := ctx.Flags.Int("local")
							assert.NoError(t, err)
							assert.True(t, localFlag.IsProvided(), "local flag should be provided")
							return nil
						},
					},
				},
			},
			vargs: []string{"app", "run", "--local", "1"},
		},
		{
			name: "should correctly identify flag provided as alias",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagString{
						Name:    "file",
						Aliases: []string{"f"},
					},
				},
				Handler: func(ctx *app.AppContext) error {
					f, err := ctx.Flags().String("file")
					assert.NoError(t, err)
					assert.True(t, f.IsProvided(), "flag should be marked as provided")
					assert.True(t, f.IsProvidedShort(), "flag should be marked as provided via alias (short)")
					assert.False(t, f.IsProvidedLong(), "flag should not be marked as provided via long name")
					return nil
				},
			},
			vargs: []string{"app", "-f", "somefile.txt"},
		},
		{
			name: "should treat subsequent value as tail arg if int flag already assigned",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagInt{Name: "num"},
				},
				Handler: func(ctx *app.AppContext) error {
					numFlag, _ := ctx.Flags().Int("num")
					val, _ := numFlag.Value()
					assert.Equal(t, 123, val)
					assert.Equal(t, []string{"456"}, ctx.TailArgs(), "should have one tail argument")
					return nil
				},
			},
			vargs: []string{"app", "--num", "123", "456"},
		},
		{
			name: "should treat subsequent non-flag arg as tail arg after a bool flag",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagBool{Name: "verbose"},
				},
				Handler: func(ctx *app.AppContext) error {
					verboseFlag, _ := ctx.Flags().Bool("verbose")
					val, _ := verboseFlag.Value()

					assert.True(t, val, "verbose flag should be true")
					assert.True(t, verboseFlag.Flag.FlagAssigned, "verbose flag should be assigned")
					assert.Equal(t, []string{"extra-arg"}, ctx.TailArgs(), "should have one tail argument")
					return nil
				},
			},
			vargs: []string{"app", "--verbose", "true", "extra-arg"},
		},
		{
			name: "should return error for invalid flag",
			ap: &app.App{
				Flags: []flag.Flag{
					flag.FlagStringSlice{Name: "itéms", Aliases: []string{"i"}},
				},
			},
			vargs:   []string{"app", "-i"},
			wantErr: fmt.Errorf("error: flag 'itéms' contains invalid characters"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actualErr := New(tt.ap).Run(tt.vargs); tt.wantErr != nil {
				assert.Error(t, actualErr, "Expected an error but got none")
				assert.Equal(t, actualErr, tt.wantErr, "Error message does not match the expected one")
			} else {
				assert.NoError(t, actualErr, "Expected no error but got one")
			}
		})
	}
}
