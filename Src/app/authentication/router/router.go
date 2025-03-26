package router

import (
	"authentication/constants"
	"authentication/handlers"
	"authentication/repo"
	"authentication/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	repoInstance := &repo.SignUpRepository{DB: db}
	serviceInstance := &service.SignUpService{Repo: repoInstance}
	handlerInstance := &handlers.SignUpHandler{Service: serviceInstance}

	repoInstanceSignin := &repo.SignInRepository{DB: db}
	serviceInstanceSignin:= &service.SignInService{Repo: repoInstanceSignin}
	handlerInstanceSignin:=&handlers.SignInHandler{Service: serviceInstanceSignin}

	r := gin.Default()

	r.POST(constants.SignUpRoute, handlerInstance.SignUp)
	r.POST(constants.SignInRoute, handlerInstanceSignin.SignIn)

	return r
}
