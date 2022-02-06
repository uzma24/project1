package handler

import (
	"log"
	"net/http"

	"github.com/project1/internal"
	"github.com/project1/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetFrequency(c *gin.Context) {
	resp, err := internal.GetFrequency(c)
	if err != nil {
		utils.HandleError(c, err)
		log.Printf("error occured while calling to GET GetFrequency: %v", err)
		return
	}
	utils.SuccessResponse(c, gin.H{
		"status":  http.StatusOK,
		"message": "frequency received successfully",
		"error":   false,
		"data":    resp,
	})
}
