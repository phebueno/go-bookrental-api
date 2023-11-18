package main

import (
	"errors"
	"example/go-bookrental-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Importa as funções auxiliares do pacote principal

type book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: 1, Title: "Do Androids Dream of Electric Sheep?", Author: "Philip K. Dick", Quantity: 5},
	{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 4},
	{ID: 3, Title: "The Witcher", Author: "Andrzej Sapkowski", Quantity: 3},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBook(c *gin.Context) {
	idStr := c.Param("id")
	id, errorMsg, errorStatus := helper.ConvertIDParam(idStr)

	if errorMsg != nil {
		c.IndentedJSON(errorStatus, gin.H{"message": errorMsg})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id int) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func rentBook(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	id, errorMsg, errorStatus := helper.ConvertQueryParam(idStr, ok)
	if errorMsg != nil {
		c.IndentedJSON(errorStatus, gin.H{"message": errorMsg})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found."})
		return
	}

	if book.Quantity == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}

	book.Quantity--
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	id, errorMsg, errorStatus := helper.ConvertQueryParam(idStr, ok)
	if errorMsg != nil {
		c.IndentedJSON(errorStatus, gin.H{"message": errorMsg})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", createBook)
	router.PATCH("/checkout", rentBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8000")
}
