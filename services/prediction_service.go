package services

import (
	"context"

	models "github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	repositories "github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
)

type (
	PredictionService interface {
		Service
		Predict(ctx context.Context, req models.PredictionRequest) (audio []byte, text string)
	}

	predictionService struct {
		*service[repositories.ChatHistoryRepository]
		replicateService ReplicateService
		openAIService    OpenAIService
	}
)

func NewPredictionService(replicateService ReplicateService, openAIService OpenAIService) PredictionService {
	return &predictionService{
		service:          &service[repositories.ChatHistoryRepository]{},
		replicateService: replicateService,
		openAIService:    openAIService,
	}
}

func (s *predictionService) Predict(ctx context.Context, req models.PredictionRequest) (audio []byte, text string) {
	sttOutput := s.openAIService.SpeechToText(ctx, req.AudioQuestionFile, req.AudioQuestionFilename)
	if s.openAIService.Error() != nil {
		s.ThrowsException(&s.exception.BadRequest, "Failed to generate speech to text!")
		s.ThrowsError(s.openAIService.Error())
		return nil, ""
	}

	replicateOutput := s.replicateService.AskImage(ctx, req.ImageFile, req.ImageFileName, sttOutput)
	if s.replicateService.Error() != nil {
		s.ThrowsException(&s.exception.ReplicateConnectionRefused, "Replicate Connection Refused!")
		s.ThrowsError(s.replicateService.Error())
		return nil, ""
	}

	ttsOutput := s.openAIService.TextToSpeech(ctx, replicateOutput)
	if s.openAIService.Error() != nil {
		s.ThrowsException(&s.exception.FailedGenerateAudio, "Failed to convert audio output!")
		s.ThrowsError(s.openAIService.Error())
		return nil, ""
	}

	savePrediction := s.repository.SaveChatHistory(ctx, req.ImageFileName, sttOutput, replicateOutput)
	if s.ThrowsRepoException() {
		return nil, ""
	}

	return ttsOutput, savePrediction.Answer
}
