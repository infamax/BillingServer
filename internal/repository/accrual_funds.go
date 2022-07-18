package repository

import (
	"context"
	"os/exec"
)

func (d *db) AccrualFunds(ctx context.Context, userId int, amount float64, reason string) error {
	const queryUpdate = `
		update user_balance
		set balance += $2
		where user_id = $1;
	`

	const queryInsert = `
		insert into accrual
		(user_id, accrual, reason)
		values ($1, $2, $3);
	`

	cmd, err := d.pool.Exec(ctx, queryUpdate, userId, amount)

	if cmd.RowsAffected() == 0 {
		err := exec.ErrNotFound
		return err
	}

	err = d.pool.QueryRow(ctx, queryInsert, userId, amount, reason).Scan()

	return err
}
