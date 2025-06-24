package services

import (
	"context"
	"fmt"

	models "github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	"github.com/meyskens/go-turnstile"
)

type AuthenticationService interface {
	Service
	Register(ctx context.Context, passPhrase string, turnstile string, ip string) string
	Login(ctx context.Context, passPhrase string) string
}

type authenticationService struct {
	*service[repositories.AccountRepository]
	turnStileClient *turnstile.Turnstile
	jwtService      JWTService
}

func NewAuthenticationService(accountRepository repositories.AccountRepository, turnStileClient *turnstile.Turnstile, jwtService JWTService) AuthenticationService {
	return &authenticationService{
		service:         &service[repositories.AccountRepository]{repository: accountRepository},
		turnStileClient: turnStileClient,
		jwtService:      jwtService,
	}
}
func (s *authenticationService) Register(ctx context.Context, passPhrase string, turnstile string, ip string) string {
	turnStileResponse, err := s.turnStileClient.Verify(turnstile, ip)
	fmt.Println(turnstile)
	fmt.Println(turnStileResponse)
	if err != nil {
		s.ThrowsException(&s.exception.Unauthorized, "Turnstile error!, Turnstile Respose :")
		s.ThrowsError(err)
		return ""
	}

	if turnStileResponse.Success {
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
		return ""
	}
}

func (s *authenticationService) Login(ctx context.Context, passPhrase string) string {
	account := s.repository.GetAccountByPassPhrase(ctx, passPhrase)

	if s.repository.IsNoRecord() {
		s.ThrowsException(&s.exception.Unauthorized, "Account not found!")
		return " "
	}
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
}
