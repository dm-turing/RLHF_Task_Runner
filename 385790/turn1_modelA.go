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
		Info:                    swaggerInfo,
		BasePath:                "/api",
		JSONRootPath:            "",
		SkipNonPublic:           true,
		ExcludeSwagger:          true,
		SwaggerFile:             "swagger.json",
		SwaggerPath:             "/swagger",
		SwaggerJSON:             "/swagger.json",
		SwaggerUIDoc:            "/swagger-ui",
		RedocPath:               "/redoc",
		RedocURL:                "https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.bundle.js",
		IgnoreRouteTags:         false,
		EnableRedoc:             true,
		EnableSwaggerUI:         true,
		WriteSwagger:            true,
		WriteSwaggerJSON:        true,
		WriteSwaggerUIDoc:       true,
		WriteRedoc:              true,
		WriteSwaggerJS:          true,
		SwaggerJS:               "/swagger.js",
		WriteSwaggerCSS:         true,
		SwaggerCSS:              "/swagger.css",
		WriteSwaggerFavicon:     true,
		SwaggerFavicon:          "/favicon.ico",
		WriteSwaggerSpec:        true,
		SwaggerSpec:             "/swagger.yaml",
		WriteRedocSpec:          true,
		RedocSpec:               "/redoc.yaml",
		SwaggerSpecDepth:        0,
		RedocSpecDepth:          0,
		EnableSwaggerValidation: true,
	}

	// Register the Swagger handler
	swag.Register("r", swaggerOptions)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
