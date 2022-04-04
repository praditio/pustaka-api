package handler

import (
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct {
	loginService book.LoginService
	jWtService   book.JWTService
}

func LoginHandler(loginService book.LoginService,
	jWtService book.JWTService) *loginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) {
	var credential book.LoginCredentialsRequest
	var token string
	err := ctx.ShouldBindJSON(&credential)
	if err != nil {
		token = "Please nominate a valid email and password"
	}

	isUserAuthenticated := controller.loginService.LoginUser(credential)
	if isUserAuthenticated {
		token = controller.jWtService.GenerateToken(credential.Email, true)

	}

	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"token": nil,
		})
	}

}
