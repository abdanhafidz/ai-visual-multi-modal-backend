package tests

import (
	"context"
	"os"
	"testing"

	"github.com/abdanhafidz/ai-visual-multi-modal-backend/config"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/services"
)

func TestSpeechToText(t *testing.T) {
	config.RunConfig()
	// Inisialisasi dependency
	var dummyRepo repositories.Repository
	openAIService := services.NewOpenAIService(dummyRepo, config.OpenAIClient)

	ctx := context.Background()
	filename := "test_audio.mp3"
	// Buka file audio dummy untuk pengujian
	audioFile, err := os.Open("test_data/" + filename)
	if err != nil {
		t.Fatalf("Failed to open audio test file: %v", err)
	}
	defer audioFile.Close()

	result := openAIService.SpeechToText(ctx, audioFile, filename)
	if openAIService.Error() != nil {
		t.Fatalf("Speech To Text failed: %v", openAIService.Error())
	}
	if result == "" {
		t.Errorf("Expected transcription result, got empty string")
	} else {
		t.Logf("Transcription result: %s", result)
	}
}

func TestTextToSpeech(t *testing.T) {
	config.RunConfig()
	// Inisialisasi dependency
	var dummyRepo repositories.Repository
	openAIService := services.NewOpenAIService(dummyRepo, config.OpenAIClient)

	ctx := context.Background()
	text := "Halo, ini adalah pengujian Text to Speech."

	audioBytes := openAIService.TextToSpeech(ctx, text)

	if openAIService.Error() != nil {
		t.Fatalf("TextToSpeech failed: %v", openAIService.Error())
	}

	if len(audioBytes) == 0 {
		t.Errorf("Expected non-empty audio bytes")
	}

	// (Opsional) Simpan hasil untuk verifikasi manual
	if err := os.WriteFile("test_data/output_test.mp3", audioBytes, 0644); err != nil {
		t.Errorf("Gagal menyimpan hasil audio: %v", err)
	}
}
