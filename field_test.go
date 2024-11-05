package gohtml

import (
	"html/template"
	"reflect"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func Test_extractTemplateFields(t *testing.T) {
	type tc struct {
		TemplateFile string
		TemplateName string
		Want         []Field
	}

	tests := []tc{
		{
			TemplateFile: "tests/nested.gohtml",
			TemplateName: "Nested",
			Want: []Field{
				{
					Path: []string{"Nested", "Organisation"},
					Name: "Name",
					Type: "any",
				}, {
					Path: []string{"Nested", "Organisation"},
					Name: "Founded",
					Type: "any",
				}, {
					Path: []string{"Nested", "Employee", "Personal", "Address"},
					Name: "Street",
					Type: "any",
				}, {
					Path: []string{"Nested", "Employee", "Personal", "Address"},
					Name: "City",
					Type: "any",
				}, {
					Path: []string{"Nested", "Employee", "Personal", "Address"},
					Name: "Country",
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
				}, {
					Path: []string{"Nested", "Organisation"},
					Name: "Founded",
					Type: "any",
				}, {
					Path: []string{"Nested", "Employee", "Personal", "Address"},
					Name: "Street",
					Type: "any",
				}, {
					Path: []string{"Nested", "Employee", "Personal", "Address"},
					Name: "City",
					Type: "any",
				}, {
					Path: []string{"Nested", "Employee", "Personal", "Address"},
					Name: "Country",
					Type: "any",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.TemplateName, func(t *testing.T) {
			tmpl, err := template.ParseFiles(tt.TemplateFile)
			if err != nil {
				t.Error(err)
				return
			}

			got := extractTemplateFields(tt.TemplateName, tmpl)

			td.Cmp(t, got, tt.Want)
			//if !assert.Equal(t, tt.Want, got) {
			//	t.Log("Got:")
			//	pp.Println(got)
			//}
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
				pipe: "$P := .People",
			},
			want: Field{
				Path: []string{"Loop"},
				Name: "People",
				Type: "[]LoopPeopleItem",
			},
		}, {
			args: args{
				path: []string{"Loop"},
				pipe: "$item := .Options",
			},
			want: Field{
				Path: []string{"Loop"},
				Name: "Options",
				Type: "[]LoopOption",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.pipe, func(t *testing.T) {
			got := parseRangeNodeField(tt.args.path, tt.args.pipe)
			td.Cmp(t, got, tt.want)
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
			//assert.Equal(t, tt.want, got)
			if !reflect.DeepEqual(got, tt.want) {
				td.Cmp(t, got, tt.want)
			}
		})
	}
}

func Test_parseIfNodeField(t *testing.T) {
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
				pipe: "not .SignedIn",
			},
			want: Field{
				Path: []string{"Person"},
				Name: "SignedIn",
				Type: "any",
			},
		},
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
			got := parseIfNodeField(tt.args.path, tt.args.pipe)
			//assert.Equal(t, tt.want, got)
			if !reflect.DeepEqual(got, tt.want) {
				td.Cmp(t, got, tt.want)
			}
		})
	}
}
