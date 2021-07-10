package render

import (
	"fmt"
	"github.com/TicketsBot/website/utils"
	"html/template"
)

func CreateTemplate(name string) *template.Template {
	template, err := template.New(name).ParseFiles(
		"./public/templates/layouts/main.tmpl",
		"./public/templates/includes/head.tmpl",
		"./public/templates/includes/navbar.tmpl",
		fmt.Sprintf("./public/templates/views/%s.tmpl", name),
	)

	utils.Must(err)

	return template
}
