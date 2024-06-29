package controllers

import (
    "encoding/json"
    "net/http"
    "magical_arena/models"
    "magical_arena/services"
)

// FightRequest struct represents the JSON request body for a fight
type FightRequest struct {
    PlayerA models.Player `json:"playerA"`
    PlayerB models.Player `json:"playerB"`
}

// FightResponse struct represents the JSON response body with the winner of the fight
type FightResponse struct {
    Winner string `json:"winner"`
}

// FightHandler handles the HTTP request for initiating a fight
func FightHandler(w http.ResponseWriter, r *http.Request) {
    // Ensure the request method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Decode the request body into a FightRequest struct
    var fightRequest FightRequest
    err := json.NewDecoder(r.Body).Decode(&fightRequest)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Initialize the Arena with the provided players
    arena := services.Arena{
        PlayerA: &fightRequest.PlayerA,
        PlayerB: &fightRequest.PlayerB,
    }

    // Simulate the fight and get the winner
    winner := arena.Fight()

    // Encode the response with the winner and send it back
    fightResponse := FightResponse{Winner: winner}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(fightResponse)
}
