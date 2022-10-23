package invoice

type Invoice struct {
	ID          string `json:"id,omitempty"`
	CompanyName string `json:"company_name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
}
