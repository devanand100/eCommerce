package user

import (
	"github.com/gofiber/fiber/v2"	
)

func Routes(route fiber.Router){
	route.Post("/signup",signUp)
}