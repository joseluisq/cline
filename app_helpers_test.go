package cline

import (
	"reflect"
	"testing"
)

func Test_validateCommands(t *testing.T) {
	type args struct {
		commands []Cmd
	}
	tests := []struct {
		name    string
		args    args
		want    []Cmd
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateCommands(tt.args.commands)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateCommands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateAndInitFlags(t *testing.T) {
	type args struct {
		flags []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    []Flag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateAndInitFlags(tt.args.flags)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateAndInitFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateAndInitFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
