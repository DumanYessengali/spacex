package handler

import (
	"garyshker"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var u garyshker.User
	if err := c.ShouldBindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, "invalid json")
		return
	}

	user, err := h.services.Authorization.CreateUser(&u)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

type singInInput struct {
	UsernameOrEmail string `json:"username or email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var u singInInput
	if err := c.ShouldBindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	//check if the user exist:
	user, err := h.services.Authorization.GetUser(u.UsernameOrEmail, u.Password)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	//since after the user logged out, we destroyed that record in the database so that same jwt token can't be used twice. We need to create the token again

	authData, err := h.services.Authorization.CreateAuth(user.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var authD garyshker.AuthDetails
	authD.UserId = authData.UserID
	authD.AuthUuid = authData.AuthUUID

	token, loginErr := h.services.Authorization.SignIn(authD)
	if loginErr != nil {
		newErrorResponse(c, http.StatusForbidden, "Please try to login later")
		return
	}
	c.JSON(http.StatusOK, token)
}

func (h *Handler) LogOut(c *gin.Context) {
	au, err := h.services.Authorization.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	delErr := h.services.Authorization.DeleteAuth(au)
	if delErr != nil {
		log.Println(delErr)
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
