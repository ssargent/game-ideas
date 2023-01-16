package engine

import (
	"fmt"

	"github.com/asdine/storm/v3"
)

type Engine struct {
	DB *storm.DB
}

func Init() (*Engine, error) {
	db, err := storm.Open("/tmp/game-ideas/game.db")
	if err != nil {
		return nil, fmt.Errorf("error creating stormdb %w", err)
	}

	return &Engine{DB: db}, nil
}
