package cline

import (
	"os"
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
	// env variables for test purposes
	os.Setenv("ENV_INT_VAR_OK", "1")
	os.Setenv("ENV_INT_VAR_ERR", "?")
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "initialize FlagInt by default value",
			fields: fields{
				Name:    "a",
				Summary: "",
				Value:   2,
				Aliases: nil,
				EnvVar:  "",
			},
		},
		{
			name: "initialize FlagInt by env value",
			fields: fields{
				Name:    "b",
				Summary: "",
				Value:   1,
				Aliases: nil,
				EnvVar:  "ENV_INT_VAR_OK",
			},
		},
		{
			name: "initialize FlagInt with wrong env value",
			fields: fields{
				Name:    "b",
				Summary: "",
				Aliases: nil,
				EnvVar:  "ENV_INT_VAR_ERR",
			},
		},
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
	// env variables for test purposes
	os.Setenv("ENV_BOOL_VAR_OK", "true")
	os.Setenv("ENV_BOOL_VAR_ERR", "?")
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "initialize FlagBool by default value",
			fields: fields{
				Name:    "a",
				Summary: "",
				Value:   true,
				Aliases: nil,
				EnvVar:  "",
			},
		},
		{
			name: "initialize FlagBool by env value",
			fields: fields{
				Name:    "b",
				Summary: "",
				Value:   false,
				Aliases: nil,
				EnvVar:  "ENV_BOOL_VAR_OK",
			},
		},
		{
			name: "initialize FlagBool with wrong env value",
			fields: fields{
				Name:    "b",
				Summary: "",
				Aliases: nil,
				EnvVar:  "ENV_BOOL_VAR_ERR",
			},
		},
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
	// env variables for test purposes
	os.Setenv("ENV_STRING_VAR_OK", "str")
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "initialize FlagString by default value",
			fields: fields{
				Name:    "a",
				Summary: "",
				Value:   "str",
				Aliases: nil,
				EnvVar:  "",
			},
		},
		{
			name: "initialize FlagString by env value",
			fields: fields{
				Name:    "b",
				Summary: "",
				Value:   "",
				Aliases: nil,
				EnvVar:  "ENV_STRING_VAR_OK",
			},
		},
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
	// env variables for test purposes
	os.Setenv("ENV_STRING_SLICE_VAR_OK", "A,b,C,d,E")
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "initialize FlagString by default value",
			fields: fields{
				Name:    "a",
				Summary: "",
				Value:   nil,
				Aliases: nil,
				EnvVar:  "",
			},
		},
		{
			name: "initialize FlagString by env value",
			fields: fields{
				Name:    "b",
				Summary: "",
				Value:   nil,
				Aliases: nil,
				EnvVar:  "ENV_STRING_SLICE_VAR_OK",
			},
		},
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
