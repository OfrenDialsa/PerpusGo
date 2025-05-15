package domain

import (
	"context"
	"database/sql"
)

type Book struct {
	Id          string       `db:"id"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	Isbn        int          `db:"isbn"`
	Created_at  sql.NullTime `db:"created_at"`
	Updated_at  sql.NullTime `db:"updated_at"`
	Deleted_at  sql.NullTime `db:"deleted_at"`
}

type BookRepository interface {
	FindAll(ctx context.Context) ([]Book, error)
	FindById(ctx context.Context, id string) (Book, error)
	Save(ctx context.Context, book *Book) error
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, id string) error
}
