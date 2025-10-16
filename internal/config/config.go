package config

import "time"

type Config struct {
    PulseInterval   time.Duration
    AnalyzeInterval time.Duration
    MaxBuffer       int
}

func Load() Config {
    return Config{
        PulseInterval:   1 * time.Second,
        AnalyzeInterval: 5 * time.Second,
        MaxBuffer:       100,
    }
}
