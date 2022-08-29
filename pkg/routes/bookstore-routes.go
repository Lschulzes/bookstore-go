package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lschulzes/bookstore-go/pkg/controllers"
)

var RegisterBookStoreRoutes = func(r *gin.Engine) {
	r.POST("/book", controllers.CreateBook)
	r.GET("/book", controllers.GetBooks)
	r.GET("/book/:id", controllers.GetBook)
	r.DELETE("/book/:id", controllers.DeleteBook)
	r.PUT("/book/:id", controllers.UpdateBook)
}
