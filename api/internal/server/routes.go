package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

type GetSensorOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello World!" doc:"Greeting message"`
	}
}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	// Serve OpenAPI documentation
	//     r.GET("/docs", func(c *gin.Context) {
	//         c.Header("Content-Type", "text/html")
	//         c.Writer.Write([]byte(`<!doctype html>
	// <html>
	//   <head>
	//     <title>API Reference</title>
	//     <meta charset="utf-8" />
	//     <meta name="viewport" content="width=device-width, initial-scale=1" />
	//   </head>
	//   <body>
	//     <script id="api-reference" data-url="/openapi.json"></script>
	//     <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
	//   </body>
	// </html>`))
	//     })

	config := huma.DefaultConfig("pih-core-go", "0.0.1")
	config.DocsPath = ""
	config.Info = &huma.Info{
		Title:       "PIH Core API",
		Description: "API for the PIH Core project",
		Version:     "0.0.1",
		License: &huma.License{
			Name: "MIT",
			URL:  "https://opensource.org/licenses/MIT",
		},
	}
	config.Tags = []*huma.Tag{
		{
			Name:        "Sensor",
			Description: "Operations related to sensors",
			ExternalDocs: &huma.ExternalDocs{
				Description: "Find more info here",
				URL:         "https://example.com/docs/sensor",
			},
		},
		{
			Name:        "SensorMetadata",
			Description: "Operations related to sensor metadata",
			ExternalDocs: &huma.ExternalDocs{
				Description: "Find more info here",
				URL:         "https://example.com/docs/sensor_metadata",
			},
		},
		{
			Name:        "SensorReading",
			Description: "Operations related to sensor readings",
			ExternalDocs: &huma.ExternalDocs{
				Description: "Find more info here",
				URL:         "https://example.com/docs/sensor_reading",
			},
		},
	}

	// config.OpenAPIPath = ""

	// Serve OpenAPI documentation
	r.GET("/v1/docs", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.Writer.Write([]byte(`
			<!doctype html>
			<html>
			<head>
				<title>API Reference</title>
				<meta charset="utf-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
			</head>
			<body>
				<script id="api-reference" data-url="openapi.json"></script>
				<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
			</body>
			</html>
		`))
	})

	v1 := r.Group("/v1")

	// Initialize humagin with the Gin router group
	// api := humagin.New(r, config)
	api := humagin.NewWithGroup(r, v1, config)

	// Define API operations
	huma.Register(api, huma.Operation{
		OperationID: "helloWorld",
		Method:      "GET",
		Path:        "/",
		Summary:     "Hello World",
		Description: "Returns a hello world message.",
	}, func(ctx context.Context, input *struct{}) (*GetSensorOutput, error) {
		resp := &GetSensorOutput{}
		resp.Body.Message = "Hello World!"
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "getSensors",
		Method:      "GET",
		Path:        "/sensor/get_sensors",
		Summary:     "Get a list of sensors",
		Description: "Returns a list of sensors.",
		Tags:        []string{"Sensor"},
	}, func(ctx context.Context, input *struct{}) (*GetSensorOutput, error) {
		// Implement your handler logic here
		resp := &GetSensorOutput{}
		resp.Body.Message = "List of sensors"
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "getSensorById",
		Method:      "GET",
		Path:        "/sensor/get_sensor_by_id",
		Summary:     "Get a sensor by ID",
		Description: "Returns a sensor by its ID.",
		Tags:        []string{"Sensor"},
	}, func(ctx context.Context, input *struct {
		SensorID string `query:"sensor_id" doc:"The ID of the sensor"`
	},
	) (*GetSensorOutput, error) {
		// Implement your handler logic here
		resp := &GetSensorOutput{}
		resp.Body.Message = fmt.Sprintf("Found sensor %s", input.SensorID)
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "createSensor",
		Method:      "POST",
		Path:        "/sensor/create_sensor",
		Summary:     "Create a new sensor",
		Description: "Creates a new sensor.",
		Tags:        []string{"Sensor"},
	}, func(ctx context.Context, input *struct {
		// Define your input fields here
	},
	) (*GetSensorOutput, error) {
		// Implement your handler logic here
		resp := &GetSensorOutput{}
		resp.Body.Message = "Sensor created"
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "updateSensor",
		Method:      "PUT",
		Path:        "/sensor/update_sensor",
		Summary:     "Update a sensor",
		Description: "Updates an existing sensor.",
		Tags:        []string{"Sensor"},
	}, func(ctx context.Context, input *struct {
		SensorID string `query:"sensor_id" doc:"The ID of the sensor"`
		// Define other input fields here
	},
	) (*GetSensorOutput, error) {
		// Implement your handler logic here
		resp := &GetSensorOutput{}
		resp.Body.Message = fmt.Sprintf("Sensor %s updated", input.SensorID)
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "deleteSensor",
		Method:      "DELETE",
		Path:        "/sensor/delete_sensor",
		Summary:     "Delete a sensor",
		Description: "Deletes an existing sensor.",
		Tags:        []string{"Sensor"},
	}, func(ctx context.Context, input *struct {
		SensorID string `query:"sensor_id" doc:"The ID of the sensor"`
	},
	) (*GetSensorOutput, error) {
		// Implement your handler logic here
		resp := &GetSensorOutput{}
		resp.Body.Message = fmt.Sprintf("Sensor %s deleted", input.SensorID)
		return resp, nil
	})

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
