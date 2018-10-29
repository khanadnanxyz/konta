package main

import (
	"fmt"
	"github.com/khanadnanxyz/konta/config"
	"net/http"
)

func main() {
	config.LoadConfig()
	println(config.AppConfig.GetString("postgres.host"))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Konta !")
}
