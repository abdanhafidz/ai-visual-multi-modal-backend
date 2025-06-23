package tests

import (
	"context"
	"os"
	"testing"

	"github.com/abdanhafidz/ai-visual-multi-modal-backend/config"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/services"
)

func TestAskImage(t *testing.T) {
	config.RunConfig()
	var dummyRepo repositories.Repository
	replicateService := services.NewReplicateService(dummyRepo, config.ReplicateClient, "spuuntries/urna-kp3l:9338a4573a17178b70515c0ef2e613d3b4213e2dc860ef23b3ad6149dacadc1e")
	ctx := context.Background()
	filename := "foto_pacarku.jpg"
	imageFile, err := os.Open("test_data/" + filename)
	if err != nil {
		t.Fatalf("Failed to open image test file: %v", err)
	}
	defer imageFile.Close()
	result := replicateService.AskImage(ctx, imageFile, filename, "What is this image about?")
	if replicateService.Error() != nil {
		t.Fatalf("AskImage failed: %v", replicateService.Error())
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	} else {
		t.Logf("AskImage result: %s", result)
	}

}
