package controller

import (
	"github.com/gin-gonic/gin"
)

type ItineraryController struct{}

func (i *ItineraryController) GenerateItinerary(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
