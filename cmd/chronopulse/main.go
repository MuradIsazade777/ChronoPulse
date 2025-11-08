package main 

import (
    "chronopulse/internal/analyzer"
    "chronopulse/internal/collector"
    "chronopulse/internal/config"
    "chronopulse/internal/logger"
    "chronopulse/internal/storage"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    logger.Init()
    cfg := config.Load()

    pulseChan := make(chan time.Time, 100)
    done := make(chan bool)

    go collector.Start(pulseChan, done, cfg)
    go analyzer.Start(pulseChan, done, cfg)

    // Signal handling
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

    logger.Info("ChronoPulse started. Press Ctrl+C to stop.")
    <-sig

    logger.Info("Shutting down...")
    done <- true
    time.Sleep(1 * time.Second)
    storage.SaveFinal()
    logger.Info("Goodbye.")
}
