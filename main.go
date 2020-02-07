package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pinguinens/mail-service/service"
	"net/http"
)

func main() {
	var svc service.ImapService
	svc = service.ImapService{}

	listHandler := httptransport.NewServer(
		makeListEndpoint(svc),
		decodeListRequest,
		encodeResponse,
		httptransport.ServerAfter(
			httptransport.SetContentType("text/json"),
		),
	)

	http.Handle("/list", listHandler)
	http.ListenAndServe(":8080", nil)
}
