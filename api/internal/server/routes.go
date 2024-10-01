package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	v1 := r.Group("/v1")
	sensors := v1.Group("/sensors")
	{
		sensors.GET("/", s.GetSensorsHandler)
		sensors.GET("/:sensor_id", s.GetSensorByIdHandler)
		sensors.POST("/", s.CreateSensorHandler)
		sensors.PUT("/:sensor_id", s.UpdateSensorHandler)
		sensors.DELETE("/:sensor_id", s.DeleteSensorHandler)
	}
	return r
}

func (s *Server) GetSensorsHandler(c *gin.Context) {
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")

	resp := make(map[string]string)
	resp["limit"] = limit
	resp["offset"] = offset
	resp["message"] = "Yeetus Deletus"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetSensorByIdHandler(c *gin.Context) {
	id := c.Query("sensor_id")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Sensor for id %v not found", id)})
	} else {
		resp := make(map[string]string)
		resp["message"] = fmt.Sprintf("Found sensor %s", id)

		c.JSON(http.StatusOK, resp)
	}
}

func (s *Server) CreateSensorHandler(c *gin.Context) {
	var sensor Sensor
	if c.ShouldBind(&sensor) == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("invalid body: %v", &sensor)})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("created sensor: %v", &sensor)})
	}
}

func (s *Server) UpdateSensorHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func (s *Server) DeleteSensorHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func (s *Server) YeetHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Yeetus Deletus"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}
