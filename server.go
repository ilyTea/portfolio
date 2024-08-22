package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	rtr := chi.NewRouter()

	rtr.Use(middleware.RequestID)
	rtr.Use(middleware.RealIP)
	rtr.Use(middleware.Logger)
	rtr.Use(middleware.Recoverer)

  	// Set a timeout value on the request context (ctx), that will signal
  	// through ctx.Done() that the request has timed out and further
  	// processing should be stopped.
  	rtr.Use(middleware.Timeout(60 * time.Second))

	rtr.Get("/dog", func(w http.ResponseWriter, req *http.Request) {
		
	})





	http.ListenAndServe(":7447", rtr)
}