package invoice

import "context"

type Service interface {
	Create(ctx context.Context, companyName string, price int) (Invoice, error)
	Get(ctx context.Context, id int) (Invoice, error)
}
