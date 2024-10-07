package server

import (
	"net/http"

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
