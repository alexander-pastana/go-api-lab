package controller

import (
	"net/http"
	"strconv"

	"github.com/alexander-pastana/go-api-lab/model"
	"github.com/alexander-pastana/go-api-lab/usecase"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	//Lê o JSON trazido da página, da requisição
	err := ctx.BindJSON(&product)
	if err != nil {
		//Retorna uma mensagem caso dê erro
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	//Chamada da função criada na usecase
	insertedProduct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	//http.StatusCreated (código 201). O 201 é o padrão quando você cria algo novo
	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	//Verifica se id não é uma string vazia
	//Desnecessário pois o Atoi já cuida disso
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	//Converte string para int
	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	//Converte string para int
	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product
	err = ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido ou campos obrigatórios ausentes"})
		return
	}

	product.ID = productId

	updatedProduct, err := p.productUsecase.UpdateProduct(product)
	if err != nil {
		//compara string com string)
		if err.Error() == "Produto não encontrado" {
			//Cria um map com o gin, para adicionar o erro no json
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno no servidor"})
		return
	}
	//quando você altera algo existente (como no Update), o padrão mais comum é usar http.StatusOK (código 200).
	ctx.JSON(http.StatusOK, updatedProduct)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	id_product, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUsecase.DeleteProduct(id_product)
	if err != nil {
		if err.Error() == "Produto não encontrado" {
			//Cria um map com o gin, para adicionar o erro no json
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar produto"})
        return

	}

	response := model.Response{
		Message: "Produto deletado com sucesso!",
	}
	ctx.JSON(http.StatusOK, response)
}

// products := []model.Product{
// 	{
// 		ID: 1,
// 		Name: "Batata frita",
// 		Price: 20,
// 	},

// }
