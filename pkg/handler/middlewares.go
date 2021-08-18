package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := h.services.Authorization.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}
		c.Next()
	}
}
