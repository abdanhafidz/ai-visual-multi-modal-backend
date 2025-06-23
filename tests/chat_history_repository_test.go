package tests

import (
	"context"
	"testing"

	"github.com/abdanhafidz/ai-visual-multi-modal-backend/config"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
)

func TestSaveChatHistoryRepository(t *testing.T) {
	config.RunConfig()
	var ctx context.Context
	ctx = context.Background()
	t.Log("DB Ptr :", config.DB)
	chatHistoryRepo := repositories.NewChatHistoryRepository(config.DB)
	t.Log("Repo Ptr :", chatHistoryRepo)
	result := chatHistoryRepo.SaveChatHistory(ctx, "Testing", "Testing", "Testing")
	t.Log("Error Log:", chatHistoryRepo.RowsError())
	expectedResult := models.ChatHistory{
		Answer:    "Testing",
		ImagePath: "Testing",
		Question:  "Testing",
	}
	if chatHistoryRepo.IsNoRecord() {
		t.Log("Is No Record:", chatHistoryRepo.IsNoRecord())
		t.Errorf("Failed to create rows!")
	}
	err := chatHistoryRepo.RowsError()
	if err != nil {
		// t.Logf("Expected Result: %v, Got: %v", expectedResult, result)
		t.Error("Error while saving chat history:", err)
	}
	if result.Answer != expectedResult.Answer || result.ImagePath != expectedResult.ImagePath || result.Question != expectedResult.Question {
		t.Logf("Expected Result: %v, Got: %v", expectedResult, result)
		t.Errorf("Wrong Result of chat history")

		return
	}
}
