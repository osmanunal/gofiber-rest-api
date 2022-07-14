package book

import (
	"github.com/osmanunal/gofiber-rest-api/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// Get - /api/book
func GetBooks(c *fiber.Ctx) error {

	db := database.DBConn
	var book []Book
	db.Find(&book)

	return c.JSON(book)
}

// Get - /api/book/id
func GetBook(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)

	return c.JSON(book)
}

// Post (Static) - /api/book
func NewBook(c *fiber.Ctx) error {

	db := database.DBConn

	var book Book
	book.Title = "Chess"
	book.Author = "Stefan Zweig"
	book.Rating = 4

	db.Create(&book)

	return c.JSON(book)
}

// Post (Dinamic) - /api/book

func PostBook(c *fiber.Ctx) error { // bu yapı başka bir projede vardı sabah bak
	db := database.DBConn
	book := new(Book)

	if err := c.BodyParser(book); err != nil {

		c.Status(503).SendString(err.Error())
	}
	db.Create(&book)

	return c.JSON(book)

}

// Delete - /api/book/id
func DeleteBook(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DBConn

	var book Book

	db.First(&book, id)
	if book.Title == "" {
		return c.Status(503).SendString("No Book Found With ID")
	}
	db.Delete(&book)

	return c.SendString("Book Seccesfully deleted")
}
