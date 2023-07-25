package main

import (
    "fmt"
    "log"
    "os"
)


func main() {
    portString := os.Getenv("PORT")
    if portString == "" {
        log.Fatal("PORT is not found in the environment")
    }

    fmt.Println("Port:", portString)
}
