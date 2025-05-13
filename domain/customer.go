package domain

import (
	"PerpusGo/dto"
	"context"
	"database/sql"
)

type Customer struct {
	ID         string       `db:"id"`
	Code       string       `db:"code"`
	Name       string       `db:"name"`
	Created_at sql.NullTime `db:"created_at"`
	Updated_at sql.NullTime `db:"updated_at"`
	Deleted_at sql.NullTime `db:"deleted_at"`
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindByID(ctx context.Context, id string) (Customer, error)
	Save(ctx context.Context, c *Customer) error
	Update(ctx context.Context, c *Customer) error
	Delete(ctx context.Context, id string) error
}

type CustomerService interface {
	Index(ctx context.Context) ([]dto.CustomerData, error)
	Create(ctx context.Context, req dto.CreateCustomerRequest) error
}
