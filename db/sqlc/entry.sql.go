package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id,
                      to_account_id,
					  user_id,
                      amount,
					  created_at) VALUES ($1, $2, $3)

RETURNING id, from_account_id, to_account_id, user_id, amount, created_at
`

const createEntry = `--- name: CreateEntry :one
BEGIN TRANSACTION;
UPDATE entries SET balance = $3
WHERE id = $1

UPDATE entries SET balance = $3
WHERE id = $2

COMMIT TRANSACTION;
`

type CreateEntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type ListEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)

	var t Transfer
	err := row.Scan(
		&t.ID,
		&t.FromAccountID,
		&t.ToAccountID,
		&t.UserID,
		&t.Amount,
		&t.CreatedAt,
	)

	return t, err
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountID, arg.Amount)
	var e Entry

	err := row.Scan(
		&e.ID,
		&e.AccountID,
		&e.Amount,
		&e.CreatedAt,
	)

	return e, err
}

const getEntry = `-- name: GetEntry: one
SELECT id, account_id, amount, created_at FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)

	var e Entry
	err := row.Scan(
		&e.ID,
		&e.AccountID,
		&e.Amount,
		&e.CreatedAt,
	)

	return e, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, account_id, amount, created_at FROM entries
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []Entry
	for rows.Next() {
		var e Entry
		if err := rows.Scan(
			&e.ID,
			&e.AccountID,
			&e.Amount,
			&e.CreatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, e)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}


