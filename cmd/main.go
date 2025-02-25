package main

import (
	"products-api/controller"
	"products-api/usecase"

	"github.com/gin-gonic/gin"
)

func main()  {
	
	server := gin.Default()

	ProductUseCase := usecase.NewProductUseCase()
	//Controller layers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}