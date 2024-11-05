package gohtml

import (
	"bytes"
	"os"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

// TestDirectory runs gohtml on a directory and compares the buffered output to that of the golden file.
// This does not test file writing.
func TestDirectory(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		dir        string
		goldenFile string
	}{
		{
			dir:        "tests",
			goldenFile: "tests/gohtml.gen.go.golden",
		},
	}
	for _, tt := range tests {
		t.Run(tt.dir, func(t *testing.T) {
			g, err := ParseDir(tt.dir)
			if err != nil {
				t.Error(err)
				return
			}

			gotBuf := new(bytes.Buffer)
			err = g.Generate(gotBuf)
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

			td.CmpString(t, gotBuf.String(), string(wantBuf))
		})
	}
}
