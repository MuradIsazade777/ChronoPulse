package visualizer

import (
    "fmt"
    "strings"
    "time"

    "github.com/fatih/color"
)

func Render(pulses []time.Time) {
    fmt.Println()
    color.Cyan("📊 Pulse Activity (%d events)", len(pulses))

    counts := make(map[string]int)
    for _, t := range pulses {
        key := t.Format("15:04")
        counts[key]++
    }

    for k, v := range counts {
        bar := strings.Repeat("█", v)
        fmt.Printf("%s | %s (%d)\n", k, bar, v)
    }
}
