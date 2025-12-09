package main

import (
	"log"
	"net/http"

	"github.com/blckfrost/pokedot.git/internal/handlers"
	"github.com/blckfrost/pokedot.git/internal/redis"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	redis.Init()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:4321"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	r.Get("/api/pokemon", handlers.GetPokeMons)

	log.Println("server running on :3030")
	http.ListenAndServe(":3030", r)
}
