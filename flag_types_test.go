package cline

import (
	"testing"
)

func TestFlagInt_initialize(t *testing.T) {
	type fields struct {
		Name          string
		Summary       string
		Value         int
		Aliases       []string
		EnvVar        string
		zflag         FlagValue
		zflagAssigned bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fi := &FlagInt{
				Name:          tt.fields.Name,
				Summary:       tt.fields.Summary,
				Value:         tt.fields.Value,
				Aliases:       tt.fields.Aliases,
				EnvVar:        tt.fields.EnvVar,
				zflag:         tt.fields.zflag,
				zflagAssigned: tt.fields.zflagAssigned,
			}
			fi.initialize()
		})
	}
}

func TestFlagBool_initialize(t *testing.T) {
	type fields struct {
		Name          string
		Summary       string
		Value         bool
		Aliases       []string
		EnvVar        string
		zflag         FlagValue
		zflagAssigned bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb := &FlagBool{
				Name:          tt.fields.Name,
				Summary:       tt.fields.Summary,
				Value:         tt.fields.Value,
				Aliases:       tt.fields.Aliases,
				EnvVar:        tt.fields.EnvVar,
				zflag:         tt.fields.zflag,
				zflagAssigned: tt.fields.zflagAssigned,
			}
			fb.initialize()
		})
	}
}

func TestFlagString_initialize(t *testing.T) {
	type fields struct {
		Name          string
		Summary       string
		Value         string
		Aliases       []string
		EnvVar        string
		zflag         FlagValue
		zflagAssigned bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagString{
				Name:          tt.fields.Name,
				Summary:       tt.fields.Summary,
				Value:         tt.fields.Value,
				Aliases:       tt.fields.Aliases,
				EnvVar:        tt.fields.EnvVar,
				zflag:         tt.fields.zflag,
				zflagAssigned: tt.fields.zflagAssigned,
			}
			fs.initialize()
		})
	}
}

func TestFlagStringSlice_initialize(t *testing.T) {
	type fields struct {
		Name          string
		Summary       string
		Value         []string
		Aliases       []string
		EnvVar        string
		zflag         FlagValue
		zflagAssigned bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlagStringSlice{
				Name:          tt.fields.Name,
				Summary:       tt.fields.Summary,
				Value:         tt.fields.Value,
				Aliases:       tt.fields.Aliases,
				EnvVar:        tt.fields.EnvVar,
				zflag:         tt.fields.zflag,
				zflagAssigned: tt.fields.zflagAssigned,
			}
			fs.initialize()
		})
	}
}
