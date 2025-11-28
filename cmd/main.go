package main

import (
	"os"
	"path/filepath"

	"core.axiom/internal/game"
)

func main() {
	cfg := game.GameConfig{
		TileWidth:    16,
		TileHeight:   16,
		PlayersCount: 2,
	}
	outputDir := "output"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		panic(err)
	}

	g := game.NewGame(cfg)
	data := g.Serialize()

	filePath := filepath.Join(outputDir, "game.json")
	if err := os.WriteFile(filePath, data, 0o644); err != nil {
		panic(err)
	}
}
