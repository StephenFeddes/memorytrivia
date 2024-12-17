package getaccount

import (
	"context"

	"github.com/StephenFeddes/memorytrivia/account/internal/domain/entity"
)

type DatabaseAccountGetter interface {
	Execute(context.Context, uint32) (*entity.Account, error)
}

type GetAccount struct{
	db DatabaseAccountGetter
}

func NewGetAccount() *GetAccount {
	return &GetAccount{}
}

func (g *GetAccount) Execute(ctx context.Context, id uint32) (*entity.Account, error) {
	account, err := g.db.Execute(ctx, id); if err != nil {
		return nil, err
	}
	return account, nil
}
