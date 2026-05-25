package main

import (
	"log"
	"os"

	"github.com/alexander-pastana/go-api-lab/controller"
	"github.com/alexander-pastana/go-api-lab/db"
	"github.com/alexander-pastana/go-api-lab/repository"
	"github.com/alexander-pastana/go-api-lab/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Puxa a chave secreta do arquivo .env
	secretKey := os.Getenv("SECRET_KEY")
	//fmt.Println("Chave Secreta carregada com sucesso:", secretKey)

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	UserRepository := repository.NewUserRepository(dbConnection)
	UserUsecase := usecase.NewUserUseCase(UserRepository, secretKey)
	UserController := controller.NewUserController(UserUsecase)

	userRoutes := server.Group("/users")
	{
		// O Gin junta o "/users" como prefixo do grupo com a rota informada
		userRoutes.POST("/signup", UserController.SignUp)
		userRoutes.POST("/signin", UserController.SignIn)
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

	productRoutes := server.Group("/products")
	productRoutes.Use(controller.Auth(UserUsecase))
	{
		productRoutes.GET("/", productController.GetProducts)
		productRoutes.POST("/", productController.CreateProduct)
		productRoutes.GET("/:productId", productController.GetProductById)
		productRoutes.PUT("/:productId", productController.UpdateProduct)
		productRoutes.DELETE("/:productId", productController.DeleteProduct)
	}
	server.Run(":8000")
}
