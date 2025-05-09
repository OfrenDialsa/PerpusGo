package repository

import (
	"context"
	"database/sql"
	"libraryRestAPI/domain"

	"github.com/doug-martin/goqu"
)

type customerRepository struct {
	db *goqu.Database
}

func NewCustomer(con *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		db: goqu.New("default", con),
	}
}

// FindAll implements domain.CustomerRepository.
func (cr *customerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	panic("unimplemented")
}

// FindByID implements domain.CustomerRepository.
func (cr *customerRepository) FindByID(ctx context.Context, id string) (domain.Customer, error) {
	panic("unimplemented")
}

// Save implements domain.CustomerRepository.
func (cr *customerRepository) Save(ctx context.Context, c *domain.Customer) error {
	panic("unimplemented")
}

// Update implements domain.CustomerRepository.
func (cr *customerRepository) Update(ctx context.Context, c *domain.Customer) error {
	panic("unimplemented")
}

// Delete implements domain.CustomerRepository.
func (cr *customerRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}
