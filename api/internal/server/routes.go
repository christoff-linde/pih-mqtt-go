package server

import (
	"fmt"
	"net/http"
	"strconv"

	db "github.com/christoff-linde/pih-core-go/api/database"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	v1 := r.Group("/v1")
	sensors := v1.Group("/sensors")
	{
		sensors.GET("/", s.GetSensorsHandler)
		sensors.GET("/:id", s.GetSensorByIdHandler)
		sensors.POST("/", s.CreateSensorHandler)
		sensors.PUT("/:id", s.UpdateSensorHandler)
		sensors.DELETE("/:id", s.DeleteSensorHandler)
	}
	return r
}

func (s *Server) GetSensorsHandler(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	nextOffset := offset + limit

	sensors, err := s.db.GetSensors(c, db.GetSensorsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if sensors == nil {
		nextOffset = 0
	}

	resp := gin.H{
		"message":   fmt.Sprintf("Found %d sensors", len(sensors)),
		"data":      sensors,
		"next_page": nextOffset,
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetSensorByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Sensor for id %v not found", id)})
		return
	}

	sensor, err := s.db.GetSensorById(c, int32(id))
	resp := gin.H{
		"message": fmt.Sprintf("Found sensor %d", sensor.ID),
		"data":    sensor,
	}

	c.JSON(http.StatusOK, resp)
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
