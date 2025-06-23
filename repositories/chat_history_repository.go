package repositories

import (
	"context"

	"github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	"gorm.io/gorm"
)

type ChatHistoryRepository interface {
	Repository
	SaveChatHistory(ctx context.Context, imagePath string, question string, answer string) (res models.ChatHistory)
}

type chatHistoryRepository struct {
	*repository[models.ChatHistory]
}

func NewChatHistoryRepository(db *gorm.DB) ChatHistoryRepository {
	return &chatHistoryRepository{
		repository: &repository[models.ChatHistory]{
			entity:      models.ChatHistory{},
			transaction: db,
		},
	}
}
func (r *chatHistoryRepository) SaveChatHistory(ctx context.Context, imagePath string, question string, answer string) (res models.ChatHistory) {
	r.entity.ImagePath = imagePath
	r.entity.Question = question
	r.entity.Answer = answer
	r.Create(ctx)
	return r.entity
}
