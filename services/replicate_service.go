package services

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	"github.com/replicate/replicate-go"
)

type ReplicateService interface {
	Service
	AskImage(ctx context.Context, imageFile multipart.File, filename, question string) string
}

type replicateService struct {
	*service[repositories.Repository]
	client          *replicate.Client
	deploymentOwner string
	deploymentName  string
}

func NewReplicateService(repo repositories.Repository, replicateClient *replicate.Client, deploymentOwner string, deploymentName string) ReplicateService {
	service := replicateService{
		service:         &service[repositories.Repository]{repository: repo},
		client:          replicateClient,
		deploymentOwner: deploymentOwner,
		deploymentName:  deploymentName, // e.g., "owner/moondream:versionHash"
	}
	return &service
}

func (s *replicateService) AskImage(ctx context.Context, imageFile multipart.File, filename string, question string) string {
	// Buat path file lokal
	filePath := fmt.Sprintf("./images/%s", filename)

	// Simpan file ke direktori ./images
	outFile, err := os.Create(filePath)
	if err != nil {
		s.ThrowsError(err)
		return ""
	}
	defer outFile.Close()

	// Salin data dari multipart.File ke file lokal
	if _, err := io.Copy(outFile, imageFile); err != nil {
		s.ThrowsError(err)
		return ""
	}

	// Gunakan path file untuk membuat file di Replicate
	file, err := s.client.CreateFileFromPath(ctx, filePath, &replicate.CreateFileOptions{Filename: filename})
	if err != nil {
		s.ThrowsError(err)
		return ""
	}

	// Buat input untuk prediksi
	input := replicate.PredictionInput{
		"image":    file,
		"question": question,
	}
	prediction, err := s.client.CreatePredictionWithDeployment(ctx, s.deploymentOwner, s.deploymentName, input, nil, false)

	if err != nil {
		s.ThrowsException(&s.exception.ReplicateConnectionRefused, "Failed to create prediction via replicate service")
		s.ThrowsError(err)
		return ""
	}

	errWait := s.client.Wait(ctx, prediction)

	if errWait != nil {
		s.ThrowsException(&s.exception.ReplicateConnectionRefused, "Time out Waiting Replicate Client")
		s.ThrowsError(err)
		return ""
	}
	rawOutput := prediction.Output
	fmt.Println(rawOutput)
	// Parsing output
	outputSlice, ok := rawOutput.([]interface{})
	var result string

	if ok {
		result = fmt.Sprintf("%v", outputSlice)
		// fmt.Println("Output slice", result)
	} else {
		result = fmt.Sprintf("%v", rawOutput)
		// fmt.Println("Output raw", result)
	}

	return result
}
