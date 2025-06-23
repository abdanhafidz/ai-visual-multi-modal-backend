package services

import (
	"context"

	models "github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	repositories "github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type JWTService interface {
	Service
	GenerateToken(ctx context.Context, payload models.JWTCustomClaims) string
	ValidateToken(ctx context.Context, tokenStr string) *models.JWTCustomClaims
}

type jwtService struct {
	*service[repositories.AccountRepository]
	secretKey string
}

func NewJWTService(repo repositories.AccountRepository, secretKey string) JWTService {
	return &jwtService{
		service:   &service[repositories.AccountRepository]{repository: repo},
		secretKey: secretKey,
	}
}
func (s *jwtService) GenerateToken(ctx context.Context, payload models.JWTCustomClaims) string {
	claims := jwt.MapClaims{
		"user_id": payload.IdUser,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		s.ThrowsException(&s.exception.Unauthorized, "Failed to generate JWT token!")
		s.ThrowsError(err)
		return ""
	}
	return tokenStr
}

func (s *jwtService) ValidateToken(ctx context.Context, tokenStr string) *models.JWTCustomClaims {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s.ThrowsException(&s.exception.Unauthorized, "Unexpected signing method")
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})

	if err != nil || !token.Valid {
		s.ThrowsException(&s.exception.Unauthorized, "Invalid token!")
		s.ThrowsError(err)
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		s.ThrowsException(&s.exception.Unauthorized, "Invalid token claims")
		return nil
	}

	return &models.JWTCustomClaims{
		IdUser: claims["user_id"].(uuid.UUID),
	}
}
