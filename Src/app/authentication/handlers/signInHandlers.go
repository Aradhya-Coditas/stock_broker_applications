package handlers

import (
	"authentication/models"
	"authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignInHandler struct {
	Service *service.SignInService
}

func (h *SignInHandler) SignIn(c *gin.Context) {
	var creds models.SignInRequest
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, models.SignInResponse{
			Error: "Invalid request: " + err.Error(),
		})
		return
	}

	msg, err := h.Service.SignIn(creds.Email, creds.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.SignInResponse{
			Error: err.Error(),
		})
		return
	}

	if msg == "invalid password" || msg == "user not found" {
		c.JSON(http.StatusUnauthorized, models.SignInResponse{
			Error: msg,
		})
		return
	}

	c.JSON(http.StatusOK, models.SignInResponse{
		Message: msg,
	})
}
