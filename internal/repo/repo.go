package repo

import (
	"context"
	"github.com/infamax/billing_server/internal/models"
)

type Repo interface {
	TopUp(ctx context.Context, user models.UserBalance, reason string) error
	GetBalance(ctx context.Context, ID int) (*models.UserBalance, error)
	Transfer(ctx context.Context, src models.UserBalance, dst int) error
}
