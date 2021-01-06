package middleware

import (
	"github.com/gin-gonic/gin"
	"go-sharex/config"
	"go-sharex/er"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		if key := ctx.Query("key"); len(key) > 0 && config.Get().Keys[key] != "" {
			ctx.Next() // key is valid go to next handler
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": er.AuthInvalidKey}) // no key provided / key is invalid
			return
		}
	}
}