package main

import (
	"github.com/go-chi/chi"
	"github.com/khanadnanxyz/konta/api"
	_ "github.com/khanadnanxyz/konta/helper"
	"net/http"
)

func main() {
	router := routers()
	http.ListenAndServe(":8080", router)
}


func routers() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/questions", api.AllQuestion)
	router.Get("/question/{id}", api.DetailQuestion)
	router.Post("/question", api.CreateQuestion)
	router.Put("/question/{id}", api.UpdateQuestion)
	router.Delete("/question/{id}", api.DeleteQuestion)
	return router
}
