package repository

import (
	"context"
	"errors"
	"github.com/infamax/billing_server/internal/models"
	"github.com/jackc/pgx/v4"
)

func (d *db) GetUserBalance(ctx context.Context, userId int) (*models.UserBalance, error) {
	const query = `
		select user_id, balance
		from user_balance
		where user_id = $1;
	`

	var userBalance models.UserBalance
	err := d.pool.QueryRow(ctx, query, userId).Scan(
		&userBalance.ID,
		&userBalance.Balance)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	return &userBalance, nil
}
