package invoice

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger log.Logger
}

func NewService(repository Repository, logger log.Logger) Service {
	return &service{
		repository: repository,
		logger: logger,
	}
}

func (s *service) Create(ctx context.Context, companyName string, price int) (Invoice, error) {
	logger := log.With(s.logger, "method Create invoice")

	uuid, _ := uuid.NewV4()
	id := uuid.String()

	invoice := Invoice{
		ID: id,
		CompanyName: companyName,
		Price: price,
	}
	invoice, err := s.repository.Create(ctx, invoice)
	if err != nil {
		level.Error(logger).Log("error during creating invoice", err)
		return Invoice{}, nil
	}

	logger.Log("succesfully created an invoice")

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