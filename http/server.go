package http

import (
	"github.com/TicketsBot/website/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

type Server struct {
	router *gin.Engine
}

func NewServer() Server {
	return Server{
		router: gin.Default(),
	}
}

func (s *Server) RegisterRoutes() {
	// middleware
	s.router.Use(gin.Recovery())

	s.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://panel.ticketsbot.net"},
		AllowMethods: []string{"GET"},
	}))

	// static files on /
	s.router.StaticFile("/manifest.webmanifest", "./public/static/manifest.webmanifest")
	s.router.StaticFile("/robots.txt", "./public/static/robots.txt")
	s.router.StaticFile("/sitemap.xml", "./public/static/sitemap.xml")

	// /assets
	s.router.Static("/assets", "./public/static/assets")

	// renderer
	s.router.HTMLRender = s.createRenderer()

	// routes
	s.router.GET("/", serveHtml("index", nil))
	s.router.GET("/settings", serveHtml("settings", nil))
	s.router.GET("/commands", serveHtml("commands", gin.H{
		"commands": config.Conf.Commands,
	}))
	s.router.GET("/tags", serveHtml("tags", nil))
	s.router.GET("/panels", serveHtml("panels", nil))
	s.router.GET("/premium", serveHtml("premium", nil))
	s.router.GET("/claiming", serveHtml("claiming", nil))
	s.router.GET("/privacy", serveHtml("privacy", nil))

	// redirect canned responses to new page
	s.router.GET("/cannedresponses", func(ctx *gin.Context) {
		ctx.Redirect(301, "/tags")
	})
}

func (s *Server) Listen() {
	if err := s.router.Run(os.Getenv("SERVER_ADDR")); err != nil {
		panic(err)
	}
}

func serveHtml(templateName string, variables map[string]interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(200, templateName, variables)
	}
}
