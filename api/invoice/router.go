package invoice

import (
	"context"

	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type invoiceHandler struct {
	service Service
}

func newInvoiceHandler(service Service) invoiceHandler {
	return invoiceHandler{
		service: service,
	}
}

func (s *invoiceHandler) getInvoiceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, "invalid id parameter")
		return
	}
	invoice, err := s.service.Get(c, id)
	if err != nil {
		c.JSON(500, "error during invoice retrieval")
		return
	}

	c.JSON(200, invoice)
}

func (s *invoiceHandler) createInvoiceHandler(c *gin.Context) {
	var apiInvoice Invoice
	err := c.BindJSON(&apiInvoice)
	fmt.Println("apiInvoice: ", apiInvoice)
	if err != nil {
		c.JSON(400, "invalid request")
		return
	}
	invoice, err := s.service.Create(c, apiInvoice.CompanyName, apiInvoice.Price)
	if err != nil {
		c.JSON(500, "error during invoice creation")
		return
	}

	c.JSON(200, invoice)
}

func NewRouter(ctx context.Context, service Service) *gin.Engine {
	invoiceHandler := newInvoiceHandler(service)

	router := gin.Default()

	// Invoices
	router.GET("/invoices/:id", invoiceHandler.getInvoiceHandler)
	router.POST("/invoices", invoiceHandler.createInvoiceHandler)

	return router
}
