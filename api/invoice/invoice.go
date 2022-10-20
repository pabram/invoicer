package invoice

type Invoice struct {
	ID          string `json:"id,omitempty"`
	CompanyName string `json:"company_name"`
	Price       int    `json:"price"`
}
