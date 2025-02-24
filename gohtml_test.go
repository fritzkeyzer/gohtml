package gohtml

import (
	"os"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// TestDirectory runs gohtml on a directory and compares the buffered output to that of the golden file.
// This does not test file writing.
func TestDirectory(t *testing.T) {
	tests := []struct {
		gen        string
		goldenFile string
	}{
		{
			gen:        "tests/gohtml.gen.go",
			goldenFile: "tests/gohtml.gen.go.golden",
		},
		{
			gen:        "tests/nested/views/gohtml.gen.go",
			goldenFile: "tests/nested/views/gohtml.gen.go.golden",
		},
	}
	for _, tt := range tests {
		t.Run(tt.gen, func(t *testing.T) {
			// got gen
			gotBuf, err := os.ReadFile(tt.gen)
			if err != nil {
				t.Error(err)
				return
			}

			// load golden file
			wantBuf, err := os.ReadFile(tt.goldenFile)
			if err != nil {
				t.Error(err)
				return
			}

			if string(gotBuf) != string(wantBuf) {
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(string(gotBuf), string(wantBuf), false)
				t.Error(dmp.DiffPrettyText(diffs))
			}
		})
	}
}
