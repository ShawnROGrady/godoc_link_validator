package main

import (
	"path/filepath"
	"testing"
)

var validatePathTests = map[string]struct {
	path      string
	expectErr bool
}{
	"all_valid": {
		path: "allvalid",
	},
}

func TestValidatePath(t *testing.T) {
	for testName, testCase := range validatePathTests {
		t.Run(testName, func(t *testing.T) {
			root := filepath.Join("testdata", testCase.path)
			err := validatePath(root)
			if err != nil {
				if !testCase.expectErr {
					t.Errorf("unexpected error: %s", err)
				}
				return
			}
			if testCase.expectErr {
				t.Error("unexpectedly no error")
			}
		})
	}
}
