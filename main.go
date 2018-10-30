package main

import (
	"github.com/go-chi/chi"
	"github.com/khanadnanxyz/konta/api"
	"github.com/khanadnanxyz/konta/conn"
	_ "github.com/khanadnanxyz/konta/helper"
	"net/http"
)

func main() {
	router := routers()
	http.ListenAndServe(":8080", router)
}

func init()  {
	db.Connect()
}

func routers() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/questions", func(r chi.Router) {
		r.Get("/", api.AllQuestion)
		r.Get("/{id}", api.DetailQuestion)
		r.Post("/", api.CreateQuestion)
		r.Put("/{id}", api.UpdateQuestion)
		r.Delete("/{id}", api.DeleteQuestion)
	})
	return r
}
