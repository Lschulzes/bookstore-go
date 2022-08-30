package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lschulzes/bookstore-go/pkg/models"
	"github.com/lschulzes/bookstore-go/pkg/utils"
	"net/http"
)

func CreateBook(c *gin.Context) {
	newBook := &models.Book{}
	parseErr := utils.ParseBody(c, newBook)
	if parseErr != nil {
		c.IndentedJSON(http.StatusBadRequest, parseErr)
		return
	}
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
	err := book.GetBook(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	idInt, _ := utils.GetParamsId(c)
	book := &models.Book{}
	err := book.DeleteBook(idInt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	idInt, _ := utils.GetParamsId(c)
	book, body := &models.Book{}, &models.Book{}
	parseErr := utils.ParseBody(c, body)
	if parseErr != nil {
		c.IndentedJSON(http.StatusBadRequest, parseErr)
		return
	}
	err := book.UpdateBook(idInt, body)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
