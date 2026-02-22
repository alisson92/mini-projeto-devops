package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// book representa a estrutura de dados dos meus livros pessoais.
type book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Authors string `json:"authors"`
	Details string `json:"details"`
}

// books slice para armazenar os dados dos meus livros pessoais.
var books = []book{
	{ID: "1", Title: "Infraestrutura como Código", Authors: "Kief Morris", Details: "3ª Edição, O'Reilly/Novatec"},
	{ID: "2", Title: "Manual de DevOps", Authors: "Gene Kim, Jez Humble, Patrick Debois, John Willis", Details: "Alta Books"},
	{ID: "3", Title: "O Projeto Fênix", Authors: "Gene Kim, Kevin Behr, George Spafford", Details: "Edição Comemorativa, Alta Books"},
	{ID: "4", Title: "Entrega Contínua", Authors: "Jez Humble, David Farley", Details: "Bookman"},
	{ID: "5", Title: "Computação em Nuvem", Authors: "Thomas Erl, Ricardo Puttini, Zaigham Mahmood", Details: "2ª Edição, Bookman"},
}

func main() {
	router := gin.Default()

	// Route of healthcheck
	router.GET("/ping", getHealthcheck)

	// Route of books
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")
}

func getHealthcheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"ping": "pong", "status": "online"})
}

// getBooks pra retornar as informações dos livros em formato JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// postBooks adds a book from JSON received in the request body.
func postBooks(c *gin.Context) {
	var newBook book

	// Call BindJSON to bind the received JSON to
	// newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Adiciona um novo livro ao slice de livros.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// getBookByID locates the book whose ID value matches the id
// parameter sent by the client, then returns that book as a response.
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of books, looking for
	// a book whose ID value matches the parameter.
	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
