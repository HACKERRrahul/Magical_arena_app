package models

// Player struct represents a player in the arena with health, strength, and attack attributes
type Player struct {
    Name     string `json:"name"`
    Health   int    `json:"health"`
    Strength int    `json:"strength"`
    Attack   int    `json:"attack"`
}

// IsAlive method checks if the player is still alive (health > 0)
func (p *Player) IsAlive() bool {
    return p.Health > 0
}
