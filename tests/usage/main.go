package main

import (
	"os"

	"github.com/fritzkeyzer/gohtml/tests"
)

func main() {
	content := tests.Nested(tests.NestedData{
		Organisation: tests.NestedOrganisation{
			Name:    "company",
			Founded: "long ago",
		},
		Employee: tests.NestedEmployee{
			Personal: tests.NestedEmployeePersonal{
				Address: tests.NestedEmployeePersonalAddress{
					Street:  "something lane",
					City:    "42",
					Country: "lala land",
				},
			},
		},
	})

	err := tests.RenderBaseLayout(os.Stdout, tests.BaseLayoutData{
		Title: "Hello World",
		Imports: tests.ImportsData{
			Imports: []tests.ImportsImport{
				{Src: "https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js"},
				{Src: "https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js", Defer: true},
			},
		},
		//language=javascript
		BodyScript: `console.log('Hello World')`,
		Nav: tests.NavbarData{
			Title: "Hello World",
			Links: []tests.NavbarLink{
				{
					Name: "Home",
					Link: tests.NavLinkData{
						Href: "/home",
						Text: "Home",
					},
				},
				{
					Name: "Home",
					Link: tests.NavLinkData{
						Href: "/home",
						Text: "Home",
					},
				},
			},
		},
		Content: content,
	})
	if err != nil {
		panic(err)
	}
}
