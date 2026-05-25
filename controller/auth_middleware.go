package controller

import (
	"net/http"
	"strings"

	"github.com/alexander-pastana/go-api-lab/usecase"
	"github.com/gin-gonic/gin"
)

func Auth(userUseCase usecase.UserUseCase) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        token := ctx.GetHeader("Authorization")
        if token == "" {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
            ctx.Abort()
            return
        }

        cleanToken := strings.TrimPrefix(token, "Bearer ")
        
        // Chama a usecase para validar o token JWT
        err := userUseCase.ValidateToken(cleanToken)
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
            ctx.Abort()
            return
        }

        // Se o token for válido, continua a requisição normalmente
        ctx.Next()
    }
}

