package http

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
)

func (s *Server) createRenderer() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	renderer = registerTemplate(renderer, "index")
	renderer = registerTemplate(renderer, "settings")
	renderer = registerTemplate(renderer, "commands")
	renderer = registerTemplate(renderer, "tags")
	renderer = registerTemplate(renderer, "panels")
	renderer = registerTemplate(renderer, "premium")

	return renderer
}

func registerTemplate(renderer multitemplate.Renderer, name string) multitemplate.Renderer {
	renderer.AddFromFiles(name,
		"./public/templates/layouts/main.tmpl",
		"./public/templates/includes/head.tmpl",
		"./public/templates/includes/navbar.tmpl",
		fmt.Sprintf("./public/templates/views/%s.tmpl", name),
		)

	return renderer
}

