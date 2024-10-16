package models

import (
	"golang/database"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Handler to get all todos
func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

// Handler to get a todo by ID
func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	// Attempt to find the todo by ID
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Todo not found", "data": err})
	}
	return c.JSON(&todo)
}

// Handler to create a new todo
func CreateTodo(c *fiber.Ctx) error {
	db := database.DBConn
	todo := new(Todo)

	// Try to parse the request body into the Todo struct
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Cannot parse JSON", "data": err})
	}

	// Save the todo to the database
	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Todo", "data": err})
	}

	// Return the new todo with the generated ID
	return c.JSON(todo)
}

// Handler to update an existing todo
func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id") // Get the todo ID from the URL
	db := database.DBConn
	var todo Todo

	// Try to find the todo with the given ID
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Todo not found", "data": err})
	}

	// Parse the request body to update the todo
	err = c.BodyParser(&todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Cannot parse JSON", "data": err})
	}

	// Save the changes to the database
	db.Save(&todo)
	return c.JSON(&todo)
}

// Handler to delete a todo
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var todo Todo

	// Try to find the todo with the given ID
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Todo not found", "data": err})
	}

	// Delete the found todo
	db.Delete(&todo)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Todo deleted"})
}
