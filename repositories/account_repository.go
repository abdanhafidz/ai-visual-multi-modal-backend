package repositories

import (
	"context"

	"github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Repository
	CreateAccount(ctx context.Context, passPhrase string) (res models.Account)
	GetAccountByPassPhrase(ctx context.Context, passPhrase string) (res models.Account)
}

type accountRepository struct {
	*repository[models.Account]
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		repository: &repository[models.Account]{
			entity:      models.Account{},
			transaction: db,
		},
	}
}

func (r *accountRepository) CreateAccount(ctx context.Context, passPhrase string) (res models.Account) {
	r.entity.PassPhrase = passPhrase
	r.Create(ctx)
	return r.entity
}

func (r *accountRepository) GetAccountByPassPhrase(ctx context.Context, passPhrase string) (res models.Account) {
	r.entity.PassPhrase = passPhrase
	r.Where(ctx)
	r.Find(ctx, &res)
	return res
}
