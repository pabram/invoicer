package invoice

import "context"

type Invoice struct {
	ID          string `json:"id,omitempty"`
	CompanyName string `json:"company_name"`
	Price       int    `json:"price"`
}

type Repository interface {
	Get(ctx context.Context, ID int) (Invoice, error)
	Create(ctx context.Context, invoice Invoice) (Invoice, error)
}
