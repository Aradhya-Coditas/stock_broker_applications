package handlers

import (
	"authentication/models"
	"authentication/service"

	"github.com/gin-gonic/gin"
)

type SignUpHandler struct {
	Service *service.SignUpService
}

func (h *SignUpHandler) SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	msg, err := h.Service.SignUp(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": msg})
}
