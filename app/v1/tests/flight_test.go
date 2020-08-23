package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"flight-api/app/core"
	v1 "flight-api/app/v1"
)

func initTest() *gin.Engine {
	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err.Error())
		os.Exit(1)
	}

	gin.SetMode(gin.ReleaseMode)

	core.InitDB()

	router := gin.Default()
	v1.Route(router)

	return router
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGetFlights(t *testing.T) {
	router := initTest()
	w := performRequest(router, "GET", "/v1/flight")

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	_, exists := response["flights"]

	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestGetFlightsWithQueryParam(t *testing.T) {
	router := initTest()
	w := performRequest(router, "GET", "/v1/flight?airline-code=QF")

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	_, exists := response["flights"]

	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestNegativeGetFlightsWithQueryParam(t *testing.T) {
	var w *httptest.ResponseRecorder
	router := initTest()

	w = performRequest(router, "GET", "/v1/flight?airline-code=Q")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	w = performRequest(router, "GET", "/v1/flight?airline-code=123")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	w = performRequest(router, "GET", "/v1/flight?airline-code=Q1")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	_, exists := response["error"]

	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestCloseDatabase(t *testing.T) {
	defer core.GetDB().Close()
}
