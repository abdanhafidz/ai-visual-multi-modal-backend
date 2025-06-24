package tests

import (
	"context"
	"testing"

	"github.com/abdanhafidz/ai-visual-multi-modal-backend/config"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
)

func TestCreateAccountRepository(t *testing.T) {
	config.RunConfig()
	var ctx context.Context
	ctx = context.Background()
	t.Log("DB Ptr :", config.DB)
	accountRepo := repositories.NewAccountRepository(config.DB)
	t.Log("Repo Ptr :", accountRepo)

	result := accountRepo.CreateAccount(ctx, "testpassphrases")
	if accountRepo.RowsError() != nil {
		t.Errorf("Error creating account: %v", accountRepo.RowsError())
		return
	}

	expectedPassPhrase := "testpassphrases"
	if result.PassPhrase != expectedPassPhrase {
		t.Errorf("Expected passPhrase %s, got %s", expectedPassPhrase, result.PassPhrase)
	}
	t.Logf("Account created successfully: %+v", result)
}
