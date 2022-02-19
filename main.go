package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	// rating
	ratingJson, _ := os.Open("rating.json")
	dataRating := json.NewDecoder(ratingJson)
	var rating Rating
	dataRating.Decode(&rating)
	// book details
	bookJson, _ := os.Open("book_detail.json")
	dataBook := json.NewDecoder(bookJson)
	var book Book
	dataBook.Decode(&book)
	// bestseller
	bestsellerJson, _ := os.Open("bestseller.json")
	dataBestseller := json.NewDecoder(bestsellerJson)
	var bestseller []BookList
	dataBestseller.Decode(&bestseller)

	router := gin.Default()

	router.GET("/books/:id", func(c *gin.Context) {
		c.JSON(200, book)
	})
	router.GET("/books/:id/bestseller", func(c *gin.Context) {
		c.JSON(200, bestseller)
	})
	router.GET("/books/:id/rating", func(c *gin.Context) {
		c.JSON(200, rating)
	})

	router.Run()
}
