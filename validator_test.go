package main

import (
	"errors"
	"fmt"
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
	ignoreRe := regexp.MustCompile(`^localhost|example`)
	reTypes := []string{"check", "ignore", "both"}
	for testName, testCase := range validatePathTests {
		t.Run(testName, func(t *testing.T) {
			for _, reType := range reTypes {
				t.Run(fmt.Sprintf("re_type=%s", reType), func(t *testing.T) {
					root := filepath.Join("testdata", testCase.path)
					var v validator
					switch reType {
					case "check":
						v = validator{
							checkRe: checkRe,
						}
					case "ignore":
						v = validator{
							ignoreRe: ignoreRe,
						}
					case "both":
						v = validator{
							checkRe:  checkRe,
							ignoreRe: ignoreRe,
						}
					}
					err := v.validatePath(root)
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
		})
	}
}
