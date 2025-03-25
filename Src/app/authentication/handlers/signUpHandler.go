package handlers

import (
    "authentication/models"
    "authentication/service"
    "net/http"

    "github.com/gin-gonic/gin"
)

type SignUpHandler struct {
    Service *service.SignUpService
}

func (h *SignUpHandler) SignUp(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    msg, err := h.Service.SignUp(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": msg}) 
    
}