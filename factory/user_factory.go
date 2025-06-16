package factory

import (
	controller "github.com/abdanhafidz/ai-visual-multi-modal-backend/controller"
	repositories "github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	services "github.com/abdanhafidz/ai-visual-multi-modal-backend/services"
)
func NewUserProfileModule() controller.UserController {

	accountDetailRepository := repositories.NewAccountDetailRepository()
	userProfileService := services.NewUserProfileService(accountDetailRepository)
	userController := controller.NewUserController(userProfileService)

	return userController
}
