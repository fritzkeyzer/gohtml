package gohtml

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
			TemplateFile: "tests/conditional.gohtml",
			WantFile:     "tests/conditional.gohtml.go",
		}, {
			TemplateFile: "tests/loops.gohtml",
			WantFile:     "tests/loops.gohtml.go",
		}, {
			TemplateFile: "tests/nested.gohtml",
			WantFile:     "tests/nested.gohtml.go",
		}, {
			TemplateFile: "tests/person.gohtml",
			WantFile:     "tests/person.gohtml.go",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.TemplateFile, func(t *testing.T) {
			tt, err := ParseTemplate(tc.TemplateFile, "tests")
			if !assert.NoError(t, err) {
				return
			}

			buf := bytes.NewBuffer(nil)
			err = tt.Generate(buf)
			if !assert.NoError(t, err) {
				return
			}

			want, err := os.ReadFile(tc.WantFile)
			if !assert.NoError(t, err) {
				return
			}

			assert.Equal(t, string(want), buf.String())
		})
	}
}
