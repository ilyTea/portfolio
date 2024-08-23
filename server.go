package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	rtr := chi.NewRouter()

	rtr.Use(middleware.RequestID)
	rtr.Use(middleware.RealIP)
	rtr.Use(middleware.Logger)

	rtr.Get("/", homeHandler)
	rtr.Get("/liltext", aLittleHandler)


	http.ListenAndServe(":7447", rtr)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {

	ctx := make(map[string]string)
	ctx["Name"] = "Cha"

	tmpt, _ := template.ParseFiles("templates/index.html")

	err := tmpt.Execute(w, ctx)

	if err != nil {
		log.Println("Could not execute template.")
	}
}

func aLittleHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("wow now im here"))
}