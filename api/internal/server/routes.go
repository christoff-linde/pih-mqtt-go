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

	sensor_metadata := v1.Group("/sensor_metadata")
	{
		sensor_metadata.GET("/", s.GetSensorMetadataHandler)
		sensor_metadata.GET("/:id", s.GetSensorMetadataByIdHandler)
		sensor_metadata.POST("/", s.CreateSensorMetadataHandler)
		sensor_metadata.PUT("/:id", s.UpdateSensorMetadataHandler)
		sensor_metadata.DELETE("/:id", s.DeleteSensorMetdataHandler)
	}
	return r
}

func (s *Server) GetSensorMetadataHandler(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	nextOffset := offset + limit

	sensor_metadata, err := s.db.GetSensorMetadata(c, db.GetSensorMetadataParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if sensor_metadata == nil {
		nextOffset = 0
	}

	resp := gin.H{
		"message":   fmt.Sprintf("Found %d sensor_metadata", len(sensor_metadata)),
		"data":      sensor_metadata,
		"next_page": nextOffset,
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetSensorMetadataByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("Invalid sensor_id parameter", err)})
		return
	}

	sensor_metadata, err := s.db.GetSensorMetadataForSensorId(c, int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Sensor Metadata for sensor_id %v not found", id)})
		return
	}
	resp := gin.H{
		"message": fmt.Sprintf("Found sensor_metadata %d", sensor_metadata.ID),
		"data":    sensor_metadata,
	}

	c.JSON(http.StatusOK, resp)
}

func (s *Server) CreateSensorMetadataHandler(c *gin.Context) {
	var sensorMetadata db.CreateSensorMetadataParams

	if err := c.BindJSON(&sensorMetadata); err != nil {
		resp := gin.H{
			"message": "Invalid input",
			"error":   err,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	sensor_metadata, err := s.db.CreateSensorMetadata(c, db.CreateSensorMetadataParams{
		SensorID:       sensorMetadata.SensorID,
		SensorType:     sensorMetadata.SensorType,
		Manufacturer:   sensorMetadata.Manufacturer,
		ModelNumber:    sensorMetadata.ModelNumber,
		SensorLocation: sensorMetadata.SensorLocation,
		AdditionalData: sensorMetadata.AdditionalData,
	})
	if err != nil {
		resp := gin.H{
			"message": "Failed to create Sensor Metadata",
			"error":   err,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := gin.H{
		"message": fmt.Sprintf("Created sensor_metadata %v for sensor_id %v", sensor_metadata.ID, sensorMetadata.SensorID),
		"data":    sensor_metadata,
	}
	c.JSON(http.StatusCreated, resp)
}

func (s *Server) UpdateSensorMetadataHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var sensorMetadata db.UpdateSensorMetadataParams
	if err := c.BindJSON(&sensorMetadata); err != nil {
		resp := gin.H{
			"message": "Invalid input",
			"error":   err,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	sensor_metadata, err := s.db.UpdateSensorMetadata(c, db.UpdateSensorMetadataParams{
		SensorID:       sensorMetadata.SensorID,
		SensorType:     sensorMetadata.SensorType,
		Manufacturer:   sensorMetadata.Manufacturer,
		ModelNumber:    sensorMetadata.ModelNumber,
		SensorLocation: sensorMetadata.SensorLocation,
		AdditionalData: sensorMetadata.AdditionalData,
	})

	if sensor_metadata.RowsAffected() == 0 {
		// TODO: improve error handling
		resp := gin.H{
			"message": "Failed to update sensor_metadata",
			"error":   fmt.Errorf("Sensor_metadata with id %v or sensor_id %v does not exist", id, sensorMetadata.SensorID),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if err != nil {
		resp := gin.H{
			"message": "Failed to update sensor_metadata",
			"error":   err,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := gin.H{
		"message": fmt.Sprintf("Updated sensor_metadata %v", id),
		"data":    sensor_metadata,
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) DeleteSensorMetdataHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Sensor Metadata for id %v not found", id)})
		return
	}

	sensor_metadata, err := s.db.DeleteSensorMetadata(c, int32(id))

	if sensor_metadata.RowsAffected() == 0 {
		// TODO: improve error handling
		resp := gin.H{
			"message": "Failed to delete sensor_metadata",
			"error":   fmt.Errorf("Sensor Metadtata with id %v does not exist", id),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if err != nil {
		resp := gin.H{
			"message": "Failed to delete sensor_metadata",
			"error":   err,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := gin.H{
		"message": fmt.Sprintf("Deleted sensor_metadata %v", id),
		"data":    sensor_metadata,
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
