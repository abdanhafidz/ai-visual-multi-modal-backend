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

func (r *replicateService) AskImage(ctx context.Context, imageFile multipart.File, filename, question string) string {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, imageFile); err != nil {
		r.ThrowsError(err)
		return ""
	}

	// Upload file ke replicate
	file, err := r.client.CreateFileFromBuffer(ctx, &buf, &replicate.CreateFileOptions{Filename: filename})
	if err != nil {
		r.ThrowsError(err)
		return ""
	}

	// Input untuk model Moondream
	input := replicate.PredictionInput{
		"image":    file,
		"question": question,
	}

	// Jalankan prediksi, tunggu selesai
	output, err := r.client.Run(ctx, r.model, input, nil)
	if err != nil {
		r.ThrowsError(err)
		return ""
	}

	return output.(string)
}
