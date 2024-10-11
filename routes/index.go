package routes

import (
	"log"
	"net/http"
	"trandung2k1/server/middlewares"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("userId")
		w.Write([]byte("UserId : " + userId))

	})
	router.HandleFunc("PUT /api/v1/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("userId")
		w.Write([]byte("UserId : " + userId))

	})

	router.HandleFunc("GET /api/v1/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("userId")
		w.Write([]byte("UserId : " + userId))

	})

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	// C1
	// server := http.Server{Addr: s.addr,Handler: RequireAuthMiddleware(RequestLoggerMiddleware(router))}

	// C2
	// server := http.Server{Addr: s.addr, Handler: MiddlewareChain(RequestLoggerMiddleware, RequireAuthMiddleware)(router)}

	// C3
	middlewareChain := middlewares.MiddlewareChain(middlewares.RequestLoggerMiddleware, middlewares.RequireAuthMiddleware)
	server := http.Server{Addr: s.addr, Handler: middlewareChain(router)}
	log.Printf("Server has started: %s", s.addr)
	return server.ListenAndServe()
}
