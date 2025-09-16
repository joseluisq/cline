package flag_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joseluisq/cline/flag"
)

func TestFlagInt_init(t *testing.T) {
	type fields struct {
		Name         string
		Summary      string
		Value        int
		Aliases      []string
		EnvVar       string
		FlagValue    flag.Value
		FlagAssigned bool
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
				Name:      "a",
				Value:     2,
				FlagValue: "2",
			},
		},
		{
			name: "initialize FlagInt by env value",
			fields: fields{
				Name:      "b",
				Value:     1,
				EnvVar:    "ENV_INT_VAR_OK",
				FlagValue: "1",
			},
		},
		{
			name: "initialize FlagInt with wrong env value",
			fields: fields{
				Name:      "b",
				Aliases:   []string{"x", "y"},
				EnvVar:    "ENV_INT_VAR_ERR",
				FlagValue: "0",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fi := &flag.FlagInt{
				Name:         tt.fields.Name,
				Summary:      tt.fields.Summary,
				Value:        tt.fields.Value,
				Aliases:      tt.fields.Aliases,
				EnvVar:       tt.fields.EnvVar,
				FlagValue:    tt.fields.FlagValue,
				FlagAssigned: tt.fields.FlagAssigned,
			}
			fi.Init()

			assert.Equal(t, tt.fields.Summary, fi.Summary)
			assert.Equal(t, tt.fields.Aliases, fi.Aliases)
			assert.Equal(t, tt.fields.EnvVar, fi.EnvVar)
			assert.Equal(t, tt.fields.Value, fi.Value)
			assert.Equal(t, tt.fields.FlagValue, fi.FlagValue)
		})
	}
}

func TestFlagBool_init(t *testing.T) {
	type fields struct {
		Name         string
		Summary      string
		Value        bool
		Aliases      []string
		EnvVar       string
		FlagValue    flag.Value
		FlagAssigned bool
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
				Name:      "a",
				Value:     true,
				FlagValue: "true",
			},
		},
		{
			name: "initialize FlagBool by env value",
			fields: fields{
				Name:      "b",
				EnvVar:    "ENV_BOOL_VAR_OK",
				FlagValue: "true",
			},
		},
		{
			name: "initialize FlagBool with wrong env value",
			fields: fields{
				Name:      "b",
				EnvVar:    "ENV_BOOL_VAR_ERR",
				FlagValue: "false",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb := &flag.FlagBool{
				Name:         tt.fields.Name,
				Summary:      tt.fields.Summary,
				Value:        tt.fields.Value,
				Aliases:      tt.fields.Aliases,
				EnvVar:       tt.fields.EnvVar,
				FlagValue:    tt.fields.FlagValue,
				FlagAssigned: tt.fields.FlagAssigned,
			}
			fb.Init()

			assert.Equal(t, tt.fields.Summary, fb.Summary)
			assert.Equal(t, tt.fields.Aliases, fb.Aliases)
			assert.Equal(t, tt.fields.EnvVar, fb.EnvVar)
			assert.Equal(t, tt.fields.Value, fb.Value)
			assert.Equal(t, tt.fields.FlagValue, fb.FlagValue)
		})
	}
}

func TestFlagString_init(t *testing.T) {
	type fields struct {
		Name         string
		Summary      string
		Value        string
		Aliases      []string
		EnvVar       string
		FlagValue    flag.Value
		FlagAssigned bool
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
				Name:      "a",
				Value:     "str",
				FlagValue: "str",
			},
		},
		{
			name: "initialize FlagString by env value",
			fields: fields{
				Name:      "b",
				EnvVar:    "ENV_STRING_VAR_OK",
				FlagValue: "str",
			},
		},
		{
			name: "initialize FlagString with wrong env value",
			fields: fields{
				Name:   "b",
				EnvVar: "ENV_STRING_VAR_ERR",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &flag.FlagString{
				Name:         tt.fields.Name,
				Summary:      tt.fields.Summary,
				Value:        tt.fields.Value,
				Aliases:      tt.fields.Aliases,
				EnvVar:       tt.fields.EnvVar,
				FlagValue:    tt.fields.FlagValue,
				FlagAssigned: tt.fields.FlagAssigned,
			}
			fs.Init()

			assert.Equal(t, tt.fields.Summary, fs.Summary)
			assert.Equal(t, tt.fields.Aliases, fs.Aliases)
			assert.Equal(t, tt.fields.EnvVar, fs.EnvVar)
			assert.Equal(t, tt.fields.Value, fs.Value)
			assert.Equal(t, tt.fields.FlagValue, fs.FlagValue)
		})
	}
}

func TestFlagStringSlice_init(t *testing.T) {
	type fields struct {
		Name         string
		Summary      string
		Value        []string
		Aliases      []string
		EnvVar       string
		FlagValue    flag.Value
		FlagAssigned bool
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
				Name: "a",
			},
		},
		{
			name: "initialize FlagString by env value",
			fields: fields{
				Name:      "b",
				EnvVar:    "ENV_STRING_SLICE_VAR_OK",
				FlagValue: flag.Value("A,b,C,d,E"),
			},
		},
		{
			name: "initialize FlagString with wrong env value",
			fields: fields{
				Name:   "b",
				EnvVar: "ENV_STRING_SLICE_VAR_ERR",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &flag.FlagStringSlice{
				Name:         tt.fields.Name,
				Summary:      tt.fields.Summary,
				Value:        tt.fields.Value,
				Aliases:      tt.fields.Aliases,
				EnvVar:       tt.fields.EnvVar,
				FlagValue:    tt.fields.FlagValue,
				FlagAssigned: tt.fields.FlagAssigned,
			}
			fs.Init()

			assert.Equal(t, tt.fields.Summary, fs.Summary)
			assert.Equal(t, tt.fields.Aliases, fs.Aliases)
			assert.Equal(t, tt.fields.EnvVar, fs.EnvVar)
			assert.Equal(t, tt.fields.Value, fs.Value)
			assert.Equal(t, tt.fields.FlagValue, fs.FlagValue)
		})
	}
}
