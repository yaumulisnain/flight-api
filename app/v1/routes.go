package v1

import (
	"github.com/gin-gonic/gin"

	"flight-api/app/v1/delivery"
)

func Route(route *gin.Engine) *gin.RouterGroup {
	api := route.Group("/v1")
	{
		flights := api.Group("/flight")
		{
			flights.GET("", delivery.GetFlights)
			flights.GET("/:id", delivery.GetFlightById)
			flights.POST("", delivery.CreateFlight)
			flights.PUT("/:id", delivery.UpdateFlight)
			flights.DELETE("/:id", delivery.DeleteFlight)
		}
	}

	return api
}
