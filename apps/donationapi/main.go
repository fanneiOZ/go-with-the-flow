package main

import (
	donationApi "donationapi/infra/http"
	"sharedinfra/httpserver"
)

func main() {
	server := httpserver.NewHttpServer()
	server.Handle(donationApi.Router())
	server.Instance()
	server.Start()
}
