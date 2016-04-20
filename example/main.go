package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/boxofrad/uniform"
)

const formTemplate = `
<h1>Sign Up</h1>

<div>
	{{.first_name}}
</div>

<div>
	{{.last_name}}
</div>

<div>
	{{.occupation}}
</div>
`

func main() {
	http.HandleFunc("/", serveHTTP)
	http.ListenAndServe(":4321", nil)
}

type field interface {
	HTML() (template.HTML, error)
}

type fieldMap struct {
	markup map[string]template.HTML
	err    error
}

func (fm *fieldMap) add(name string, field field) {
	if fm.err == nil {
		fm.markup[name], fm.err = field.HTML()
	}
}

func buildFields() fieldMap {
	fields := fieldMap{markup: make(map[string]template.HTML)}

	fields.add("first_name", uniform.TextInput{
		Name:        "first_name",
		Placeholder: "John",
	})

	fields.add("last_name", uniform.TextInput{
		Name:        "last_name",
		Placeholder: "Doe",
	})

	fields.add("occupation", uniform.Select{
		Name: "occupation",
		Options: []uniform.SelectChild{
			uniform.OptGroup{
				Label: "Technology",
				Options: []uniform.Option{
					{
						Value: "developer",
						Text:  "Developer",
					},
					{
						Value:    "product_manager",
						Text:     "Product Manager",
						Selected: true,
					},
				},
			},
		},
	})

	return fields
}

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	handleErr := func(err error) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "ERROR")
		log.Println(err)
	}

	fields := buildFields()
	if fields.err != nil {
		handleErr(fields.err)
		return
	}

	tmpl, err := template.New("page").Parse(formTemplate)

	if err != nil {
		handleErr(err)
		return
	}

	if err := tmpl.Execute(w, fields.markup); err != nil {
		handleErr(err)
	}
}
