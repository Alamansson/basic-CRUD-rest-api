package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/createUser", createUser)
	router.GET("/readUsers", readUsers)
	router.PUT("/updateUser", updateUser)
	router.DELETE("/deleteUser", deleteUser)

	return router
}
