package detectpath

import (
	"testing"
)

func TestDetectPath(t *testing.T) {
	testCases := []struct {
		Input  string
		Output string
	}{
		{
			"./feer/fwef.go",
			"./feer/fwef.go",
		},
		{
			"./feer-23/fwef-23.go",
			"./feer-23/fwef-23.go",
		},
		{
			"./feer-23/fwef-23.go:fjw",
			"./feer-23/fwef-23.go",
		},
		{
			"~/aaa/aaa.go",
			"~/aaa/aaa.go",
		},
		{
			"aaa.go",
			"aaa.go",
		},
	}
	for _, v := range testCases {
		if ret := DetectPath(v.Input); ret == nil || v.Output != ret.File {
			t.Fatalf("\nInput: %#v\n\nOutput: %#v", v.Output, ret.File)
		}
	}
}