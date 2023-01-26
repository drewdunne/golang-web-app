package render

import (
	"github.com/drewdunne/go-course/pkg/helpers"
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err1 := template.ParseFiles(helpers.Pwd() + "/templates/" + tmpl)
	if err1 != nil {
		fmt.Println("issue getting template:", err1)
	}
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}