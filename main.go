package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	ID    string `json: "id"`
	Name  string `json: "name"`
	Age   string `json: "age"`
	Email string `json: "email"`
}

var users = make(map[string]User)

func main() {
	app := fiber.New()

	app.Post("/user", createUser)
	app.Get("/user/:id", getUser)
	app.Delete("/user/:id", deleteUser)
	app.Put("/user", updateUser)
	app.Listen(":3000")
}

func createUser(c *fiber.Ctx) error {
	var newUser User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Can't parse JSON"})
	}

	if newUser.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name is required"})
	}

	newUser.ID = uuid.New().String()
	users[newUser.ID] = newUser

	return c.Status(201).JSON(fiber.Map{"id": newUser.ID})
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, exists := users[id]
	if !exists {
		return c.Status(400).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(200).JSON(user)
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, exists := users[id]; !exists {
		return c.Status(400).JSON(fiber.Map{"error": "User not found"})
	}

	delete(users, id)
	return c.SendStatus(200)
}

func updateUser(c *fiber.Ctx) error {
	var updatedUser User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if updatedUser.ID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "ID us required"})
	}

	if _, exists := users[updatedUser.ID]; !exists {
		return c.Status(400).JSON(fiber.Map{"error": "User not found"})
	}

	users[updatedUser.ID] = updatedUser
	return c.Status(200).JSON(fiber.Map{"id": updatedUser.ID})
}
