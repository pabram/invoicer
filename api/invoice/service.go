package invoice

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Service interface {
	Create(ctx context.Context, companyName string, price int) (Invoice, error)
	Get(ctx context.Context, id int) (Invoice, error)
}

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(repository Repository, logger log.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s *service) Create(ctx context.Context, companyName string, price int) (Invoice, error) {
	logger := log.With(s.logger, "method Create invoice")

	invoice := Invoice{
		CompanyName: companyName,
		Price:       price,
	}
	invoice, err := s.repository.Create(ctx, invoice)
	if err != nil {
		level.Error(logger).Log("error during creating invoice", err)
		return Invoice{}, nil
	}

	return invoice, nil
}

func (s *service) Get(ctx context.Context, id int) (Invoice, error) {
	logger := log.With(s.logger, "method Get invoice")

	invoice, err := s.repository.Get(ctx, id)
	if err != nil {
		level.Error(logger).Log("error during retrieving an invoice")
		return Invoice{}, err
	}

	return invoice, nil
}
