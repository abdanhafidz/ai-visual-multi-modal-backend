package router

import (
	factory "github.com/abdanhafidz/ai-visual-multi-modal-backend/factory"
	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	authModule := factory.NewAuthenticationModule()
	routerGroup := router.Group("/api/v1/auth")
	{
		routerGroup.POST("/login", authModule.Login)
		routerGroup.POST("/register", authModule.Register)
		// routerGroup.PUT("/change-password", authModule.ChangePassword)
	}
}
