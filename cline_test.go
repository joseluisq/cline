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
		// TODO: Add test cases.
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
	type fields struct {
		Name      string
		Summary   string
		Version   string
		BuildTime string
		Flags     []Flag
		Commands  []Cmd
		Handler   AppHandler
	}
	type args struct {
		vArgs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				Name:      tt.fields.Name,
				Summary:   tt.fields.Summary,
				Version:   tt.fields.Version,
				BuildTime: tt.fields.BuildTime,
				Flags:     tt.fields.Flags,
				Commands:  tt.fields.Commands,
				Handler:   tt.fields.Handler,
			}
			if err := app.Run(tt.args.vArgs); (err != nil) != tt.wantErr {
				t.Errorf("App.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
