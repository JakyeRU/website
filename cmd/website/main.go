package main

import (
	"github.com/TicketsBot/website/config"
	"github.com/TicketsBot/website/http"
)

func main() {
	config.LoadConfig()

	server := http.NewServer()
	server.RegisterRoutes()
	server.Listen()
}
