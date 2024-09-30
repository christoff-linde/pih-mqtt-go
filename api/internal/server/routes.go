package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	v1 := r.Group("/v1/sensor")
	{
		v1.GET("/", s.HelloWorldHandler)
		v1.GET("/get_sensors", s.GetSensorsHandler)
		v1.GET("/get_sensor_by_id", s.GetSensorByIdHandler)
		v1.POST("/create_sensor", s.CreateSensorHandler)
		v1.PUT("/update_sensor", s.UpdateSensorHandler)
		v1.DELETE("/delete_sensor", s.DeleteSensorHandler)
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
