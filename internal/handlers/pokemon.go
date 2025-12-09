package handlers

import (
	"encoding/json"
	"net/http"
)

// fetch list of pokemon

type PokemonData struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

func GetPokeMons(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=151")
	if err != nil {
		http.Error(w, "failed to fetch pokemons", 500)
		return
	}

	defer response.Body.Close()

	var pokemonList PokemonData
	json.NewDecoder(response.Body).Decode(&pokemonList)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemonList)

}
