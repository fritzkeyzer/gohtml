package gohtml

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestParseTemplate(t *testing.T) {
	type tc struct {
		TemplateFile string
		PackageName  string
		Want         GoHTML
	}

	tcs := []tc{
		{
			TemplateFile: "tests/basic.gohtml",
			PackageName:  "tests",
			Want: GoHTML{
				Name:        "Basic",
				FilePath:    "tests/basic.gohtml",
				PackageName: "tests",
				Templates: []Template{
					{
						Name:           "Basic",
						EmbedFilePath:  "basic.gohtml",
						TemplateString: "<h1>Hello {{.Name}}</h1>\n<footer>Today's date is {{.Date}}</footer>\n",
						Structs: []StructDef{
							{
								Name: "BasicData",
								Fields: []Field{
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
						},
					},
				},
			},
		},
		{
			TemplateFile: "tests/conditional.gohtml",
			PackageName:  "tests",
			Want: GoHTML{
				Name:        "Conditional",
				FilePath:    "tests/conditional.gohtml",
				PackageName: "tests",
				Templates: []Template{
					{
						Name:           "Conditional",
						EmbedFilePath:  "conditional.gohtml",
						TemplateString: "<button>\n    {{if .SignedIn}}\n        Sign Out\n    {{end}}\n</button>\n",
						Structs: []StructDef{
							{
								Name: "ConditionalData",
								Fields: []Field{
									{
										Path: []string{"Conditional"},
										Name: "SignedIn",
										Type: "any",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			TemplateFile: "tests/nested.gohtml",
			PackageName:  "tests",
			Want: GoHTML{
				Name:        "Nested",
				FilePath:    "tests/nested.gohtml",
				PackageName: "tests",
				Templates: []Template{
					{
						Name:           "Nested",
						EmbedFilePath:  "nested.gohtml",
						TemplateString: "<h1>{{.Organisation.Name}}</h1>\n<p>{{.Organisation.Founded}}</p>\n<div>\n    <p>{{.Employee.Personal.Address.Street}}</p>\n    <p>{{.Employee.Personal.Address.City}}</p>\n    <p>{{.Employee.Personal.Address.Country}}</p>\n</div>\n",
						Structs: []StructDef{
							{
								Name: "NestedData",
								Fields: []Field{
									{
										Path: []string{"Nested"},
										Name: "Organisation",
										Type: "NestedOrganisation",
									}, {
										Path: []string{"Nested"},
										Name: "Employee",
										Type: "NestedEmployee",
									},
								},
							}, {
								Name: "NestedEmployee",
								Fields: []Field{
									{
										Path: []string{"Nested", "Employee"},
										Name: "Personal",
										Type: "NestedEmployeePersonal",
									},
								},
							}, {
								Name: "NestedEmployeePersonal",
								Fields: []Field{
									{
										Path: []string{"Nested", "Employee", "Personal"},
										Name: "Address",
										Type: "NestedEmployeePersonalAddress",
									},
								},
							}, {
								Name: "NestedEmployeePersonalAddress",
								Fields: []Field{
									{
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
							}, {
								Name: "NestedOrganisation",
								Fields: []Field{
									{
										Path: []string{"Nested", "Organisation"},
										Name: "Name",
										Type: "any",
									}, {
										Path: []string{"Nested", "Organisation"},
										Name: "Founded",
										Type: "any",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			TemplateFile: "tests/person.gohtml",
			PackageName:  "tests",
			Want: GoHTML{
				Name:        "Person",
				FilePath:    "tests/person.gohtml",
				PackageName: "tests",
				Templates: []Template{
					{
						Name:           "Person",
						EmbedFilePath:  "person.gohtml",
						TemplateString: "<p>Name: {{.Name}}</p>\n<p>Age: {{.Age}}</p>\n<p>Phone: {{.Contact.Phone}}</p>\n<p>Email: {{.Contact.Email}}</p>\n<ul>\n    {{range $link := .Socials}}\n        <li>{{$link.Name}} {{$link.Href}}</li>\n    {{end}}\n</ul>\n",
						Structs: []StructDef{
							{
								Name: "PersonData",
								Fields: []Field{
									{
										Path: []string{"Person"},
										Name: "Name",
										Type: "any",
									}, {
										Path: []string{"Person"},
										Name: "Age",
										Type: "any",
									}, {
										Path: []string{"Person"},
										Name: "Contact",
										Type: "PersonContact",
									}, {
										Path: []string{"Person"},
										Name: "Socials",
										Type: "[]PersonSocialsLink",
									},
								},
							}, {
								Name: "PersonContact",
								Fields: []Field{
									{
										Path: []string{"Person", "Contact"},
										Name: "Phone",
										Type: "any",
									}, {
										Path: []string{"Person", "Contact"},
										Name: "Email",
										Type: "any",
									},
								},
							}, {
								Name: "PersonSocialsLink",
								Fields: []Field{
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
						},
					},
				},
			},
		},
		{
			TemplateFile: "tests/loops.gohtml",
			PackageName:  "tests",
			Want: GoHTML{
				Name:        "Loops",
				FilePath:    "tests/loops.gohtml",
				PackageName: "tests",
				Templates: []Template{
					{
						Name:           "Loops",
						EmbedFilePath:  "loops.gohtml",
						TemplateString: "{{range .Widgets}}\n    {{.Name}} - {{.Price}}\n{{end}}\n\n{{range $link := .Socials}}\n    {{$link.Name}} {{$link.Href}}\n{{end}}",
						Structs: []StructDef{
							{
								Name: "LoopsData",
								Fields: []Field{
									{
										Path: []string{"Loops"},
										Name: "Widgets",
										Type: "[]LoopsWidget",
									}, {
										Path: []string{"Loops"},
										Name: "Socials",
										Type: "[]LoopsSocialsLink",
									},
								},
							}, {
								Name: "LoopsSocialsLink",
								Fields: []Field{
									{
										Path: []string{"LoopsSocialsLink"},
										Name: "Name",
										Type: "any",
									}, {
										Path: []string{"LoopsSocialsLink"},
										Name: "Href",
										Type: "any",
									},
								},
							}, {
								Name: "LoopsWidget",
								Fields: []Field{
									{
										Path: []string{"LoopsWidget"},
										Name: "Name",
										Type: "any",
									}, {
										Path: []string{"LoopsWidget"},
										Name: "Price",
										Type: "any",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.TemplateFile, func(t *testing.T) {
			got, err := ParseTemplate(tt.TemplateFile, tt.PackageName)
			if err != nil {
				t.Error("ParseTemplate returned unexpected error:", err)
				return
			}
			if got == nil {
				t.Error("ParseTemplate returned nil")
				return
			}
			td.Cmp(t, *got, tt.Want)
		})
	}
}
