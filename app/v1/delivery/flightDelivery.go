package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"flight-api/app/v1/models"
	"flight-api/app/v1/usecase"
)

func GetFlightById(c *gin.Context) {
	var (
		user models.Flight
		err  error
		id   int64
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

	user, err = usecase.GetFlightByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"code":    http.StatusNotFound,
			"message": "ID Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}

func GetFlights(c *gin.Context) {
	var (
		Flights []models.Flight
		err     error
	)

	//TODO: Use search query param

	Flights, err = usecase.GetFlight()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"code":    http.StatusNotFound,
			"message": "Data Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    Flights,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}

func CreateFlight(c *gin.Context) {
	var (
		Flight models.Flight
		err    error
	)

	v := validator.New()
	err = c.ShouldBindJSON(&Flight)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Failed when parsing request",
		})
		return
	}

	if err = v.Struct(Flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Validation Error",
		})
		return
	}

	Flight, err = usecase.CreateFlight(Flight)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Failed to create data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    Flight,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}

func UpdateFlight(c *gin.Context) {
	var (
		updateData map[string]interface{}
		Flight     models.Flight
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

	Flight, err = usecase.UpdateFlight(int32(id), updateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"code":    http.StatusBadRequest,
			"message": "Failed to update data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    Flight,
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
