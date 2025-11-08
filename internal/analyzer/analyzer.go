package analyzer 

import (
    "chronopulse/internal/config"
    "chronopulse/internal/logger"
    "chronopulse/internal/storage"
    "chronopulse/internal/visualizer"
    "time"
)

func Start(pulseChan <-chan time.Time, done <-chan bool, cfg config.Config) {
    buffer := make([]time.Time, 0, cfg.MaxBuffer)

    ticker := time.NewTicker(cfg.AnalyzeInterval)
    defer ticker.Stop()

    for {
        select {
        case <-done:
            logger.Info("Analyzer stopped.")
            return
        case pulse := <-pulseChan:
            buffer = append(buffer, pulse)
            storage.Append(pulse)
        case <-ticker.C:
            if len(buffer) > 0 {
                visualizer.Render(buffer)
                buffer = buffer[:0]
            }
        }
    }
}
