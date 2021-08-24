package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getUser(c *gin.Context) {
	toker, err := h.services.Authorization.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	foundAuth, err := h.services.Authorization.FetchAuth(toker)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	userTable, err := h.services.Users.GetUserByUserId(foundAuth.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":           userTable.UserId,
		"user_name":         userTable.Name,
		"user_username":     userTable.Username,
		"user_email":        userTable.Email,
		"user_avatar_image": userTable.UserAvatarImage,
		"user_city":         userTable.UserCity,
		"user_age":          userTable.UserAge,
	})
}

type UpdateUserInfo struct {
	UserAge         *int    `json:"user_age"`
	UserAvatarImage *string `json:"user_avatar_image"`
	UserCity        *string `json:"user_city"`
	Name            *string `json:"name"`
	Username        *string `json:"username"`
	Email           *string `json:"email"`
}

func (h *Handler) updateUser(c *gin.Context) {
	token, err := h.services.Authorization.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	foundAuth, err := h.services.Authorization.FetchAuth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var u UpdateUserInfo
	if err := c.ShouldBindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userInfoTable, err := h.services.Users.GetUserInfo(foundAuth.UserID)
	userTable, err := h.services.Users.GetUser(foundAuth.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	updateUserTable, err := h.services.Users.UpdateUser(u.UserAge, u.UserAvatarImage, u.UserCity, userInfoTable, u.Username, u.Name, u.Email, userTable)
	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":           updateUserTable.UserId,
		"user_name":         updateUserTable.Name,
		"user_username":     updateUserTable.Username,
		"user_email":        updateUserTable.Email,
		"user_avatar_image": updateUserTable.UserAvatarImage,
		"user_city":         updateUserTable.UserCity,
		"user_age":          updateUserTable.UserAge,
	})
}
