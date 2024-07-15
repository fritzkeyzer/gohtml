package parse

import (
	"bytes"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTemplate_Generate(t *testing.T) {
	type tc struct {
		TemplateFile string
		WantFile     string
	}

	tcs := []tc{
		{
			TemplateFile: "tests/basic.gohtml",
			WantFile:     "tests/basic.gohtml.go",
		}, {
			TemplateFile: "tests/person.gohtml",
			WantFile:     "tests/person.gohtml.go",
		}, {
			TemplateFile: "tests/conditional.gohtml",
			WantFile:     "tests/conditional.gohtml.go",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.TemplateFile, func(t *testing.T) {
			tt := MustParseTemplate(tc.TemplateFile)

			buf := bytes.NewBuffer(nil)
			tt.Generate(buf)

			want, err := os.ReadFile(tc.WantFile)
			if err != nil {
				panic(err)
			}

			diff := cmp.Diff(string(want), buf.String())
			if diff != "" {
				t.Errorf("got != want, diff:\n%s", diff)
			}
		})
	}
}
