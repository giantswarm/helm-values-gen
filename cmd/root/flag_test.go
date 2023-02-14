package root

import (
	"testing"
)

func TestValidateFlags(t *testing.T) {
	testCases := []struct {
		name        string
		output      string
		force       bool
		expectError bool
	}{
		{
			name:        "no output, no force",
			output:      "",
			force:       false,
			expectError: false,
		},
		{
			name:        "output, no force",
			output:      "output",
			force:       false,
			expectError: false,
		},
		{
			name:        "no output, force",
			output:      "",
			force:       true,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			flag := &flag{
				output: tc.output,
				force:  tc.force,
			}

			err := flag.validate()

			if err != nil && !tc.expectError {
				t.Fatalf("Unexpected error in test case '%s': %s", tc.name, err)
			}
			if err == nil && tc.expectError {
				t.Fatalf("Expected error, got none in test case '%s'", tc.name)
			}
		})
	}
}
