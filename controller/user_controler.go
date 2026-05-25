package controller

import (
	"net/http"

	"github.com/alexander-pastana/go-api-lab/model"
	"github.com/alexander-pastana/go-api-lab/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Usecase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) UserController {
	return UserController{
		Usecase: usecase,
	}
}

func (uc *UserController) SignUp(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = uc.Usecase.SignUp(user)
	if err != nil {
		if err.Error() == "Usuário já cadastrado" {
			//Cria um map com o gin, para adicionar o erro no json
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		// Se for qualquer outro erro (tipo o banco de dados ter caído)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno no servidor"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Message": "Cadastro efetuado com sucesso"})

}

func (uc *UserController) SignIn(ctx *gin.Context) {
	//Trata o json trazido pelo cliente
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	token, err := uc.Usecase.SignIn(user)
	if err != nil {
		if err.Error() == "usuário ou senha inválidos" {
			//Cria um map com o gin, para adicionar o erro no json
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno no servidor"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}
