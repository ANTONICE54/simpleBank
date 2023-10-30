package db

import "context"

type CreateEntryParams struct {
	AccountId int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	query := "INSERT INTO entries (account_id, amount) VALUES ($1, $2) RETURNING id, account_id, amount, created_at"
	row := q.db.QueryRowContext(ctx, query, arg.AccountId, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountId,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err

}

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	query := "SELECT id, account_id, amount, created_at FROM entries WHERE id=$1 LIMIT 1"
	row := q.db.QueryRowContext(ctx, query, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountId,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err

}

type ListEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int64 `json:"limit"`
	Offset    int64 `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	query := "SELECT id, account_id, amount, created_at FROM entries WHERE account_id=$3 LIMIT $1 OFFSET $2"
	rows, err := q.db.QueryContext(ctx, query, arg.Limit, arg.Offset, arg.AccountID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountId,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
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

	return items, nil

}
