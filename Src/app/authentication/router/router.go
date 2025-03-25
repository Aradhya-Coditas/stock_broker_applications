package router

import (
	"authentication/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handlers.SignUpHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/signup", handler.SignUp)
	// r.POST("/signin", handler.SignIn)
	return r
}
