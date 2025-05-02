package test

import "infrastructure/pkg/httpserver"

func main() {
	server := httpserver.NewHttpServer()
	server.Start()
}
