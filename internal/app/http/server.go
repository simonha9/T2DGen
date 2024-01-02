package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/simonha9/T2DGen/internal/app/http/controller"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	itinController := controller.NewItineraryController()
	itinGroup := r.Group("/itinerary")
	{
		itinGroup.GET("/generate", itinController.GenerateItinerary)
	}

	r.Run() // listen and serve on
}
