package router

import "github.com/gin-gonic/gin"

func Initialize(port string) {
	router := gin.Default()
	initializeRoutes(router)
	
	router.Run(":" + port)
}
