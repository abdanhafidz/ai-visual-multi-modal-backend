package repositories

import (
	"context"

	"github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
)

type ChatHistoryRepository interface {
	Repository
	SaveChatHistory(ctx context.Context, imagePath string, question string, answer string) (res models.ChatHistory)
}

type chatHistoryRepository struct {
	*repository[models.ChatHistory]
}

func (r *chatHistoryRepository) SaveChatHistory(ctx context.Context, imagePath string, question string, answer string) (res models.ChatHistory) {
	r.entity.ImagePath = imagePath
	r.entity.Question = question
	r.entity.Answer = answer
	r.Create(ctx)
	return r.entity
}
