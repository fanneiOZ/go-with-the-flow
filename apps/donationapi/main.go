package main

import (
	internalHttp "donationapi/internal/infrastructure/http"
	"infrastructure/pkg/httpserver"
)

func main() {
	server := httpserver.NewHttpServer()
	server.Handle(internalHttp.Router())
	server.Instance()
	server.Start()
}
