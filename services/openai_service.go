package services

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/sashabaranov/go-openai"

	repositories "github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
)

type (
	openAIService struct {
		*service[repositories.Repository]
		client *openai.Client
	}

	OpenAIService interface {
		Service
		SpeechToText(ctx context.Context, audioFile multipart.File, filename string) string
		TextToSpeech(ctx context.Context, text string) []byte
	}
)

func NewOpenAIService(repo repositories.Repository, openAIClient *openai.Client) OpenAIService {
	return &openAIService{
		service: &service[repositories.Repository]{
			repository: repo,
		},
		client: openAIClient,
	}
}

func (s *openAIService) SpeechToText(ctx context.Context, audioFile multipart.File, filename string) string {

	audioDir := "audio"

	if err := os.MkdirAll(audioDir, os.ModePerm); err != nil {
		s.ThrowsException(&s.exception.InternalServerError, "Failed to create directory!")
		s.ThrowsError(err)
		return "failed to create directory!"
	}

	savedPath := filepath.Join(audioDir, filepath.Base(filename))
	outFile, err := os.Create(savedPath)

	if err != nil {
		s.ThrowsException(&s.exception.AudioFileError, "Failed to save audio!")
		s.ThrowsError(err)
		return "Failed to save audio!"
	}

	defer outFile.Close()

	if _, err := io.Copy(outFile, audioFile); err != nil {
		s.ThrowsException(&s.exception.AudioFileError, "Failed to save audio!")
		s.ThrowsError(err)
		return "Failed to save audio!"
	}

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: savedPath,
	}

	resp, err := s.client.CreateTranscription(ctx, req)
	if err != nil {
		s.ThrowsException(&s.exception.FailedTranscripting, "Failed to create transcription!")
		s.ThrowsError(err)
		return "Failed to create transcription!"
	}

	// Privacy Consideration bro!
	if err := os.Remove(savedPath); err != nil {
		s.ThrowsError(err)
	}

	return resp.Text
}

func (s *openAIService) TextToSpeech(ctx context.Context, text string) []byte {
	req := openai.CreateSpeechRequest{
		Model:          openai.TTSModel1,
		Input:          text,
		Voice:          openai.VoiceNova,
		ResponseFormat: openai.SpeechResponseFormatMp3,
	}

	audioResp, err := s.client.CreateSpeech(ctx, req)
	if err != nil {
		s.ThrowsException(&s.exception.FailedGenerateAudio, "Failed to generate speech audio!")
		s.ThrowsError(err)
		return nil
	}

	audioData, err := io.ReadAll(audioResp)
	if err != nil {
		s.ThrowsException(&s.exception.AudioFileError, "Failed to read audio response!")
		s.ThrowsError(err)
		return nil
	}
	return audioData
}
