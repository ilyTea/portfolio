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
	rtr.Use(middleware.Recoverer)

	rtr.Get("/", landingHandler)
	rtr.Get("/projects", teaProjectHandler)
	rtr.Get("/profile", teaProfileHandler)


	http.ListenAndServe(":7447", rtr)
}

func landingHandler(w http.ResponseWriter, req *http.Request) {
	ctx := make(map[string]string)
	ctx["Name"] = "Cha"

	tmpt, _ := template.ParseFiles("templates/index.html")

	err := tmpt.Execute(w, ctx)

	if err != nil {
		log.Println("Could not execute template.")
	}
}

func teaProjectHandler(w http.ResponseWriter, req *http.Request) {
	ctx := make(map[string]string)
	ctx["Name"] = "Cha"

	tmpt, _ := template.ParseFiles("templates/index.html")

	err := tmpt.Execute(w, ctx)
	if err != nil {
		log.Println("Could not execute template.")
	}
}

func teaProfileHandler(w http.ResponseWriter, req *http.Request) {
	ctx := make(map[string]string)
	ctx["Name"] = "Cha"

	tmpt, _ := template.ParseFiles("templates/profile.html")

	err := tmpt.Execute(w, ctx)
	if err != nil {
		log.Println("Could not execute template.")
	}
}