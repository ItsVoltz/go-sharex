package main

import (
	"github.com/gin-gonic/gin"
	_ "go-sharex/config"
	"go-sharex/middleware"
	"go-sharex/routes"
	"log"
)

func main() {
	gin.SetMode("release") // Set gin to release
	router := gin.Default()

	/* Load Templates */
	router.LoadHTMLGlob("./templates/*")

	/* Define Routes */
	v1 := router.Group("/v1")
	{
		v1.Use(middleware.Auth())
		v1.POST("/upload", routes.Upload)
	}

	router.GET("/:file", routes.Files)
	//router.StaticFS("/", http.Dir("./uploads")) if you just want static fs uncomment this and comment the /:file GET request

	if err := router.Run(); err != nil { // start webserver on :8080
		log.Fatalf("Failed to start gin: %s", err)
	}
}