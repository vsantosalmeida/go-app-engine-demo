package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	var tests = []struct {
		name, want, varName string
		function            func() string
	}{
		{"GetHashKey", "person", hashKey, GetHashKey},
		{"GetProjectId", "xpto", projectId, GetProjectId},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.varName, tt.want)
			got := os.Getenv(tt.varName)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConfigFail(t *testing.T) {
	var tests = []struct {
		name, varName string
		function      func() string
	}{
		{"HashKeyFail", hashKey, GetHashKey},
		{"ProjectIdFail", projectId, GetProjectId},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					assert.NotNil(t, r, "Should throw panic")
				}
			}()
			os.Setenv(tt.varName, "")
			tt.function()
		})
	}
}
