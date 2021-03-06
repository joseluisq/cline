// Package cline is a fast and lightweight CLI package for Go.
//
package cline

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *App
	}{
		{
			name: "valid instance",
			want: &App{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_Run(t *testing.T) {
	appEmptyFlags := newApp(nil, nil)
	for i, v := range appEmptyFlags.Flags {
		switch f := v.(type) {
		case FlagBool:
			f.Name = ""
			appEmptyFlags.Flags[i] = f
		}
	}
	appEmptyCommmands := newApp(nil, nil)
	for i, c := range appEmptyFlags.Commands {
		c.Name = ""
		appEmptyCommmands.Commands[i] = c
	}
	appInvalidHandlers := newApp(nil, nil)
	appInvalidHandlers.Handler = nil
	appInvalidHandlers.Commands[0].Handler = nil

	type args struct {
		vArgs []string
	}
	tests := []struct {
		name    string
		app     *App
		args    args
		wantErr bool
	}{
		{
			name: "not recognized argument",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "--unknown", "123"},
			},
			wantErr: true,
		},
		{
			name: "run an app instance and flags",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "--AA", "sdev.env", "--BB", "--CC", "22", "-d", "xyz"},
			},
		},
		{
			name: "run an app instance and command flags",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{
					"", "--AA", "sdev.env", "--BB", "-c", "22", "-d", "2,2",
					"info", "--ZZ", "--II", "1,2,3",
				},
			},
		},
		{
			name: "run an app instance and command flags with tail args",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{
					"", "info", "--FF", ".env2", "-g", "11", "--ZZ", "-i", "a,b", "sdasdas",
				},
			},
		},
		{
			name: "run version flag",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "--version"},
			},
		},
		{
			name: "run help flag",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "--help"},
			},
		},
		{
			name: "run command help flag",
			app:  newApp(nil, nil),
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
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "--ZZ"},
			},
		},
		{
			name: "run valid command bool flag long",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "--ZZ", "true"},
			},
		},
		{
			name: "run invalid command bool flag long",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "--ZZ", "1f9"},
			},
		},
		{
			name: "run command tail args",
			app:  newApp(nil, nil),
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
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "---unknown", "123"},
			},
		},
		{
			name: "run valid bool flag short",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "--BB"},
			},
		},
		{
			name: "run valid bool flag long",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "--BB", "false"},
			},
		},
		{
			name: "run valid int flag",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "--CC", "2"},
			},
		},
		{
			name: "run invalid int flag value",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "--CC", "s1f0"},
			},
			wantErr: true,
		},
		{
			name: "run valid string slice flag",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "-d", "2,4,5"},
			},
		},
		{
			name: "run valid command string slice flag",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "--II", "2,4,5"},
			},
		},
		{
			name: "run valid command string slice and bool flags",
			app:  newApp(nil, nil),
			args: args{
				vArgs: []string{"", "info", "--II", "2,4,5", "-z"},
			},
		},
		{
			name: "run valid command int flag",
			app:  newApp(nil, nil),
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
			app := tt.app
			if err := app.Run(tt.args.vArgs); (err != nil) != tt.wantErr {
				t.Errorf("App.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
