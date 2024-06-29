package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "magical_arena/controllers"
    "magical_arena/models"
)

func TestFightHandler(t *testing.T) {
    // Initialize test players
    playerA := models.Player{Name: "Player A", Health: 50, Strength: 5, Attack: 10}
    playerB := models.Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}

    // Create a fight request
    fightRequest := controllers.FightRequest{
        PlayerA: playerA,
        PlayerB: playerB,
    }

    // Marshal the request body to JSON
    requestBody, _ := json.Marshal(fightRequest)
    req, err := http.NewRequest("POST", "/fight", bytes.NewBuffer(requestBody))
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.FightHandler)

    // Serve the HTTP request
    handler.ServeHTTP(rr, req)

    // Check if the status code is OK
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Decode the response body
    var fightResponse controllers.FightResponse
    err = json.NewDecoder(rr.Body).Decode(&fightResponse)
    if err != nil {
        t.Errorf("could not decode response: %v", err)
    }

    // Check if a winner is returned
    if fightResponse.Winner == "" {
        t.Errorf("expected a winner but got none")
    }
}
