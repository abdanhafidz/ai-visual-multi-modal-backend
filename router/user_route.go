package router

import (
	factory "github.com/abdanhafidz/ai-visual-multi-modal-backend/factory"
	middleware "github.com/abdanhafidz/ai-visual-multi-modal-backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	userModule := factory.NewUserProfileModule()
	routerGroup := router.Group("/api/v1/user")
	{
		routerGroup.GET("/me", middleware.AuthUser, userModule.GetProfile)
		routerGroup.PUT("/me", middleware.AuthUser, userModule.UpdateProfile)
	}
}
