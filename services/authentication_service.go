package services

import "context"

type AuthenticationService interface {
	Register(ctx context.Context, fingerPrint string)
	Login(ctx context.Context, fingerPrint string)
}

type authenticationService interface {
}
