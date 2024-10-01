package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	db "github.com/christoff-linde/pih-core-go/api/database"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
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
	var sensorData db.CreateSensorParams
	if err := c.BindJSON(&sensorData); err != nil {
		resp := gin.H{
			"message": "Invalid input",
			"error":   err,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	sensor, err := s.db.CreateSensor(c, db.CreateSensorParams{
		SensorName: sensorData.SensorName,
		CreatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		resp := gin.H{
			"message": "Failed to create Sensor",
			"error":   err,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := gin.H{
		"message": fmt.Sprintf("Created sensor %v", sensor.SensorName),
		"data":    sensor,
	}
	c.JSON(http.StatusCreated, resp)
}

func (s *Server) UpdateSensorHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func (s *Server) DeleteSensorHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Sensor for id %v not found", id)})
		return
	}

	sensor, err := s.db.DeleteSensor(c, int32(id))

	if sensor.RowsAffected() == 0 {
		// TODO: improve error handling
		resp := gin.H{
			"message": "Failed to delete sensor",
			"error":   fmt.Errorf("Sensor with id %v does not exist", id),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if err != nil {
		resp := gin.H{
			"message": "Failed to delete sensor",
			"error":   err,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := gin.H{
		"message": fmt.Sprintf("Deleted sensor %v", id),
		"data":    sensor,
	}
	c.JSON(http.StatusOK, resp)
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
