package controller

import (
	models "github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	services "github.com/abdanhafidz/ai-visual-multi-modal-backend/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetProfile(gCtx *gin.Context)
	UpdateProfile(gCtx *gin.Context)
}

type userController struct {
	*controller[services.UserProfileService]
}

func NewUserController(userProfileService services.UserProfileService) UserController {
	return &userController{
		controller: &controller[services.UserProfileService]{
			service: userProfileService,
		},
	}
}
func (c *userController) GetProfile(ctx *gin.Context) {
	c.HeaderParse(ctx)
	userProfile := c.service.Retrieve(ctx.Request.Context(), uint(c.accountData.UserID))
	c.Response(ctx, userProfile)
}

func (c *userController) UpdateProfile(ctx *gin.Context) {
	var updateProfileRequest models.AccountDetails
	c.RequestJSON(ctx, updateProfileRequest)
	updatedProfile := c.service.Update(ctx.Request.Context(), uint(c.accountData.UserID), updateProfileRequest)
	c.Response(ctx, updatedProfile)
}
