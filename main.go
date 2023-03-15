package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Chronicles of Narnia", Author: "prince caspian", Quantity: 2},
	{ID: "2", Title: "Spare", Author: "prince Harry", Quantity: 3},
	{ID: "3", Title: "Golden bells", Author: " unknown", Quantity: 1},
}

func getBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}
func createBook(c echo.Context) error {
	var newBook book
	if err := c.Bind(&newBook); err != nil {
		log.Fatalln("error creating book")
	}
	books = append(books, newBook)
	return c.JSON(http.StatusOK, newBook)
}
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
func bookById(c echo.Context) error {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		err := c.JSON(http.StatusNotFound, "Book not available")
		return err
	}
	c.JSON(http.StatusOK, book)
	return err
}
func checkOutBook(c echo.Context) error {
	id := c.QueryParam("id")
	book, err := getBookById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "missing query parameter")
		return err
	}
	if book.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, "book not available")
		return nil
	}
	book.Quantity -= 1
	c.JSON(http.StatusOK, book)
	return nil
}
func returnBook(c echo.Context) error {
	id := c.QueryParam("id")
	book, err := getBookById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "missing query parameter")
		return err
	}
	book.Quantity += 1
	c.JSON(http.StatusOK, book)
	return nil
}
func main() {
	e := echo.New()
	e.GET("/books", getBooks)
	e.POST("/books", createBook)
	e.GET("/books/:id", bookById)
	e.PATCH("/checkout", checkOutBook)
	e.PATCH("/return", returnBook)
	err := e.Start(":1323")
	if err != nil {
		return
	}
}
