package main

import (
	internalHttp "donationapi/internal/infra/http"
	"sharedinfra/httpserver"
)

func main() {
	server := httpserver.NewHttpServer()
	server.Handle(internalHttp.Router())
	server.Instance()
	server.Start()
}
