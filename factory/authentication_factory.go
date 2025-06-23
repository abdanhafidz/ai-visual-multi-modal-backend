package factory

import (
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/config"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/controller"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/services"
)

func NewAuthenticationModule() controller.AuhenticationController {
	accountRepository := repositories.NewAccountRepository(config.DB)
	jwtService := services.NewJWTService(accountRepository, config.Salt)
	authenticationService := services.NewAuthenticationService(accountRepository, config.TurnstileClient, jwtService)
	authenticationController := controller.NewAuthenticationController(authenticationService)
	return authenticationController
}
