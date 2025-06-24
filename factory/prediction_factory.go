package factory

import (
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/config"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/controller"
	repositories "github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/services"
)

func NewPredictionModule() controller.PredictionController {
	chatHistoryRepository := repositories.NewChatHistoryRepository(config.DB)
	openAIService := services.NewOpenAIService(
		chatHistoryRepository,
		config.OpenAIClient,
	)
	replicateService := services.NewReplicateService(
		chatHistoryRepository,
		config.ReplicateClient,
		"spuuntries",
		"kp3l",
	)
	predictionService := services.NewPredictionService(
		chatHistoryRepository,
		replicateService,
		openAIService,
	)
	predictionController := controller.NewPredictionController(predictionService)
	return predictionController
}
