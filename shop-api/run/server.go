package main

import (
	"github.com/gin-gonic/gin"
	"github.com/johnnydacosta/apotheose-shop/shop-api/internal/product"
)

func main() {
	productController := &product.ProductController{}

	r := gin.Default()

	// Product resource handlers
	r.GET("/products", productController.Find)
	r.POST("/products", productController.Create)
	r.GET("/products/:id", productController.FindById)
	r.PUT("/products/:id", productController.Update)
	r.DELETE("/products/:id", productController.DeleteById)

	r.Run() // listen and serve on 0.0.0.0:8080
}
