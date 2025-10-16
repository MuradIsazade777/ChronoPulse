package storage

import (
    "encoding/json"
    "os"
    "sync"
    "time"

    "chronopulse/internal/logger"
)

var (
    mu     sync.Mutex
    events []time.Time
)

func Append(t time.Time) {
    mu.Lock()
    defer mu.Unlock()
    events = append(events, t)
}

func SaveFinal() {
    mu.Lock()
    defer mu.Unlock()

    file, err := os.Create("assets/pulses.json")
    if err != nil {
        logger.Error("Failed to create file: " + err.Error())
        return
    }
    defer file.Close()

    enc := json.NewEncoder(file)
    enc.SetIndent("", "  ")
    if err := enc.Encode(events); err != nil {
        logger.Error("Failed to encode JSON: " + err.Error())
    }
}
