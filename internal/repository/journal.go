package repository

import (
	"PerpusGo/domain"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type JournalRepository struct {
	db *goqu.Database
}

func NewJournal(con *sql.DB) domain.JournalRepository {
	return &JournalRepository{
		db: goqu.New("default", con),
	}
}

// Find implements domain.JournalRepository.
func (j *JournalRepository) Find(ctx context.Context, se domain.JournalSearch) ([]domain.Journal, error) {
	panic("unimplemented")
}

// FindById implements domain.JournalRepository.
func (j *JournalRepository) FindById(ctx context.Context, id string) (domain.Journal, error) {
	panic("unimplemented")
}

// Save implements domain.JournalRepository.
func (j *JournalRepository) Save(ctx context.Context, journal *domain.Journal) error {
	panic("unimplemented")
}

// Update implements domain.JournalRepository.
func (j *JournalRepository) Update(ctx context.Context, journal *domain.Journal) error {
	panic("unimplemented")
}
