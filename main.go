package main

import (
    "log"
    "net/http"
    "magical_arena/controllers"
)

func main() {
    // Register the FightHandler to handle requests to the /fight endpoint
    http.HandleFunc("/fight", controllers.FightHandler)
    
    // Start the HTTP server on port 8080
    log.Fatal(http.ListenAndServe(":8080", nil))
}
