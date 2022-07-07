package main

import (
	"crud-golang/config"
	"crud-golang/handler"
	"crud-golang/product"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func main() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Infof(".env is not loaded properly")
	}

	cfg := config.NewConfig()
	cfg.DB().Master().AutoMigrate(&product.Product{})
	fmt.Printf("Application : %s \nEnvironment : %s\n", cfg.ServiceName(), cfg.ENV())

	productRepository := product.NewRepository(cfg)
	productService := product.NewService(productRepository)
	productHandler := handler.NewHandler(productService)

	router := gin.Default()
	api := router.Group("/api/v1")

	//health check
	api.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})
	// endpoint product
	api.POST("/products", productHandler.CreateProductHandler)
	api.GET("/products", productHandler.GetAllProductHandler)
	api.GET("/products/:id", productHandler.GetProductByIDHandler)
	api.PUT("/products/:id", productHandler.UpdateProductHandler)
	api.DELETE("/products/:id", productHandler.DeleteProductHandler)
	err = router.Run(fmt.Sprintf(":%d", cfg.Port()))
	if err != nil {
		log.Panic(err)
	}
}
