package repository

import (
	"PerpusGo/domain"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type chargerepository struct {
	db *goqu.Database
}

func NewCharge(con *sql.DB) domain.Chargerepository {
	return &chargerepository{
		db: goqu.New("default", con),
	}
}

// Save implements domain.Chargerepository.
func (c *chargerepository) Save(ctx context.Context, charge *domain.Charge) error {
	executor := c.db.Insert("charges").Rows(charge).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
