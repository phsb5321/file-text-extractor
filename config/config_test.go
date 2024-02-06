package config

import (
	"os"
	"testing"
)

func TestParseCommandLineArguments(t *testing.T) {
	// Create input directory
	if err := os.Mkdir("input", 0777); err != nil {
		t.Fatalf("Failed to create input directory: %v", err)
	}
	defer os.RemoveAll("input")

	tests := []struct {
		args []string
		want *Config
	}{
		{
			args: []string{"-d", "input", "-o", "output.txt", "-i", ".jpg,.png"},
			want: &Config{
				InputDir:    "input",
				OutputFile:  "output.txt",
				IgnoredExts: []string{".jpg", ".png"},
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := ParseCommandLineArguments(tt.args)
			if err != nil {
				t.Errorf("Failed to parse command-line arguments: %v", err)
				return
			}
			if !compareConfigs(got, tt.want) {
				t.Errorf("Unexpected config, expected %+v, got %+v", tt.want, got)
			}
		})
	}
}

// compareConfigs compares two Config structs and returns true if they are equal, false otherwise.
func compareConfigs(c1, c2 *Config) bool {
	return c1.InputDir == c2.InputDir &&
		c1.OutputFile == c2.OutputFile &&
		compareStringSlices(c1.IgnoredExts, c2.IgnoredExts)
}

// compareStringSlices compares two string slices and returns true if they are equal, false otherwise.
func compareStringSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
