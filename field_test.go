package gohtml

import (
	"html/template"
	"reflect"
	"testing"

	"github.com/maxatome/go-testdeep/td"
	"github.com/stretchr/testify/assert"
)

func Test_extractTemplateFields(t *testing.T) {
	type tc struct {
		TemplateFile string
		TemplateName string
		Want         []Field
	}

	tests := []tc{
		{
			TemplateFile: "tests/basic.gohtml",
			TemplateName: "Basic",
			Want: []Field{
				{
					Path: []string{"Basic"},
					Name: "Name",
					Type: "any",
				},
				{
					Path: []string{"Basic"},
					Name: "Date",
					Type: "any",
				},
			},
		},
		{
			TemplateFile: "tests/conditional.gohtml",
			TemplateName: "Conditional",
			Want: []Field{
				{
					Path: []string{"Conditional"},
					Name: "SignedIn",
					Type: "any",
				},
			},
		},
		{
			TemplateFile: "tests/loops.gohtml",
			TemplateName: "Loops",
			Want: []Field{
				{
					Path: []string{"Loops"},
					Name: "Widgets",
					Type: "[]LoopsWidget",
				},
				{
					Path: []string{"LoopsWidget"},
					Name: "Name",
					Type: "any",
				},
				{
					Path: []string{"LoopsWidget"},
					Name: "Price",
					Type: "any",
				},
				{
					Path: []string{"Loops"},
					Name: "Socials",
					Type: "[]LoopsSocialsLink",
				},
				{
					Path: []string{"LoopsSocialsLink"},
					Name: "Name",
					Type: "any",
				},
				{
					Path: []string{"LoopsSocialsLink"},
					Name: "Href",
					Type: "any",
				},
			},
		},
		{
			TemplateFile: "tests/nested.gohtml",
			TemplateName: "Nested",
			Want: []Field{
				{
					Path: []string{"Nested", "Organisation"},
					Name: "Name",
					Type: "any",
				},
				{
					Path: []string{"Nested", "Organisation"},
					Name: "Founded",
					Type: "any",
				},
				{
					Path: []string{"Nested", "Employee", "Personal", "Address"},
					Name: "Street",
					Type: "any",
				},
				{
					Path: []string{"Nested", "Employee", "Personal", "Address"},
					Name: "City",
					Type: "any",
				},
				{
					Path: []string{"Nested", "Employee", "Personal", "Address"},
					Name: "Country",
					Type: "any",
				},
			},
		},
		{
			TemplateFile: "tests/person.gohtml",
			TemplateName: "Person",
			Want: []Field{
				{
					Path: []string{"Person"},
					Name: "Name",
					Type: "any",
				},
				{
					Path: []string{"Person"},
					Name: "Age",
					Type: "any",
				},
				{
					Path: []string{"Person", "Contact"},
					Name: "Phone",
					Type: "any",
				},
				{
					Path: []string{"Person", "Contact"},
					Name: "Email",
					Type: "any",
				},
				{
					Path: []string{"Person"},
					Name: "Socials",
					Type: "[]PersonSocialsLink",
				},
				{
					Path: []string{"PersonSocialsLink"},
					Name: "Name",
					Type: "any",
				},
				{
					Path: []string{"PersonSocialsLink"},
					Name: "Href",
					Type: "any",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.TemplateName, func(t *testing.T) {
			tmpl, err := template.ParseFiles(tt.TemplateFile)
			if !assert.NoError(t, err) {
				return
			}

			td.Cmp(t, extractTemplateFields(tt.TemplateName, tmpl), tt.Want)
		})
	}
}

func Test_parseRangeNodeField(t *testing.T) {
	type args struct {
		path []string
		pipe string
	}
	tests := []struct {
		args args
		want Field
	}{
		{
			args: args{
				path: []string{"Loop"},
				pipe: ".List",
			},
			want: Field{
				Path: []string{"Loop"},
				Name: "List",
				Type: "[]LoopListItem",
			},
		}, {
			args: args{
				path: []string{"Loop"},
				pipe: ".Options",
			},
			want: Field{
				Path: []string{"Loop"},
				Name: "Options",
				Type: "[]LoopOption",
			},
		}, {
			args: args{
				path: []string{"Loop"},
				pipe: "Item := .Options",
			},
			want: Field{
				Path: []string{"Loop"},
				Name: "Options",
				Type: "[]LoopOptionsItem",
			},
		}, {
			args: args{
				path: []string{"Loop"},
				pipe: "$item := .Options",
			},
			want: Field{
				Path: []string{"Loop"},
				Name: "Options",
				Type: "[]LoopOptionsItem",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.pipe, func(t *testing.T) {
			got := parseRangeNodeField(tt.args.path, tt.args.pipe)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRangeNodeField()\ngot: %+#v\nwant: %+#v", got, tt.want)
			}
		})
	}
}

func Test_parseActionNodeField(t *testing.T) {
	type args struct {
		path []string
		pipe string
	}
	tests := []struct {
		args args
		want Field
	}{
		{
			args: args{
				path: []string{"Person"},
				pipe: ".SignedIn",
			},
			want: Field{
				Path: []string{"Person"},
				Name: "SignedIn",
				Type: "any",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.pipe, func(t *testing.T) {
			got := parseActionNodeField(tt.args.path, tt.args.pipe)
			assert.Equal(t, tt.want, got)
		})
	}
}
