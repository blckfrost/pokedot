package db

import (
	"database/sql"
	"fmt"

	"github.com/blckfrost/pokedot.git/config"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type Pokemon struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	PokedexID int    `json:"pokedex_id"`
	Type      string `json:"type"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	SpriteURL string `json:"sprite_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func Init(config *config.Config) error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBName,
		config.DBPassword,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	return nil
}
