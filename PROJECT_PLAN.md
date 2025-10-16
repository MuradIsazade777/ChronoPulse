# ChronoPulse Project Plan

## 🧠 Concept

ChronoPulse is a real-time event tracking engine that collects timestamped pulses, analyzes their frequency, and visualizes them in the terminal. It is designed for performance, modularity, and open-source clarity.

## 🏗️ Architecture

- **cmd/**: Entry point
- **internal/**: Modular logic
  - `collector`: Goroutine-based pulse ingestion
  - `analyzer`: Time-windowed analysis
  - `visualizer`: ASCII chart rendering
  - `storage`: JSON persistence
  - `config`: Runtime parameters
  - `logger`: Logging abstraction

## 🔄 Flow Diagram
[User Input / Timer]
        ↓
   [collector.go]
        ↓
[pulseChan (channel)]
        ↓
   [analyzer.go]
        ↓
   [visualizer.go]
        ↓
 Terminal Output
        ↓
   [storage.go]
        ↓
assets/pulses.json


Code

## 📈 Module Breakdown

| Module      | Responsibility                  | Key Concepts         |
|-------------|----------------------------------|----------------------|
| collector   | Collects pulses via ticker       | goroutine, channel   |
| analyzer    | Groups pulses by time            | slice, map, sync     |
| visualizer  | Renders ASCII charts             | color, terminal UX   |
| storage     | Saves pulses to JSON             | encoding/json, file  |
| config      | Defines intervals and limits     | struct, duration     |
| logger      | Logs system messages             | log, stdout/stderr   |

## 🧪 Testing Strategy

- Manual testing via `go run`
- JSON output validation
- Terminal visualization accuracy
- Graceful shutdown handling

## 📦 Future Enhancements

- Web dashboard (optional)
- Pulse tagging
- Export to CSV
- Configurable visualization styles
