package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		ticker := c.Query("ticker")
		results := SearchTicker(ticker)

		return c.Render("results", fiber.Map{
			"Results": results,
		})
	})

	app.Get("/values/:ticker", func(c *fiber.Ctx) error {
		ticker := c.Params("ticker")
		values := GetDailyValues(ticker)

		return c.Render("values", fiber.Map{
			"Ticker": ticker,
			"Values": values,
		})
	})

	app.Listen(":3000")
}
