package app

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joseluisq/cline/flag"
)

func TestAppContext_App(t *testing.T) {
	tests := []struct {
		name       string
		app        *App
		flagValues *flag.FlagValues
		tailArgs   []string
		want       *App
	}{
		{
			name: "should return the app instance",
			app: &App{
				Name: "test-app",
			},
			flagValues: &flag.FlagValues{},
			tailArgs:   []string{},
			want: &App{
				Name: "test-app",
			},
		},
		{
			name:       "should handle a nil app instance",
			app:        nil,
			flagValues: &flag.FlagValues{},
			tailArgs:   []string{},
			want:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewContext(tt.app, tt.flagValues, tt.tailArgs)
			got := c.App()
			assert.Equal(t, got, tt.want, "AppContext.App() = %v, want %v", got, tt.want)
		})
	}
}

func TestNewContext(t *testing.T) {
	type args struct {
		app        *App
		flagValues *flag.FlagValues
		tailArgs   []string
	}
	tests := []struct {
		name string
		args args
		want AppContext
	}{
		{
			name: "should create a new context with values",
			args: args{
				app:        &App{Name: "my-app"},
				flagValues: flag.NewFlagValues([]flag.Flag{}),
				tailArgs:   []string{"arg1", "arg2"},
			},
			want: AppContext{
				app:      &App{Name: "my-app"},
				flags:    flag.NewFlagValues([]flag.Flag{}),
				tailArgs: []string{"arg1", "arg2"},
			},
		},
		{
			name: "should create a new context with nil values",
			args: args{
				app:        nil,
				flagValues: nil,
				tailArgs:   nil,
			},
			want: AppContext{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewContext(tt.args.app, tt.args.flagValues, tt.args.tailArgs)
			assert.Equal(t, got.App(), tt.want.App())
			assert.Equal(t, got.Flags(), tt.want.Flags())
			assert.Equal(t, got.TailArgs(), tt.want.TailArgs())
		})
	}
}

func TestAppContext_Flags(t *testing.T) {
	flags := flag.NewFlagValues([]flag.Flag{flag.FlagBool{Name: "verbose"}})
	c := NewContext(nil, flags, nil)

	t.Run("should return the flag values", func(t *testing.T) {
		got := c.Flags()
		assert.Equal(t, got, flags)
	})
}

func TestAppContext_TailArgs(t *testing.T) {
	tailArgs := []string{"arg1", "arg2"}
	c := NewContext(nil, nil, tailArgs)

	t.Run("should return the tail arguments", func(t *testing.T) {
		got := c.TailArgs()
		assert.Equal(t, got, tailArgs)
	})

	t.Run("should return nil for nil tail arguments", func(t *testing.T) {
		cNil := NewContext(nil, nil, nil)
		got := cNil.TailArgs()
		assert.Nil(t, got, "TailArgs() should return nil for nil tail arguments")
	})
}
