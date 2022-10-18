package invoice

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"strconv"
)

type Endpoints struct {
	Create endpoint.Endpoint
	Get    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		invoice, err := s.Create(ctx, req.CompanyName, req.Price)
		return CreateResponse{ID: invoice.ID}, err
	}
}

func makeGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		id, err := strconv.Atoi(req.ID)
		invoice, err := s.Get(ctx, id)
		return GetResponse{ID: invoice.ID}, err
	}
}
