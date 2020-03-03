package main

import (
	"errors"
	"path/filepath"
	"regexp"
	"testing"
)

var validatePathTests = map[string]struct {
	path             string
	expectErr        bool
	expectedErrCount int
}{
	"all_valid": {
		path: "allvalid",
	},
	"ignored": {
		path: "ignored",
	},
	"invalid": {
		path:             "invalid",
		expectErr:        true,
		expectedErrCount: 5,
	},
}

func TestValidatePath(t *testing.T) {
	checkRe := regexp.MustCompile(`golang\.org`)
	for testName, testCase := range validatePathTests {
		t.Run(testName, func(t *testing.T) {
			root := filepath.Join("testdata", testCase.path)
			err := validatePath(root, checkRe)
			if err != nil {
				if !testCase.expectErr || testCase.expectedErrCount == 0 {
					t.Errorf("unexpected error: %s", err)
				}
				if testCase.expectedErrCount != 0 {
					var errSet errorSet
					if errors.As(err, &errSet) {
						if len(errSet) != testCase.expectedErrCount {
							t.Errorf("unexpected error count (expected=%d, actual=%d)", testCase.expectedErrCount, len(errSet))
						}
					} else {
						t.Errorf("unexpected error type: %v [%T]", err, err)
					}
				}
				return
			}
			if testCase.expectErr {
				t.Error("unexpectedly no error")
			}
		})
	}
}
