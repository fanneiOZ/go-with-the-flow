package main

import "infrastructure/pkg/httpserver"

func main() {
	server := httpserver.NewHttpServer()
	server.Start()
}
