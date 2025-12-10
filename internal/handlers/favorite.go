package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/blckfrost/pokedot.git/internal/redis"
)

var FavoritesKey = "favorite_pokemons"

func GetFavoritePokemons(w http.ResponseWriter, r *http.Request) {
	favoritePokemons, err := redis.Rdb.SMembers(redis.Ctx, FavoritesKey).Result()
	if err != nil {
		http.Error(w, "failed to fetch favorites", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(favoritePokemons)

}

func AddToFavorites(w http.ResponseWriter, r *http.Request) {
	var body FavoriteRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if body.Name == "" {
		http.Error(w, "pokemon name required", http.StatusBadRequest)
		return
	}

	err = redis.Rdb.SAdd(redis.Ctx, FavoritesKey, body.Name).Err()
	if err != nil {
		http.Error(w, "failed to add favorites", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("added to favorites"))

}

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	var body FavoriteRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "invalid body request", http.StatusBadRequest)
		return
	}

	if body.Name == "" {
		http.Error(w, "pokemon name required", http.StatusBadRequest)
		return
	}

	err = redis.Rdb.SRem(redis.Ctx, FavoritesKey, body.Name).Err()
	if err != nil {
		http.Error(w, "failed to delete favorite ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("removed from favorite"))

}

type FavoriteRequest struct {
	Name string `json:"name"`
}
