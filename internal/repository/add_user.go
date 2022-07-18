package repository

import "context"

func (d *db) AddUser(ctx context.Context, userId int) error {
	const query = `
		insert into user_balance
		(user_id, balance)
		values ($1, $2);
	`

	err := d.pool.QueryRow(ctx, query, userId, 0).Scan()

	if err != nil {
		return err
	}

	return nil
}
