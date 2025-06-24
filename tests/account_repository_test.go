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
	passPhrase := "testpassphrases"
	accountRepo.DeleteByPassPhrase(ctx, passPhrase)
	result := accountRepo.CreateAccount(ctx, passPhrase)
	if accountRepo.IsNoRecord() {
		t.Errorf("No Record Account: %v", accountRepo.RowsError())
		return
	}
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

func TestGetAccountByPassPhraseRepository(t *testing.T) {
	config.RunConfig()
	var ctx context.Context
	ctx = context.Background()
	t.Log("DB Ptr :", config.DB)
	accountRepo := repositories.NewAccountRepository(config.DB)
	t.Log("Repo Ptr :", accountRepo)

	passPhrase := "testpassphrases"

	result := accountRepo.GetAccountByPassPhrase(ctx, passPhrase)
	if accountRepo.IsNoRecord() {
		t.Errorf("No Record Account: %v", accountRepo.IsNoRecord())
		return
	}
	if accountRepo.RowsError() != nil {
		t.Errorf("Error getting account: %v", accountRepo.RowsError())
		return
	}

	if result.PassPhrase != passPhrase {
		t.Errorf("Expected passPhrase %s, got %s", passPhrase, result.PassPhrase)
	} else {
		t.Logf("Account retrieved successfully: %+v", result)
	}
}
