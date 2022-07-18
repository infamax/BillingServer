package repository

import (
	"context"
	"github.com/infamax/billing_server/internal/models"
)

type Repo interface {
	GetUserBalance(ctx context.Context, userId int) (*models.UserBalance, error)
	AddUser(ctx context.Context, userId int) error
	AccrualFunds(ctx context.Context, userId int, amount float64, reason string) error
	WithdrawMoney(ctx context.Context, userId int, amount float64, reason string) error
}
