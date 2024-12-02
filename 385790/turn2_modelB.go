package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

func main() {
	r := gin.Default()

	// Register your Gin routes
	// ...

	// Set up Swagger documentation
	swaggerInfo := &swag.Info{
		Title:       "My API",
		Version:     "1.0.0",
		Description: "A simple API",
	}

	swaggerFile, err := swag.SaveJSON(swaggerInfo)
	if err != nil {
		panic(err)
	}

	r.GET("/swagger/*any", gin.WrapH(middleware.Swagger(swaggerFile)))

	fmt.Println("Server is running on http://localhost:8080")
}
