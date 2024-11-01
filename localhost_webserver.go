package main

import (
    "fmt"
    "log"
    "net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
    // Write a simple HTML response
    fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Welcome to My Go Web Server</title>
        </head>
        <body>
            <h1>Welcome to the Go Web Server!</h1>
        </body>
        </html>
    `)
}

func main() {
    // Set up the HTTP route and handler
    http.HandleFunc("/", HomePageHandler)

    // Start the server on port 8080
    log.Println("Starting server on http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}

