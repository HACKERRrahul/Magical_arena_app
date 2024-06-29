package tests

import (
	"magical_arena/models"
	"magical_arena/services"
	"testing"
)

// TestRollDice checks if RollDice function returns values within the expected range
func TestRollDice(t *testing.T) {
	arena := services.Arena{}

	for i := 0; i < 100; i++ { // Run the test multiple times to check randomness
		roll := arena.RollDice()
		if roll < 1 || roll > 6 {
			t.Errorf("RollDice() returned %d, expected value between 1 and 6", roll)
		}
	}
}
// TestAttack checks if the attack function reduces the defender's health correctly
func TestAttack(t *testing.T) {
	playerA := &models.Player{Name: "Player A", Health: 50, Strength: 5, Attack: 10}
	playerB := &models.Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}
	arena := services.Arena{PlayerA: playerA, PlayerB: playerB}
	original_PlayerB_health:=playerB.Health
    rolls:=arena.Attack(playerA,playerB)
	attackRoll:=rolls[0]
	defendRoll:=rolls[1]
	damage:=playerA.Attack*attackRoll-playerB.Strength*defendRoll
	if damage<0 {
		damage=0
	}
	if playerB.Health!=original_PlayerB_health-damage{
		t.Errorf("Attack function is not working properly")
	}
}

