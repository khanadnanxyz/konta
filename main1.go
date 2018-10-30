package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	_ "github.com/khanadnanxyz/konta/helper"
)

func main() {
	routers()
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	var err error

	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)
	db, err = sql.Open("mysql", dbSource)
	catch(err)
}

var router *chi.Mux
var db *sql.DB

func routers() *chi.Mux {
	router.Get("/questions", AllQuestion)
	router.Get("/question/{id}", DetailQuestion)
	router.Post("/question", CreateQuestion)
	router.Put("/question/{id}", UpdateQuestion)
	router.Delete("/question/{id}", DeleteQuestion)

	return router
}

func AllQuestion(w http.ResponseWriter, r *http.Request) {}
func DetailQuestion(w http.ResponseWriter, r *http.Request) {}
func CreateQuestion(w http.ResponseWriter, r *http.Request) {}
func UpdateQuestion(w http.ResponseWriter, r *http.Request) {}
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {}
