package main

import (
	"context"
	"encoding/json"
	"goimap/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeListEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(listRequest)
		v := svc.List(req.URL, req.Username, req.Password)

		return listResponse{v}, nil
	}
}

func decodeListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request listRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	request.Username, request.Password, _ = r.BasicAuth()

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type listRequest struct {
	URL      string `json:"url"`
	Username string
	Password string
}

type listResponse struct {
	Mailboxes []string `json:"mailboxes"`
}
