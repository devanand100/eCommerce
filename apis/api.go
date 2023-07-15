package apis

import (
	"github.com/devanand100/eCommerce/apis/user"
	"github.com/gofiber/fiber/v2"
)

func Setup (app *fiber.App){
	userRoutes := app.Group("/api/user")
	user.Routes(userRoutes)
}