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

func (s *HttpServer) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			return
		}
	})

	http.Handle("/", mux)
	err := s.server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
