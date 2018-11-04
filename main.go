package main

import (
	"github.com/go-chi/chi"
	"github.com/khanadnanxyz/konta/api"
	"github.com/khanadnanxyz/konta/config"
	"github.com/khanadnanxyz/konta/conn"
	_ "github.com/khanadnanxyz/konta/helper"
	"net/http"
)

func main() {
	router := routers()
	http.ListenAndServe(":8080", router)
}

func init()  {
	config.LoadConfig()
	db.Connect()
}

func routers() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/questions", func(r chi.Router) {
		r.Get("/", api.AllQuestions)
		r.Get("/{id}", api.DetailQuestion)
		r.Post("/", api.CreateQuestion)
		r.Put("/{id}", api.UpdateQuestion)
		r.Delete("/{id}", api.DeleteQuestion)
	})

	r.Route("/answers", func(r chi.Router) {
		r.Get("/", api.AllAnswers)
		r.Get("/{id}", api.DetailAnswer)
		r.Post("/", api.CreateAnswer)
		r.Put("/{id}", api.UpdateAnswer)
		r.Delete("/{id}", api.DeleteAnswer)
	})

	return r
}
