package middleware

import (
	"net/http"

	"github.com/andrersp/financeiro-go/src/infra/auth"
	"github.com/andrersp/financeiro-go/src/interfaces/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := auth.TokenValid(c.Request)

		if err != nil {
			response.Error(c, http.StatusUnauthorized, err)
			return
		}

		c.Next()

	}
}
