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
	// --- Configurações Iniciais ---
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	secretKey := os.Getenv("SECRET_KEY")

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	server := gin.Default()

	// --- Inicialização de Camadas: Usuário ---
	UserRepository := repository.NewUserRepository(dbConnection)
	UserUsecase := usecase.NewUserUseCase(UserRepository, secretKey)
	UserController := controller.NewUserController(UserUsecase)

	// --- Inicialização de Camadas: Produto ---
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUsecase := usecase.NewProductUseCase(ProductRepository)
	productController := controller.NewProductController(ProductUsecase)

	// --- Rotas ---
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userRoutes := server.Group("/users")
	{
		userRoutes.POST("/signup", UserController.SignUp)
		userRoutes.POST("/signin", UserController.SignIn)
	}

	productRoutes := server.Group("/products")
	productRoutes.Use(controller.Auth(UserUsecase))
	{
		productRoutes.GET("/", productController.GetProducts)
		productRoutes.POST("/", productController.CreateProduct)
		productRoutes.GET("/:productId", productController.GetProductById)
		productRoutes.PUT("/:productId", productController.UpdateProduct)
		productRoutes.DELETE("/:productId", productController.DeleteProduct)
	}

	// --- Execução ---
	server.Run(":8000")
}
