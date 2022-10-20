package invoice

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetInvoice(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var logger log.Logger
	repository := new(MockRepository)
	service := NewService(repository, logger)
	ctx := context.TODO()

	mockInvoiceResponse := Invoice{
		CompanyName: "Pioterpol",
		Price:       2137,
	}

	repository.On("Get", mock.Anything, 1).Return(mockInvoiceResponse, nil)

	invoice, err := service.Get(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, Invoice{CompanyName: "Pioterpol", Price: 2137}, invoice)
}

func TestCreateInvoice(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var logger log.Logger
	repository := new(MockRepository)
	service := NewService(repository, logger)
	ctx := context.TODO()

	mockInvoiceResponse := Invoice{
		CompanyName: "Pioterpol",
		Price:       2137,
	}

	repository.On("Create", mock.Anything, Invoice{CompanyName: "Pioterpol", Price: 2137}).Return(mockInvoiceResponse, nil)

	invoice, err := service.Create(ctx, "Pioterpol", 2137)
	assert.NoError(t, err)
	assert.Equal(t, Invoice{CompanyName: "Pioterpol", Price: 2137}, invoice)
}
