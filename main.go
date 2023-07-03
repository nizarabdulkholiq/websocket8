package main

import (
	"log"

	"github.com/aiteung/musik"
	"github.com/nizarabdulkholiq/websocketd8/module"
	"github.com/nizarabdulkholiq/websocketd8/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go module.RunHub()

	site := fiber.New()
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
