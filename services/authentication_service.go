package services

import (
	"context"

	"github.com/9ssi7/turnstile"
	models "github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
)

type AuthenticationService interface {
	Service
	Register(ctx context.Context, passPhrase string, turnstile string, ip string) string
	Login(ctx context.Context, passPhrase string) string
}

type authenticationService struct {
	*service[repositories.AccountRepository]
	turnStileClient turnstile.Service
	jwtService      JWTService
}

func NewAuthenticationService(accountRepository repositories.AccountRepository, turnStileClient turnstile.Service, jwtService JWTService) AuthenticationService {
	return &authenticationService{
		service:         &service[repositories.AccountRepository]{repository: accountRepository},
		turnStileClient: turnStileClient,
		jwtService:      jwtService,
	}
}
func (s *authenticationService) Register(ctx context.Context, passPhrase string, turnstile string, ip string) string {
	verifiedTurnStile, err := s.turnStileClient.Verify(ctx, turnstile, ip)

	if err != nil {
		s.ThrowsException(&s.exception.Unauthorized, "Turnstile error!")
		return ""
	}

	if verifiedTurnStile {
		account := s.repository.CreateAccount(ctx, passPhrase)
		if s.ThrowsRepoException() {
			return ""
		}
		token := s.jwtService.GenerateToken(ctx, models.JWTCustomClaims{IdUser: account.ID})
		if s.jwtService.Error() != nil {
			s.ThrowsException(&s.exception.Unauthorized, "JWTService Error")
			s.ThrowsError(s.jwtService.Error())
			return ""
		}

		return token
	} else {
		s.ThrowsException(&s.exception.Unauthorized, "Invalid turnstile payload!")
	}

	return ""
}

func (s *authenticationService) Login(ctx context.Context, passPhrase string) string {
	account := s.repository.GetAccountByPassPhrase(ctx, passPhrase)
	if s.ThrowsRepoException() {
		return ""
	}

	if s.repository.IsNoRecord() {
		s.ThrowsException(&s.exception.Unauthorized, "Account not found!")
		return " "
	}

	token := s.jwtService.GenerateToken(ctx, models.JWTCustomClaims{IdUser: account.ID})
	if s.jwtService.Error() != nil {
		s.ThrowsException(&s.exception.Unauthorized, "JWTService Error")
		s.ThrowsError(s.jwtService.Error())
		return ""
	}
	return token
}
