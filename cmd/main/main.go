package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lschulzes/bookstore-go/pkg/routes"
)

func main() {
	r := gin.Default()
	routes.RegisterBookStoreRoutes(r)
	err := r.Run("localhost:3000")
	if err != nil {
		return
	}
	fmt.Println("Running Application on port 3000")
}
