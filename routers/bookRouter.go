package routers

import (
	"golang-rest-api-book/controllers"

	"github.com/gin-gonic/gin"
)


func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/books", controllers.CreateBook)


	
	return router
}