package proxy

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/proxy"
)

// reverse proxy
func To(baseURL string) fiber.Handler {
	return func(c fiber.Ctx) error {
		frontendURL := os.Getenv("FRONTEND_URL")
		target := baseURL + c.OriginalURL()
		log.Println(target)
		log.Println(c.OriginalURL())

		if err := proxy.Do(c, target); err != nil {
			return c.Status(502).JSON(fiber.Map{"error": "service unavailable"})
		}

		c.Set("Access-Control-Allow-Origin", frontendURL)
		c.Set("Access-Control-Allow-Credentials", "true")

		return nil
	}
}
