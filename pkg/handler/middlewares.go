package handler

import (
	"garyshker"
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

func (h *Handler) requireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		foundAuth, err := h.tokenCheck(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, "You don't have access")
			c.Abort()
			return
		}
		role, err := h.services.Users.GetRole(foundAuth.UserID)
		if err != nil {
			c.JSON(http.StatusNotFound, "role error")
			c.Abort()
			return
		}
		if role != garyshker.AdminRole {
			c.JSON(http.StatusBadRequest, "You don't have access")
			c.Abort()
			return
		}
		c.Next()
	}
}
