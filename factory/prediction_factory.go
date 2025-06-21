package factory

import (
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/config"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/controller"
	repositories "github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/services"
)

func NewPredictionModule() controller.PredictionController {
	chatHistoryRepository := repositories.NewChatHistoryRepository()
	openAIService := services.NewOpenAIService(chatHistoryRepository, config.OpenAIClient)
	replicateService := services.NewReplicateService(chatHistoryRepository, config.ReplicateClient, "lucataco/moondream2:72ccb656353c348c1385df54b237eeb7bfa874bf11486cf0b9473e691b662d31")
	predictionService := services.NewPredictionService(chatHistoryRepository, replicateService, openAIService)
	predictionController := controller.NewPredictionController(predictionService)

	return predictionController
}
