# ChronoPulse

ChronoPulse is a modular, terminal-based pulse tracking engine built entirely in Go. It collects timestamped events in real time, analyzes their frequency, and visualizes activity using ASCII charts. Designed for performance, clarity, and open-source extensibility. 

## 🚀 Features 

- Real-time pulse collection via goroutines and channels
- Time-based statistical analysis 
- ASCII-based terminal visualization 
- Modular architecture with internal packages 
- Local JSON storage for event history
- Graceful shutdown with signal handling
- Fully responsive CLI experience 

## 📦 Tech Stack

- **Language**: Go 1.21+
- **Concurrency**: Goroutines, Channels
- **Storage**: Local JSON
- **Visualization**: ASCII + `fatih/color`
- **Structure**: `cmd/`, `internal/`, `assets/`

## 🧪 Getting Started

```bash
git clone https://github.com/yourusername/chronopulse.git
cd chronopulse
go run cmd/chronopulse/main.go
``` 

📊 Output Example
Code
📊 Pulse Activity (5 events)
19:13 | ███ (3)
19:14 | ██ (2)
📜 License
ChronoPulse is released under the MIT License.

🤝 Contributing
We welcome contributions! See CONTRIBUTING.md for guidelines.

🔐 Security
Please report vulnerabilities via SECURITY.md.
