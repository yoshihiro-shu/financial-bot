// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: news.sql

package postgresql

import (
	"context"
	"time"
)

const createNews = `-- name: CreateNews :one
INSERT INTO news (
  title, link, published_at
) VALUES (
  $1, $2, $3
)
RETURNING id, title, description, link, thumbnail, published_at
`

type CreateNewsParams struct {
	Title       string    `db:"title"`
	Link        string    `db:"link"`
	PublishedAt time.Time `db:"published_at"`
}

func (q *Queries) CreateNews(ctx context.Context, arg CreateNewsParams) (News, error) {
	row := q.db.QueryRowContext(ctx, createNews, arg.Title, arg.Link, arg.PublishedAt)
	var i News
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Link,
		&i.Thumbnail,
		&i.PublishedAt,
	)
	return i, err
}

const deleteNews = `-- name: DeleteNews :exec
DELETE FROM news
WHERE id = $1
`

func (q *Queries) DeleteNews(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteNews, id)
	return err
}

const getNews = `-- name: GetNews :one
SELECT id, title, description, link, thumbnail, published_at FROM news
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetNews(ctx context.Context, id int32) (News, error) {
	row := q.db.QueryRowContext(ctx, getNews, id)
	var i News
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Link,
		&i.Thumbnail,
		&i.PublishedAt,
	)
	return i, err
}

const listNews = `-- name: ListNews :many
SELECT id, title, description, link, thumbnail, published_at FROM news
ORDER BY published_at
`

func (q *Queries) ListNews(ctx context.Context) ([]News, error) {
	rows, err := q.db.QueryContext(ctx, listNews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []News{}
	for rows.Next() {
		var i News
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Link,
			&i.Thumbnail,
			&i.PublishedAt,
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
