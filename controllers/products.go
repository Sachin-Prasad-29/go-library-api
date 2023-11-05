package controllers

import "github.com/gofiber/fiber/v2"

func GetAllBooks(c *fiber.Ctx) error {

	return c.SendString("Get All Books")
}

func CreateBook(c *fiber.Ctx) error {

	return c.SendString("Add a Book")
}

func GetBook(c *fiber.Ctx) error {

	return c.SendString("Get Book By ID")
}

func EditBook(c *fiber.Ctx) error {

	return c.SendString("Edit Book By Id")
}

func DeleteBook(c *fiber.Ctx) error {

	return c.SendString("Delete A Booy By Id")
}
