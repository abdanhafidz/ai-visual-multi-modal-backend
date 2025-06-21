package models

type LoginRequest struct {
	FingerPrintToken string `json:"email" binding:"required"`
}
