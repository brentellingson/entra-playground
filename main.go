package main

import (
	"net/http"

	"github.com/brentellingson/entra-playground/internal/api"
	"github.com/brentellingson/entra-playground/internal/config"
	_ "github.com/brentellingson/entra-playground/internal/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func main() {
	cfg := config.NewConfig()
	server := api.NewServer(cfg)
	mux := server.Register()

	mux.HandleFunc("GET /swagger/", httpSwagger.Handler())
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	})
	http.ListenAndServe(":8080", mux)
}
