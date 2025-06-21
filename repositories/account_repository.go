package repositories

import (
	"context"

	"github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, fingerPrint string) (res models.Account)
	GetAccountByFingerPrint(ctx context.Context, fingePrint string) (res models.Account)
}

type accountRepository struct {
	*repository[models.Account]
}

func (r *accountRepository) CreateAccount(ctx context.Context, fingerPrint string) (res models.Account) {
	r.entity.Fingerprint = fingerPrint
	r.Create(ctx)
	return r.entity
}

func (r *accountRepository) GetAccountByFingerPrint(ctx context.Context, fingerPrint string) (res models.Account) {
	r.entity.Fingerprint = fingerPrint
	r.Where(ctx)
	r.Find(ctx, res)
	return res
}
