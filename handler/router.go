package handler

import (
	"github.com/gin-gonic/gin"
)

//using gin framework to set routes and call GetFrequency handler
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/text", GetFrequency)
	return router
}
