package router

import (
	factory "github.com/abdanhafidz/ai-visual-multi-modal-backend/factory"
	"github.com/gin-gonic/gin"
)

func AuthenticationRoute(router *gin.Engine) {
	routerGroup := router.Group("/api/v1")
	{
		routerGroup.POST("/register", func(c *gin.Context) {
			authenticationModule := factory.NewAuthenticationModule()
			authenticationModule.Register(c)
		})
		routerGroup.POST("/login", func(c *gin.Context) {
			authenticationModule := factory.NewAuthenticationModule()
			authenticationModule.Login(c)
		})
	}
}
