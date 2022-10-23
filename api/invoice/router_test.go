package invoice

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetInvoiceHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := new(MockService)
	invoiceHandler := newInvoiceHandler(service)
	r := SetupRouter()
	r.GET("/:id", invoiceHandler.getInvoiceHandler)
	mockInvoiceResponse := Invoice{
		CompanyName: "Pioterpol",
		Price:       2137,
	}
	service.On("Get", mock.Anything, 1).Return(mockInvoiceResponse, nil)

	req, err := http.NewRequest("GET", "/1", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := `{"company_name":"Pioterpol","price":2137}`

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectedResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetInvoiceHandlerInvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := new(MockService)
	invoiceHandler := newInvoiceHandler(service)
	r := SetupRouter()
	r.GET("/:id", invoiceHandler.getInvoiceHandler)
	service.On("Get", mock.Anything, 1).Return(Invoice{}, nil)

	req, err := http.NewRequest("GET", "/1p", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := `"invalid id parameter"`

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectedResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateInvoiceHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := new(MockService)
	invoiceHandler := newInvoiceHandler(service)
	r := SetupRouter()
	r.GET("/", invoiceHandler.createInvoiceHandler)
	mockInvoiceResponse := Invoice{
		CompanyName: "Pioterpol",
		Price:       2137,
	}
	service.On("Create", mock.Anything, "Pioterpol", 2137).Return(mockInvoiceResponse, nil)

	invoice := Invoice{
		CompanyName: "Pioterpol",
		Price:       2137,
	}
	payload, err := json.Marshal(invoice)
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := `{"company_name":"Pioterpol","price":2137}`

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectedResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateInvoiceHandlerInvalid(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := new(MockService)
	invoiceHandler := newInvoiceHandler(service)
	r := SetupRouter()
	r.GET("/", invoiceHandler.createInvoiceHandler)
	mockInvoiceResponse := Invoice{
		CompanyName: "Pioterpol",
		Price:       2137,
	}
	service.On("Create", mock.Anything, "Pioterpol", 2137).Return(mockInvoiceResponse, nil)

	type InvalidInvoice struct {
		wrongFielName string
	}
	invoice := InvalidInvoice{
		wrongFielName: "Pioterpol",
	}
	payload, err := json.Marshal(invoice)
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := `"invalid request"`

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectedResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
