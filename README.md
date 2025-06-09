Dựa trên nội dung từ **Documentation.pdf** và bố cục hiện tại của **README.md**, mình đã cải tiến lại file `README.md` sao cho đồng bộ, chuyên nghiệp và chứa đầy đủ thông tin kỹ thuật, hướng dẫn triển khai, cách chơi, và kiến trúc hệ thống.

Dưới đây là phiên bản đã được làm lại:

---

```markdown
# 🏰 TCR - Tower Defense TCP Game (Multiplayer)

<div align="center">
  <img src="https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20MacOS-blue?style=for-the-badge" />
</div>

<div align="center">
  <h3>⚔️ Multiplayer tower defense game using Go and TCP socket programming</h3>
</div>

---

## 🧩 System Overview

TCR is a real-time, multiplayer tower defense game built using Go. It operates on a client-server model with TCP sockets for communication. Two players battle using assigned troops to destroy the opponent's King Tower.

### 🧱 Architecture Diagram

```

┌────────────┐   TCP Socket   ┌────────────┐
│   Client   │◀──────────────▶│   Server   │
│ clients.go │                │ server.go  │
└────────────┘                └────────────┘
│
▼
┌─────────────┐
│  JSON Data  │
└─────────────┘

```

---

## 📦 Project Structure

```

tcr\_project/
├── main.go              # Entry point
├── clients.go           # Client logic
├── server/server.go     # Server logic
├── models/entities.go   # Data models: Player, Troop, Tower
├── utils/
│   ├── json\_utils.go    # JSON read/write
│   └── level\_utils.go   # Leveling system
├── data/
│   ├── players.json     # User credentials
│   ├── troops.json      # Troop definitions
│   └── towers.json      # Tower definitions
└── go.mod               # Go module definition

````

---

## 🔑 Features

- 🔐 **User Authentication**
- 🎮 **Real-time Battle with Turn-Based Logic**
- 🏹 **Troop vs Tower Mechanics**
- 📈 **EXP System with Level Ups**
- 💬 **Text-Based Interaction**
- 🌐 **Multiplayer over TCP**

---

## 🚀 Getting Started

### Prerequisites

- Go 1.24+ installed
- Open TCP port 8080
- OS: Windows, macOS, or Linux

### Setup

```bash
# Clone project
git clone https://github.com/yourusername/tcr-tower-defense.git
cd tcr-tower-defense

# Create data directory and add required files
mkdir -p data
# Ensure data/players.json, troops.json, and towers.json are correctly filled

# Initialize Go module
go mod tidy
````

---

## ▶️ Running the Game

### Start the Server

```bash
go run main.go
# Output: Server đang chạy trên cổng 8080...
```

### Start Clients (in separate terminals)

```bash
go run clients.go
```

Repeat for second client.

---

## 🎮 Game Flow

1. Players authenticate with username/password.
2. Server matches two players.
3. Each player receives 3 random troops.
4. Players take turns attacking towers.
5. Game ends when a King Tower is destroyed.
6. Winner gains EXP and may level up.

---

## 📡 Protocol Communication (PDUs)

| Phase      | Direction       | Message Type      | Example                                 |
| ---------- | --------------- | ----------------- | --------------------------------------- |
| Auth       | Server → Client | Username Prompt   | "Nhập username:"                        |
|            | Client → Server | Username          | "player1\n"                             |
|            | Server → Client | Auth Result       | "Đăng nhập thành công!"                 |
| Game Setup | Server → Client | Troop Assignment  | "Bạn đã nhận được 3 quân: ..."          |
|            | Server → Client | Game Start        | "Trận đấu bắt đầu! Bạn là Người chơi 1" |
| Battle     | Server → Client | Turn Notification | "Lượt của bạn (Player 1)"               |
|            | Client → Server | Troop Choice      | "1\n"                                   |
|            | Server → Client | Battle Result     | "Pawn tấn công Guard Tower 1: ..."      |
| Game End   | Server → Client | Victory Message   | "Người chơi 1 đã chiến thắng!"          |

---

## ⚙️ Configuration

You can adjust the following in `server.go` and `clients.go`:

```go
// server.go
const serverPort = ":8080"
const lobbyQueueSize = 10
const connectionTimeout = 30

// clients.go
const serverAddress = "localhost:8080"
const maxRetries = 3
const retryDelay = 2 * time.Second
```

---

## 🧪 Troubleshooting

| Issue                   | Solution                               |
| ----------------------- | -------------------------------------- |
| Port already in use     | Change port in `server.go`             |
| Connection refused      | Ensure server is running first         |
| JSON file not found     | Verify `data/` folder and its contents |
| Wrong login credentials | Use valid entries from `players.json`  |

---

## 🧱 Extending the Game

* ➕ Add new troops/towers → Edit `troops.json` or `towers.json`
* 👥 Add more players → Edit `players.json`
* 🔄 Improve logic → Modify `server.go`, `clients.go`, or utils

---

## 🧪 Testing Tips

* Simulate 2 clients concurrently
* Verify EXP and level persistence
* Test timeout and disconnect scenarios
* Balance troop vs tower damage values

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more info.

---

<div align="center">
  <sub>Developed with ❤️ by Nguyễn Hải Quân - ITITWE21104</sub>
</div>
```
