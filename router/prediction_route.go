package router

import (
	factory "github.com/abdanhafidz/ai-visual-multi-modal-backend/factory"
	"github.com/gin-gonic/gin"
)

func PredictionRoute(router *gin.Engine) {
	routerGroup := router.Group("/api/v1")
	{
		routerGroup.POST("/predict", func(c *gin.Context) {
			predictionModule := factory.NewPredictionModule()
			predictionModule.Predict(c)
		})
	}
}
