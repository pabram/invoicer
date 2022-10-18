package invoice

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repository{
		db: db,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, invoice Invoice) (Invoice, error) {
	// TODO
	return Invoice{}, nil
}

func (r *repository) Get(ctx context.Context, ID int) (Invoice, error) {
	// TODO
	return Invoice{}, nil
}