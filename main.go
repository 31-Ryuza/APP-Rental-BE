package main

import (
	productcontroller "example/api-rental/Controllers"
	models "example/api-rental/Models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	port := "3000"

	// Konfigurasi middleware CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"} // Ganti dengan origin Anda
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run(":" + port)
}
