package invoice

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CreateRequest struct {
	CompanyName string `json:"company_name"`
	Price int `json:"price"`
}

type GetRequest struct {
	ID string `json:"id"`
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetRequest
	vars := mux.Vars(r)

	req = GetRequest{
		ID: vars["id"],
	}

	return req, nil
}