package db

import "context"

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	query := "INSERT INTO accounts (owner, balance, currency) VALUES ($1, $2, $3) RETURNING id, owner, balance, currency, created_at"
	row := q.db.QueryRowContext(ctx, query, arg.Owner, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err

}

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	query := "SELECT id, owner, balance, currency, created_at FROM accounts WHERE id=$1 LIMIT 1"
	row := q.db.QueryRowContext(ctx, query, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err

}

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int64) (Account, error) {
	query := "SELECT id, owner, balance, currency, created_at FROM accounts WHERE id=$1 LIMIT 1 FOR NO KEY UPDATE"
	row := q.db.QueryRowContext(ctx, query, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err

}

type ListAccountsParams struct {
	Owner  string `json:"owner"`
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	query := "SELECT id, owner, balance, currency, created_at FROM accounts WHERE owner = $1 LIMIT $2 OFFSET $3"
	rows, err := q.db.QueryContext(ctx, query, arg.Owner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err

	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		err = rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err

	}
	if err := rows.Err(); err != nil {
		return nil, err

	}

	return items, err

}

type UpdateAccountParams struct {
	Balance int64 `json:"balance"`
	ID      int64 `json:"id"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	query := "UPDATE accounts SET balance=$2 WHERE id=$1 RETURNING id, owner, balance, currency, created_at"
	row := q.db.QueryRowContext(ctx, query, arg.ID, arg.Balance)

	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err

}

type AddAccountBalanceParams struct {
	Amount int64 `json:"balance"`
	ID     int64 `json:"id"`
}

func (q *Queries) AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error) {
	query := "UPDATE accounts SET balance= balance + $2 WHERE id=$1 RETURNING id, owner, balance, currency, created_at"
	row := q.db.QueryRowContext(ctx, query, arg.ID, arg.Amount)

	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err

}

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	query := "DELETE FROM accounts WHERE id=$1"
	_, err := q.db.ExecContext(ctx, query, id)

	return err

}
