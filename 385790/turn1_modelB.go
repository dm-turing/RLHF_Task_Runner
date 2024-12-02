package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swaggo/swag"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Healthy")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "User data")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/user", userHandler).Methods("GET")

	// Set up Swagger documentation
	swaggerInfo := &swag.Info{
		Title:       "My API",
		Version:     "1.0.0",
		Description: "A simple API with dynamic Swagger documentation",
	}
	swaggerOptions := swag.Options{
		Info:                         swaggerInfo,
		BasePath:                     "/api",
		JSONRootPath:                 "",
		SkipNonPublic:                true,
		ExcludeSwagger:               true,
		SwaggerFile:                  "swagger.json",
		SwaggerPath:                  "/swagger",
		SwaggerJSON:                  "/swagger.json",
		SwaggerUIDoc:                 "/swagger-ui",
		RedocPath:                    "/redoc",
		RedocURL:                     "https://cdn.jsdelivr.net/npm/redoc@next/dist/redoc.min.js",
		EnableSwaggerUI:              true,
		EnableRedoc:                  true,
		EnableSwaggerJSON:            true,
		DisableSwaggerJSONValidation: true,
	}

	// Register the Swagger handler
	swag.Register("r", swaggerOptions)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
