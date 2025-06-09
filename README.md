# 🎮 TCP Game Server-Client

<div align="center">
  <img src="https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version"/>
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge" alt="License"/>
  <img src="https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20MacOS-blue?style=for-the-badge" alt="Platform"/>
</div>

<br/>

<div align="center">
  <h3>🚀 A powerful TCP-based game server-client implementation with advanced features</h3>
</div>

<br/>

## 📋 Key Features

<table>
<tr>
<td width="50%">

### 🎯 Authentication
- Secure user login system
- Account registration
- Password encryption
- Session management

</td>
<td width="50%">

### 🎮 Game Management
- Real-time game sessions
- Player matchmaking
- Game state synchronization
- Multiplayer support

</td>
</tr>
<tr>
<td width="50%">

### 📊 Ranking System
- Player statistics
- Global leaderboards
- Achievement tracking
- Performance metrics

</td>
<td width="50%">

### 🔒 Security
- Data encryption
- Anti-cheat protection
- DDoS prevention
- Secure communication

</td>
</tr>
</table>

## 🛠️ Installation

### Prerequisites
- Go 1.16 or higher
- Git
- Basic knowledge of TCP networking

### Quick Start

```bash
# Clone the repository
git clone https://github.com/yourusername/tcp-game-server.git

# Navigate to project directory
cd tcp-game-server

# Install dependencies
go mod download

# Build the project
go build -o server.exe server.go
go build -o client.exe client.go
```

## 🎮 How to Play

### Starting the Server
```bash
./server.exe
```

### Starting the Client
```bash
./client.exe
```

### Available Commands

<table>
<tr>
<th>Command</th>
<th>Description</th>
</tr>
<tr>
<td><code>/login [username] [password]</code></td>
<td>Login to your account</td>
</tr>
<tr>
<td><code>/register [username] [password]</code></td>
<td>Create a new account</td>
</tr>
<tr>
<td><code>/play</code></td>
<td>Start a new game session</td>
</tr>
<tr>
<td><code>/quit</code></td>
<td>Exit the game</td>
</tr>
</table>

## 📁 Project Structure

```
.
├── server.go          # Main server implementation
├── client.go          # Main client implementation
├── clients.go         # Client management system
├── json_utils.go      # JSON data handling
├── level_utils.go     # Level progression system
├── go.mod            # Go module dependencies
└── README.md         # Project documentation
```

## 🔒 Security Features

<div align="center">
<table>
<tr>
<td align="center">
  <b>Data Protection</b><br/>
  <img src="https://img.shields.io/badge/Encryption-AES256-green?style=flat-square" alt="Encryption"/>
</td>
<td align="center">
  <b>Authentication</b><br/>
  <img src="https://img.shields.io/badge/Auth-JWT-blue?style=flat-square" alt="Authentication"/>
</td>
<td align="center">
  <b>Protection</b><br/>
  <img src="https://img.shields.io/badge/Security-DDoS%20Protection-red?style=flat-square" alt="Security"/>
</td>
</tr>
</table>
</div>

## 🤝 Contributing

We welcome contributions! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

<div align="center">
  <img src="https://img.shields.io/badge/PRs-Welcome-brightgreen.svg?style=for-the-badge" alt="PRs Welcome"/>
</div>

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

<div align="center">
  <sub>Built with ❤️ by Your Name</sub>
</div> 