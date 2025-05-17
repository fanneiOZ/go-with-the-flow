package httpserver

import "net/http"

type HttpServer struct {
	server *http.Server
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		server: &http.Server{
			Addr: ":3001",
		},
	}
}

func (s *HttpServer) Instance() *http.Server {
	return s.server
}

func (s *HttpServer) Handle(handler http.Handler) *HttpServer {
	http.Handle("/", handler)

	return s
}

func (s *HttpServer) Start() {
	err := s.server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
