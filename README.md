# Magical_arena_app

## Overview
This is a Go application that simulates a magical arena where two players fight until one of them loses all health. Each player has health, strength, and attack attributes. The players take turns attacking each other, and the damage dealt is determined by dice rolls.

## How to Run

1. **Clone the repository**:
    ```sh
    git clone <repository_url>
    cd magical_arena
    ```

2. **Run the application**:
    ```sh
    go run main.go
    ```

3. **Run tests**:
    ```sh
    go test ./...
    ```

## Files

- `main.go`: Contains the entry point for the application.
- `controllers/arena_controller.go`: Contains the HTTP handler for the fight API.
- `models/player.go`: Contains the Player struct and related methods.
- `services/arena_service.go`: Contains the Arena struct and methods to handle the fighting logic.
- `tests/arena_test.go`: Contains unit tests for checking controllers.FightHandler is working or not.
-`tests/arena_service.go`:Contains unit test for checking arena.Rolldice and arena.attack are working or not.
- `README.md`: Provides information about the application and instructions on how to run it.
- `go.mod`: Go module file for dependencies management.
-`Verification_screenshots/`: Contains api_success.png,test_success.png which shows that api and unit tests are working fine. 




## API

### `http://localhost:8080/fight`
**Method**: `POST`

**Request Body**:
```json
{
  "playerA": {
    "name": "Player A",
    "health": 50,
    "strength": 5,
    "attack": 10
  },
  "playerB": {
    "name": "Player B",
    "health": 100,
    "strength": 10,
    "attack": 5
  }
}


