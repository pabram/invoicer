package invoice

import (
	"context"
	"encoding/json"
	"net/http"
)

type CreateResponse struct {
	ID string `json:"id"`
}

type GetResponse struct {
	ID string `json:"id"`
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}