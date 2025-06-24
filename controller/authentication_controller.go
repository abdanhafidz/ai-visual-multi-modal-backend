package controller

import (
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	services "github.com/abdanhafidz/ai-visual-multi-modal-backend/services"
	"github.com/gin-gonic/gin"
)

type AuhenticationController interface {
	Controller
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}
type authenticationController struct {
	*controller[services.AuthenticationService]
}

func NewAuthenticationController(authenticationService services.AuthenticationService) AuhenticationController {
	return &authenticationController{
		controller: &controller[services.AuthenticationService]{service: authenticationService},
	}
}
func (c *authenticationController) Register(ctx *gin.Context) {
	var loginRequest models.LoginRequest
	c.RequestJSON(ctx, &loginRequest)
	if loginRequest.IPAddress == "" {
		loginRequest.IPAddress = ctx.ClientIP()
	}
	token := c.service.Register(ctx.Request.Context(), loginRequest.PassPhrase, loginRequest.TurnStile, loginRequest.IPAddress)

	c.Response(ctx, token)
}
func (c *authenticationController) Login(ctx *gin.Context) {
	var loginRequest models.LoginRequest
	c.RequestJSON(ctx, &loginRequest)

	token := c.service.Login(ctx.Request.Context(), loginRequest.PassPhrase)

	c.Response(ctx, token)
}
