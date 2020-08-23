package delivery

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"flight-api/app/helper"
	"flight-api/app/v1/models"
	"flight-api/app/v1/usecase"
)

func GetFlightById(c *gin.Context) {
	var (
		flight models.Flight
		err    error
		id     int64
	)
	id, err = strconv.ParseInt(c.Params.ByName("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Failed to parse id",
		})
		return
	}

	flight, err = usecase.GetFlightByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"code":    http.StatusNotFound,
			"message": "ID Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    flight,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}

func GetFlights(c *gin.Context) {
	var (
		flights []models.Flight
		err     error
	)

	flightCodeParam := c.DefaultQuery("airline-code", "")

	if !helper.ValidationAirlineCode(flightCodeParam) && flightCodeParam != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Airline Code is not valid",
			"code":    http.StatusBadRequest,
			"message": "Airline Code is required 2-3 alphabetical character",
		})
		return
	}

	flightCodeParam = strings.ToUpper(flightCodeParam)

	flights, err = usecase.GetFlight(flightCodeParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"code":    http.StatusNotFound,
			"message": "Data Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"flights": flights,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}

func CreateFlight(c *gin.Context) {
	var (
		flight models.Flight
		err    error
	)

	v := validator.New()
	err = c.ShouldBindJSON(&flight)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Failed when parsing request",
		})
		return
	}

	if err = v.Struct(flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Validation Error",
		})
		return
	}

	flight, err = usecase.CreateFlight(flight)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Failed to create data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    flight,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}

func UpdateFlight(c *gin.Context) {
	var (
		updateData map[string]interface{}
		flight     models.Flight
		err        error
		id         int64
	)
	id, err = strconv.ParseInt(c.Params.ByName("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Failed to parse id",
		})
		return
	}

	err = c.ShouldBind(&updateData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Failed when parsing request",
		})
		return
	}

	flight, err = usecase.UpdateFlight(int32(id), updateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Failed to update data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    flight,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}

func DeleteFlight(c *gin.Context) {
	var (
		err error
		id  int64
	)
	id, err = strconv.ParseInt(c.Params.ByName("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Failed to parse id",
		})
		return
	}

	err = usecase.DeleteFlight(int32(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Failed to delete data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    "success",
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}
