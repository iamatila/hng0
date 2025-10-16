package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CatFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/me", func(c *fiber.Ctx) error {
		// fetch a random cat fact from the API
		fact := ""
		resp, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			log.Println("error fetching cat fact:", err)
		} else {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println("error reading response body:", err)
			} else {
				var cf CatFactResponse
				if err := json.Unmarshal(body, &cf); err != nil {
					log.Println("error unmarshalling json:", err)
				} else {
					fact = cf.Fact
				}
			}
		}

		return c.JSON(fiber.Map{
			"status": "success",
			"user": fiber.Map{
				"email": "tobiatilola@gmail.com",
				"name":  "Oluwatobiloba Atilola",
				"stack": "Go/Fiber, Node/Express",
			},
			"timestamp": time.Now().UTC(),
			"fact":      fact,
		})
	})

	log.Fatal(app.Listen(":8000"))
}
