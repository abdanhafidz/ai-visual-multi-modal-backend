package services

import (
	"bytes"
	"context"
	"fmt"
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

	file, err := s.client.CreateFileFromBuffer(ctx, &buf, &replicate.CreateFileOptions{Filename: filename})
	if err != nil {
		s.ThrowsError(err)
		return ""
	}

	input := replicate.PredictionInput{
		"image":    file,
		"question": question,
	}
	rawOutput, err := s.client.Run(ctx, s.model, input, nil)
	if err != nil {
		s.ThrowsError(err)
		return ""
	}

	outputSlice, _ := rawOutput.([]interface{})
	result := fmt.Sprintf("%v", outputSlice)
	fmt.Println("Output slice", result)

	// if !ok {
	// 	s.ThrowsError(errors.New("failed to parse output as string"))
	// 	return ""
	// }

	return result
}
