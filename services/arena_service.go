package services

import (
    "fmt"
    "math/rand"
    "time"
    "magical_arena/models"
)

// Arena struct represents the battle arena where two players fight
type Arena struct {
    PlayerA *models.Player
    PlayerB *models.Player
}

// RollDice method simulates rolling a 6-sided dice
func (a *Arena) RollDice() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(6) + 1
}

// Fight method simulates the fight between PlayerA and PlayerB until one of them dies
func (a *Arena) Fight() string {
	
	var playerAttacking *models.Player
	var playerDefending *models.Player
	if a.PlayerA.Health <= a.PlayerB.Health {
		playerAttacking=a.PlayerA
		playerDefending=a.PlayerB
	} else {
		playerAttacking=a.PlayerB
		playerDefending=a.PlayerA
	}

    for {
        if !a.PlayerA.IsAlive() {
            return fmt.Sprintf("%s wins!", a.PlayerB.Name)
        }
        if !a.PlayerB.IsAlive() {
            return fmt.Sprintf("%s wins!", a.PlayerA.Name)
        }
        a.Attack(playerAttacking,playerDefending)
		playerAttacking,playerDefending=playerDefending,playerAttacking
        // Player with lower health attacks first
       
    }
}

// attack method handles the attack logic between two players
func (a *Arena) Attack(attacker, defender *models.Player) []int {
    attackRoll := a.RollDice()      // Attacker rolls the dice
    defendRoll := a.RollDice()      // Defender rolls the dice

    damage := attacker.Attack * attackRoll      // Calculate attack damage
    defense := defender.Strength * defendRoll   // Calculate defense

    actualDamage := damage - defense     // Calculate the actual damage after defense
    if actualDamage > 0 {
        defender.Health -= actualDamage  // Reduce defender's health by the actual damage
    }

    fmt.Printf("%s attacks %s! Attack roll: %d, Defense roll: %d, Damage: %d, Defense: %d, Actual Damage: %d, %s Health: %d\n",
        attacker.Name, defender.Name, attackRoll, defendRoll, damage, defense, actualDamage, defender.Name, defender.Health)
	var rolls []int
	rolls = append(rolls, attackRoll)
	rolls=append(rolls, defendRoll)
	return rolls
}
