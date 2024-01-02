package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/simonha9/T2DGen/internal/app/service"
)

type ItineraryController struct {
	ItineraryService service.ItineraryService
}

func NewItineraryController() *ItineraryController {
	return &ItineraryController{
		ItineraryService: service.NewItineraryService(),
	}
}

func (i *ItineraryController) GenerateItinerary(c *gin.Context) {

	itin, err := i.ItineraryService.GenerateItinerary()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": itin,
	})
}
