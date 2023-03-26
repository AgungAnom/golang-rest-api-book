package routers

import (
	"golang-rest-api-book/controllers"

	"github.com/gin-gonic/gin"
)


func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/books", controllers.GetAllBook)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.GetBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	
	return router
}