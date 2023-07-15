package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	// "github.com/database"
)

func signUp(c *fiber.Ctx) error {
	body := c.Body()

	fmt.Printf("body:-->%s\n",body)

	return nil
}
