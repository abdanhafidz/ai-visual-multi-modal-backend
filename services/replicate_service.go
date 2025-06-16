package services

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"

	"github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	"github.com/replicate/replicate-go"
)

type ReplicateService interface {
	Service
	AskImage(ctx context.Context, imageFile multipart.File, filename, question string) string
}

type replicateService struct {
	*service[repositories.Repository]
	client *replicate.Client
	model  string
}

func NewReplicateService(repo repositories.Repository, replicateClient *replicate.Client, model string) ReplicateService {
	service := replicateService{
		service: &service[repositories.Repository]{repository: repo},
		client:  replicateClient,
		model:   model, // e.g., "owner/moondream:versionHash"
	}
	return &service
}

func (s *replicateService) AskImage(ctx context.Context, imageFile multipart.File, filename, question string) string {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, imageFile); err != nil {

		s.ThrowsError(err)
		return ""
	}

	// Upload file ke replicate
	file, err := s.client.CreateFileFromBuffer(ctx, &buf, &replicate.CreateFileOptions{Filename: filename})
	if err != nil {
		s.ThrowsError(err)
		return ""
	}

	// Input untuk model Moondream
	input := replicate.PredictionInput{
		"image":    file,
		"question": question,
	}

	// Jalankan prediksi, tunggu selesai
	output, err := s.client.Run(ctx, s.model, input, nil)
	if err != nil {
		s.ThrowsError(err)
		return ""
	}

	return output.(string)
}
