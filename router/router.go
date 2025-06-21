package router

import (
	config "github.com/abdanhafidz/ai-visual-multi-modal-backend/config"
	controller "github.com/abdanhafidz/ai-visual-multi-modal-backend/controller"
	"github.com/gin-gonic/gin"
)

func StartService() {
	router := gin.Default()
	router.GET("/", controller.HomeController)
	PredictionRoute(router)
	router.Run(config.TCP_ADDRESS)
}
