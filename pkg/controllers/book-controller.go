package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lschulzes/bookstore-go/pkg/models"
	"github.com/lschulzes/bookstore-go/pkg/utils"
	"net/http"
)

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
	idInt, _ := utils.GetParamsId(c)
	book := &models.Book{}
	book.GetBook(idInt)
	c.IndentedJSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	idInt, _ := utils.GetParamsId(c)
	book := &models.Book{}
	err := book.DeleteBook(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	idInt, _ := utils.GetParamsId(c)
	book := &models.Book{}
	utils.ParseBody(c, book)
	err := book.UpdateBook(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
