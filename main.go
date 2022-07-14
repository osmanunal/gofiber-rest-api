package main

import (
	"fmt"

	"github.com/osmanunal/gofiber-rest-api/book"
	"github.com/osmanunal/gofiber-rest-api/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	app := fiber.New()

	initDatabase()

	setupRoutes(app)
	app.Listen(":8080")

	defer database.DBConn.Close()
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("Failed connect database...")
	}
	fmt.Println("Connection Opened to Database...")

	database.DBConn.AutoMigrate(&book.Book{})

}

func setupRoutes(app *fiber.App) {

	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	// app.Post("/api/book", book.NewBook)
	app.Post("/api/book", book.PostBook) // Dinamic
	app.Delete("/api/book/:id", book.DeleteBook)
}
