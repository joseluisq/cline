package cline

import (
	"testing"
)

func Test_printHelp(t *testing.T) {
	type args struct {
		app *App
		cmd *Cmd
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
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
	type fields struct {
		Name      string
		Summary   string
		Version   string
		BuildTime string
		Flags     []Flag
		Commands  []Cmd
		Handler   AppHandler
	}
	tests := []struct {
		name   string
		fields fields
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
			app.printVersion()
		})
	}
}
