// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: stocks.sql

package postgresql

import (
	"context"
	"time"
)

const createStocks = `-- name: CreateStocks :exec
INSERT INTO stocks (
    symbol, name, open, close, high, low, volume, date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
`

type CreateStocksParams struct {
	Symbol string    `db:"symbol"`
	Name   string    `db:"name"`
	Open   float32   `db:"open"`
	Close  float32   `db:"close"`
	High   float32   `db:"high"`
	Low    float32   `db:"low"`
	Volume int32     `db:"volume"`
	Date   time.Time `db:"date"`
}

func (q *Queries) CreateStocks(ctx context.Context, arg CreateStocksParams) error {
	_, err := q.db.ExecContext(ctx, createStocks,
		arg.Symbol,
		arg.Name,
		arg.Open,
		arg.Close,
		arg.High,
		arg.Low,
		arg.Volume,
		arg.Date,
	)
	return err
}