package controllers

import "github.com/gofiber/fiber/v2"

func SignUp(c *fiber.Ctx) error {
	return c.SendString("Sign Up User")
}

func UserSignIn(c *fiber.Ctx) error {
	return c.SendString("Sign In Normal User")
}

func AdminSignIn(c *fiber.Ctx) error {
	return c.SendString("Sign In Admin User")
}