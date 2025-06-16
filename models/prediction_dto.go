package models

import "mime/multipart"

type PredictionRequest struct {
	ImageFile             multipart.File
	ImageFileName         string
	AudioQuestionFile     multipart.File
	AudioQuestionFilename string
}
