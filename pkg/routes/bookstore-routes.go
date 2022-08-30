package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lschulzes/bookstore-go/pkg/controllers"
)

var RegisterBookStoreRoutes = func(r *gin.Engine) {
	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
}
