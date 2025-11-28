package game

import (
	"encoding/json"

	"core.axiom/internal/game/simulation"
	"core.axiom/internal/players"
	"core.axiom/internal/tilemap"
)

type GameCore struct {
	Players    []players.Player
	Map        tilemap.TileMap
	Simulation simulation.GameSimulation
}

func NewGame(cfg GameConfig) *GameCore {
	return nil
}

func (c *GameCore) Serialize() []byte {
	data, _ := json.Marshal(c)
	return data
}
