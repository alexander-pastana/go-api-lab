package main

import (
	"github.com/alexander-pastana/go-api-lab/controller"
	"github.com/alexander-pastana/go-api-lab/db"
	"github.com/alexander-pastana/go-api-lab/repository"
	"github.com/alexander-pastana/go-api-lab/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada Usecase
	ProductUsecase := usecase.NewProductUseCase(ProductRepository)
	//Camada de controllers
	productController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})

	})

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductById)
	server.PUT("/product/:productId", productController.UpdateProduct)
	server.DELETE("/product/:productId", productController.DeleteProduct)

	server.Run(":8000")

}
