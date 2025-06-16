package controller

import (
	"mime/multipart"

	models "github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	services "github.com/abdanhafidz/ai-visual-multi-modal-backend/services"
	utils "github.com/abdanhafidz/ai-visual-multi-modal-backend/utils"
	"github.com/gin-gonic/gin"
)

type (
	PredictionController interface {
		Controller
		Predict(ctx *gin.Context)
	}

	predictionController struct {
		*controller[services.PredictionService]
	}
)

func requestImage(ctx *gin.Context, image *multipart.File, imageFilename *string) {
	imageHeader, err := ctx.FormFile("image_file")
	if err != nil {
		utils.ResponseFAIL(ctx, 400, models.Exception{
			BadRequest: true,
			Message:    "Image file is required",
		})
		return
	}
	imageFile, err := imageHeader.Open()
	if err != nil {
		utils.ResponseFAIL(ctx, 400, models.Exception{
			BadRequest: true,
			Message:    "Failed to open image file",
		})
		return
	}
	image = &imageFile
	imageFilename = &imageHeader.Filename
	defer imageFile.Close()
}

func requestAudio(ctx *gin.Context, audio *multipart.File, audioFilename *string) {
	audioHeader, err := ctx.FormFile("audio_file")
	if err != nil {
		utils.ResponseFAIL(ctx, 400, models.Exception{
			BadRequest: true,
			Message:    "Audio file is required",
		})
		return
	}
	audioFile, err := audioHeader.Open()
	if err != nil {
		utils.ResponseFAIL(ctx, 400, models.Exception{
			BadRequest: true,
			Message:    "Failed to open audio file",
		})
		return
	}
	audio = &audioFile
	audioFilename = &audioHeader.Filename
	defer audioFile.Close()
}

func NewPredictionController(predictionService services.PredictionService) PredictionController {
	controller := predictionController{}
	controller.service = predictionService
	return &controller
}
func (c *predictionController) Predict(ctx *gin.Context) {

	var predictionRequest models.PredictionRequest
	requestImage(ctx, &predictionRequest.ImageFile, &predictionRequest.ImageFileName)
	requestAudio(ctx, &predictionRequest.AudioQuestionFile, &predictionRequest.AudioQuestionFilename)
	predictionResult := c.service.Predict(ctx, predictionRequest)

	if c.service.Error() != nil {
		c.Response(ctx, nil)
		return
	}

	ctx.Header("Content-Type", "audio/mpeg")
	ctx.Header("Content-Disposition", "inline; filename=response.mp3")
	ctx.Data(200, "audio/mpeg", predictionResult)

}
