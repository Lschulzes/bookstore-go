package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lschulzes/bookstore-go/pkg/models"
	"github.com/lschulzes/bookstore-go/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func CreateBook(c *gin.Context) {
	newBook := &models.Book{}
	utils.ParseBody(c, newBook)
	b := newBook.CreateBook()
	c.IndentedJSON(http.StatusOK, b)
}
func GetBooks(c *gin.Context) {
	books := models.GetAllBooks()
	c.IndentedJSON(http.StatusOK, books)
}
func GetBook(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book, _ := models.GetBook(idInt)
	c.IndentedJSON(http.StatusOK, book)
}
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(idInt)
	c.IndentedJSON(http.StatusOK, book)
}

//func UpdateBook(c *gin.Context) {
//	id := c.Param("id")
//	idInt, err := strconv.ParseInt(id, 10, 64)
//	if err != nil {
//		fmt.Println("Error while parsing")
//	}
//}
