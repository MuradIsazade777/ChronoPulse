package collector

import (
    "chronopulse/internal/config"
    "chronopulse/internal/logger"
    "time"
)

func Start(pulseChan chan<- time.Time, done <-chan bool, cfg config.Config) {
    ticker := time.NewTicker(cfg.PulseInterval)
    defer ticker.Stop()

    for {
        select {
        case <-done:
            logger.Info("Collector stopped.")
            return
        case t := <-ticker.C:
            pulseChan <- t
            logger.Debug("Pulse collected at " + t.Format(time.RFC3339))
        }
    }
}
