package factory

import (
	controller "github.com/abdanhafidz/ai-visual-multi-modal-backend/controller"
	repositories "github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	services "github.com/abdanhafidz/ai-visual-multi-modal-backend/services"
)

func NewAuthenticationModule() controller.AuthController {

	accountRepository := repositories.NewAccountRepository()
	accountDetailRepository := repositories.NewAccountDetailRepository()
	userProfileService := services.NewUserProfileService(accountDetailRepository)
	accountService := services.NewAuthenticationService(accountRepository, userProfileService)
	accountController := controller.NewAuthController(accountService)

	return accountController
}

func NewUserProfileModule() controller.UserController {

	accountDetailRepository := repositories.NewAccountDetailRepository()
	userProfileService := services.NewUserProfileService(accountDetailRepository)
	userController := controller.NewUserController(userProfileService)

	return userController
}
