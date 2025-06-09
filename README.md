# Game Tower Defense Multiplayer

Một game Tower Defense đơn giản cho phép nhiều người chơi tham gia thông qua mạng LAN. Game được viết bằng Go, sử dụng TCP socket để giao tiếp giữa client và server.

## 🎮 Tính Năng

- Đăng nhập với username/password
- Hệ thống level và EXP
- Nhiều loại quân với khả năng khác nhau
- 3 loại tháp: 2 Guard Tower và 1 King Tower
- Queen có khả năng hồi máu cho tháp
- Tăng 50% sát thương khi tấn công King Tower
- Lưu trữ tiến trình chơi game

## 🛠️ Yêu Cầu Hệ Thống

- Go 1.20 trở lên
- Hệ điều hành: Windows/Linux/MacOS
- Kết nối mạng LAN

## 📦 Cài Đặt

1. Clone repository:
```bash
git clone https://github.com/your-username/tower-defense-game.git
cd tower-defense-game
```

2. Cài đặt dependencies:
```bash
go mod download
```

3. Tạo file cấu hình:
```bash
mkdir data
```

4. Tạo file `data/players.json`:
```json
[
    {
        "username": "player1",
        "password": "pass1",
        "level": 1,
        "exp": 0
    },
    {
        "username": "player2",
        "password": "pass2",
        "level": 1,
        "exp": 0
    }
]
```

5. Tạo file `data/troops.json`:
```json
[
    {
        "name": "Warrior",
        "atk": 100,
        "def": 50,
        "mana": 0
    },
    {
        "name": "Archer",
        "atk": 150,
        "def": 30,
        "mana": 0
    },
    {
        "name": "Queen",
        "atk": 80,
        "def": 40,
        "mana": 200
    }
]
```

## 🚀 Chạy Game

1. Khởi động server:
```bash
go run server/server.go
```

2. Khởi động client (trên máy khác):
```bash
go run clients/clients.go
```

## 🎯 Cách Chơi

1. Đăng nhập với username và password
2. Chờ người chơi khác tham gia
3. Mỗi người chơi nhận được 3 quân ngẫu nhiên
4. Lần lượt chọn quân để tấn công
5. Quân sẽ tự động tấn công tháp đầu tiên còn sống
6. Queen có thể hồi máu cho tháp đồng minh
7. Phá hủy King Tower của đối phương để chiến thắng
8. Người thắng nhận được 30 EXP

## 📊 Hệ Thống Level

- Level 1: 100 EXP
- Level 2: 110 EXP (tăng 10%)
- Level 3: 121 EXP (tăng 10%)
- Và cứ tiếp tục...

## 🏗️ Cấu Trúc Project

```
tower-defense-game/
├── server/
│   └── server.go
├── clients/
│   └── clients.go
├── utils/
│   ├── json_utils.go
│   └── level_utils.go
├── models/
│   ├── player.go
│   ├── troop.go
│   └── tower.go
└── data/
    ├── players.json
    ├── troops.json
    └── towers.json
```

## ⚠️ Lưu Ý

- Server chạy trên cổng 8080
- Timeout cho mỗi lượt chơi là 30 giây
- Tối đa 10 người chơi trong hàng đợi
- Chỉ hỗ trợ 2 người chơi mỗi trận

## 🔧 Cấu Hình

Các thông số có thể điều chỉnh trong code:
- Port: 8080
- Timeout: 30 giây
- Số lượng quân mỗi người: 3
- HP Guard Tower: 1000
- HP King Tower: 2000
- EXP thưởng: 30
- Hệ số tăng EXP: 10%

## 🤝 Đóng Góp

Mọi đóng góp đều được hoan nghênh! Vui lòng tạo issue hoặc pull request.

## 📝 License

MIT License 