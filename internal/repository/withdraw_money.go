package repository

import (
	"context"
	"errors"
	"os/exec"
)

func (d *db) WithdrawMoney(ctx context.Context, userId int, amount float64, reason string) error {
	userBalance, err := d.GetUserBalance(ctx, userId)

	if err != nil {
		return err
	}

	if userBalance.Balance < amount {
		return errors.New("withdrawing money more than balance")
	}

	const queryUpdate = `
		update user_balance
		set balance = balance - amount
		where user_id = $1;
	`

	const queryInsert = `
		insert into payment
		(user_id, payment, reason)
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
