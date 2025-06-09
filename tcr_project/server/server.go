// Package server chứa các hàm xử lý server game
package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"tcr_project/models"

	"tcr_project/utils"

	"math/rand"
	"time"
)

// StartServer khởi động server game trên cổng 8080
func StartServer() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Không thể mở cổng:", err)
		return
	}
	defer ln.Close()

	fmt.Println("🔌 Server đang chạy trên cổng 8080...")

	go matchPlayers()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Lỗi kết nối:", err)
			continue
		}
		go handleClient(conn)
	}

}

// matchPlayers ghép cặp người chơi từ hàng đợi
func matchPlayers() {
	for {
		// Lấy 2 người chơi từ hàng đợi
		player1 := <-lobbyQueue
		player2 := <-lobbyQueue

		// Kiểm tra kết nối còn sống không
		if _, err := player1.Write([]byte{}); err != nil {
			fmt.Println("Player 1 mất kết nối trước khi bắt đầu game")
			player1.Close()
			continue
		}
		if _, err := player2.Write([]byte{}); err != nil {
			fmt.Println("Player 2 mất kết nối trước khi bắt đầu game")
			player2.Close()
			continue
		}

		// Bắt đầu game mới với 2 người chơi
		go startGame(player1, player2)
	}
}

// startGame khởi tạo và chạy một trận đấu mới
func startGame(conn1, conn2 net.Conn) {
	defer conn1.Close()
	defer conn2.Close()

	// Load danh sách quân từ file JSON
	troopsList, err := utils.LoadTroops("data/troops.json")
	if err != nil {
		sendMessage(conn1, "❌ Lỗi load troop data!")
		sendMessage(conn2, "❌ Lỗi load troop data!")
		return
	}

	// Phát ngẫu nhiên 3 quân cho mỗi người chơi
	player1Troops := getRandomTroops(troopsList, 3)
	player2Troops := getRandomTroops(troopsList, 3)

	// Hiển thị quân cho người chơi 1
	sendMessage(conn1, "🎲 Bạn đã nhận được 3 quân:")
	for _, t := range player1Troops {
		sendMessage(conn1, fmt.Sprintf("  • %s (ATK: %d, DEF: %d, Mana: %d)", t.Name, t.ATK, t.DEF, t.Mana))
	}
	sendMessage(conn1, "\n📋 Hướng dẫn:")
	sendMessage(conn1, "  • Mỗi lượt bạn sẽ chọn 1 quân để tấn công")
	sendMessage(conn1, "  • Quân sẽ tự động tấn công tower đầu tiên còn sống")
	sendMessage(conn1, "  • King Tower sẽ nhận 50% sát thương")
	sendMessage(conn1, "  • Phá hủy King Tower để chiến thắng")
	sendMessage(conn1, "") // Thêm dòng trống để tách biệt

	// Hiển thị quân cho người chơi 2
	sendMessage(conn2, "🎲 Bạn đã nhận được 3 quân:")
	for _, t := range player2Troops {
		sendMessage(conn2, fmt.Sprintf("  • %s (ATK: %d, DEF: %d, Mana: %d)", t.Name, t.ATK, t.DEF, t.Mana))
	}
	sendMessage(conn2, "\n📋 Hướng dẫn:")
	sendMessage(conn2, "  • Mỗi lượt bạn sẽ chọn 1 quân để tấn công")
	sendMessage(conn2, "  • Quân sẽ tự động tấn công tower đầu tiên còn sống")
	sendMessage(conn2, "  • King Tower sẽ nhận   150%   sát thương")
	sendMessage(conn2, "  • Phá hủy King Tower để chiến thắng")
	sendMessage(conn2, "") // Thêm dòng trống để tách biệt

	sendMessage(conn1, "🎮 Trận đấu bắt đầu! Bạn là Người chơi 1")
	sendMessage(conn2, "🎮 Trận đấu bắt đầu! Bạn là Người chơi 2")

	// Load thông tin người chơi từ file JSON
	players, err := utils.LoadPlayers("data/players.json")
	if err != nil {
		sendMessage(conn1, "❌ Lỗi load player data!")
		sendMessage(conn2, "❌ Lỗi load player data!")
		return
	}

	player1 := utils.FindPlayer(conn1, players)
	player2 := utils.FindPlayer(conn2, players)

	// Chạy vòng lặp trận đấu
	winner := runSimpleBattleLoop(conn1, conn2, player1Troops, player2Troops)

	// Tính EXP cho người thắng
	if winner == 0 && player1 != nil {
		player1.EXP += 30
		levelsGained := utils.CheckLevelUp(player1)
		if levelsGained > 0 {
			sendMessage(conn1, fmt.Sprintf("🎉 Bạn đã lên %d level mới!", levelsGained))
		}
		sendMessage(conn1, "🎉 Bạn nhận được +30 EXP!")
	} else if winner == 1 && player2 != nil {
		player2.EXP += 30
		levelsGained := utils.CheckLevelUp(player2)
		if levelsGained > 0 {
			sendMessage(conn2, fmt.Sprintf("🎉 Bạn đã lên %d level mới!", levelsGained))
		}
		sendMessage(conn2, "🎉 Bạn nhận được +30 EXP!")
	}

	// Lưu lại thông tin người chơi
	err = utils.SavePlayers("data/players.json", players)
	if err != nil {
		sendMessage(conn1, "❌ Không thể lưu dữ liệu người chơi.")
		sendMessage(conn2, "❌ Không thể lưu dữ liệu người chơi.")
	}
}

// sendMessage gửi tin nhắn đến người chơi với format nhất quán
func sendMessage(conn net.Conn, message string) {
	conn.Write([]byte(message + "\n"))
}

// handleClient xử lý kết nối từ một client mới
func handleClient(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Printf("Client %s đã ngắt kết nối\n", conn.RemoteAddr().String())
	}()

	reader := bufio.NewReader(conn)

	// Set timeout cho các hoạt động đọc/ghi
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	conn.SetWriteDeadline(time.Now().Add(30 * time.Second))

	// Yêu cầu đăng nhập
	sendMessage(conn, "👤 Nhập username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Lỗi đọc username từ %s: %v\n", conn.RemoteAddr().String(), err)
		return
	}
	username = strings.TrimSpace(username)

	sendMessage(conn, "🔑 Nhập password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Lỗi đọc password từ %s: %v\n", conn.RemoteAddr().String(), err)
		return
	}
	password = strings.TrimSpace(password)

	if isValidLogin(username, password) {
		sendMessage(conn, "✅ Đăng nhập thành công! Đang chờ người chơi khác...")

		// Reset deadline trước khi thêm vào hàng đợi
		conn.SetReadDeadline(time.Time{}) // Không timeout
		conn.SetWriteDeadline(time.Time{})

		select {
		case lobbyQueue <- conn:
			// Kết nối đã được thêm vào hàng đợi
			// Giữ kết nối mở cho đến khi game bắt đầu
			for {
				// Kiểm tra kết nối còn sống không
				_, err := conn.Write([]byte{})
				if err != nil {
					fmt.Printf("Client %s mất kết nối trong phòng chờ\n", conn.RemoteAddr().String())
					return
				}
				time.Sleep(5 * time.Second)
			}
		default:
			sendMessage(conn, "❌ Phòng chờ đã đầy, vui lòng thử lại sau.")
			return
		}
	} else {
		sendMessage(conn, "❌ Sai thông tin đăng nhập.")
		return
	}
}

// isValidLogin kiểm tra thông tin đăng nhập từ file players.json
func isValidLogin(username, password string) bool {
	file, err := os.Open("data/players.json")
	if err != nil {
		fmt.Println("Không thể mở file players.json:", err)
		return false
	}
	defer file.Close()

	var players []models.Player
	if err := json.NewDecoder(file).Decode(&players); err != nil {
		fmt.Println("Lỗi đọc JSON:", err)
		return false
	}

	for _, p := range players {
		if p.Username == username && p.Password == password {
			return true
		}
	}
	return false
}

// Hàng đợi chờ người chơi, tối đa 10 người
var (
	lobbyQueue = make(chan net.Conn, 10)
)

// runSimpleBattleLoop chạy vòng lặp trận đấu chính
func runSimpleBattleLoop(conn1 net.Conn, conn2 net.Conn, troops1, troops2 []models.Troop) int {
	playerConns := []net.Conn{conn1, conn2}
	currentPlayer := 0

	// Tạo towers mặc định cho mỗi người chơi
	towerSet := []models.Tower{
		{Name: "Guard Tower 1", HP: 1000, ATK: 300, DEF: 50}, // Tower bảo vệ 1
		{Name: "Guard Tower 2", HP: 1000, ATK: 300, DEF: 50}, // Tower bảo vệ 2
		{Name: "King Tower", HP: 2000, ATK: 500, DEF: 50},    // Tower chính
	}

	// Khởi tạo towers cho cả 2 người chơi
	playerTowers := [][]models.Tower{
		append([]models.Tower{}, towerSet...),
		append([]models.Tower{}, towerSet...),
	}

	// Hiển thị trạng thái ban đầu cho cả hai người chơi
	for i := 0; i < 2; i++ {
		sendMessage(playerConns[i], "\n🏰 Trạng thái ban đầu:")
		for _, t := range playerTowers[i] {
			sendMessage(playerConns[i], fmt.Sprintf("  • %s: HP %d", t.Name, t.HP))
		}
	}

	// Vòng lặp chính của trận đấu
	for {
		attacker := playerConns[currentPlayer]
		defender := playerConns[1-currentPlayer]

		var myTroops []models.Troop
		if currentPlayer == 0 {
			myTroops = troops1
		} else {
			myTroops = troops2
		}

		// Hiển thị lượt chơi
		sendMessage(attacker, fmt.Sprintf("\n🕐 Lượt của bạn (Player %d)", currentPlayer+1))
		sendMessage(attacker, "💂 Chọn quân (1, 2, 3):")

		// Hiển thị danh sách quân
		for i, t := range myTroops {
			sendMessage(attacker, fmt.Sprintf("  %d. %s (ATK: %d, DEF: %d)", i+1, t.Name, t.ATK, t.DEF))
		}

		choiceStr := readLine(attacker)
		if choiceStr == "" {
			sendMessage(attacker, "❌ Bạn đã ngắt kết nối!")
			sendMessage(defender, fmt.Sprintf("🎉 Người chơi %d đã ngắt kết nối. Bạn thắng!", currentPlayer+1))
			return 1 - currentPlayer // người chơi còn lại thắng
		}

		idx, err := strconv.Atoi(choiceStr)
		if err != nil || idx < 1 || idx > 3 {
			sendMessage(attacker, "❌ Lựa chọn không hợp lệ! Vui lòng chọn lại (1-3)")
			continue
		}

		troop := myTroops[idx-1]

		// Tìm tower đầu tiên còn sống để tấn công
		var targetTower *models.Tower
		var targetIndex int

		// Nếu là Queen, tìm tower có HP thấp nhất để hồi máu
		if troop.Name == "Queen" {
			minHP := 9999
			for i := 0; i < 3; i++ {
				if playerTowers[currentPlayer][i].HP > 0 && playerTowers[currentPlayer][i].HP < minHP {
					targetTower = &playerTowers[currentPlayer][i]
					targetIndex = i
					minHP = playerTowers[currentPlayer][i].HP
				}
			}
		} else {
			// Các quân khác tìm tower địch để tấn công
			for i := 0; i < 3; i++ {
				if playerTowers[1-currentPlayer][i].HP > 0 {
					targetTower = &playerTowers[1-currentPlayer][i]
					targetIndex = i
					break
				}
			}
		}

		if targetTower == nil {
			sendMessage(attacker, "❌ Không còn tower nào để tấn công!")
			continue
		}

		// Xử lý Queen hồi máu
		if troop.Name == "Queen" {
			healAmount := 200 // Queen hồi 200 HP
			targetTower.HP += healAmount

			// Hiển thị thông tin hồi máu
			sendMessage(attacker, fmt.Sprintf("\n💖 Queen hồi máu cho %s:", targetTower.Name))
			sendMessage(attacker, fmt.Sprintf("   • Hồi: %d HP", healAmount))
			sendMessage(attacker, fmt.Sprintf("   • HP hiện tại: %d", targetTower.HP))

			// Hiển thị trạng thái trụ cho người chơi
			sendMessage(attacker, "\n🏰 Trạng thái trụ của bạn:")
			for _, t := range playerTowers[currentPlayer] {
				if t.HP > 0 {
					sendMessage(attacker, fmt.Sprintf("  • %s: HP %d", t.Name, t.HP))
				} else {
					sendMessage(attacker, fmt.Sprintf("  • %s: Đã bị phá hủy", t.Name))
				}
			}
		} else {
			// Tính sát thương cho các quân khác
			baseDamage := troop.ATK - targetTower.DEF
			if baseDamage < 0 {
				baseDamage = troop.ATK / 2 // Đảm bảo luôn có sát thương tối thiểu bằng 50% ATK
			}

			// Áp dụng hệ số sát thương đặc biệt cho King Tower
			if targetIndex == 2 { // King Tower
				baseDamage = int(float64(baseDamage) * 1.5) // Tăng 50% sát thương cho King Tower
			}

			// Đảm bảo sát thương tối thiểu là 1
			if baseDamage < 1 {
				baseDamage = 1
			}

			// Cập nhật HP
			targetTower.HP -= baseDamage
			if targetTower.HP < 0 {
				targetTower.HP = 0
			}

			// Hiển thị thông tin tấn công
			sendMessage(attacker, fmt.Sprintf("\n🎯 %s tấn công %s:", troop.Name, targetTower.Name))
			sendMessage(attacker, fmt.Sprintf("   • Gây ra: %d sát thương", baseDamage))
			sendMessage(attacker, fmt.Sprintf("   • HP còn lại: %d", targetTower.HP))

			// Hiển thị trạng thái trụ cho người phòng thủ
			sendMessage(defender, "\n🏰 Trạng thái trụ của bạn:")
			for _, t := range playerTowers[1-currentPlayer] {
				if t.HP > 0 {
					sendMessage(defender, fmt.Sprintf("  • %s: HP %d", t.Name, t.HP))
				} else {
					sendMessage(defender, fmt.Sprintf("  • %s: Đã bị phá hủy", t.Name))
				}
			}

			if targetTower.HP <= 0 {
				notice := fmt.Sprintf("\n💥 %s đã bị phá hủy!", targetTower.Name)
				sendMessage(attacker, notice)
				sendMessage(defender, notice)
			}
		}

		// Kiểm tra King Tower bị phá chưa
		if playerTowers[1-currentPlayer][2].HP <= 0 {
			winMsg := fmt.Sprintf("\n🏆 Người chơi %d đã chiến thắng!", currentPlayer+1)
			sendMessage(attacker, winMsg)
			sendMessage(defender, winMsg)

			// Gửi thông báo kết thúc game
			sendMessage(attacker, "\n🎮 Trận đấu kết thúc!")
			sendMessage(defender, "\n🎮 Trận đấu kết thúc!")
			break
		}

		// Đổi lượt
		currentPlayer = 1 - currentPlayer
	}
	return currentPlayer
}

// readLine đọc một dòng từ kết nối với timeout
func readLine(conn net.Conn) string {
	// Reset deadline mỗi lần đọc
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	reader := bufio.NewReader(conn)
	text, err := reader.ReadString('\n')
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			conn.Write([]byte("⚠️ Hết thời gian chờ, vui lòng thử lại.\n"))
		} else {
			conn.Write([]byte("🔌 Mất kết nối.\n"))
		}
		conn.Close()
		return ""
	}
	return strings.TrimSpace(text)
}

// getRandomTroops chọn ngẫu nhiên n quân từ danh sách
func getRandomTroops(all []models.Troop, count int) []models.Troop {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var chosen []models.Troop
	used := map[int]bool{}

	for len(chosen) < count {
		idx := r.Intn(len(all))
		if !used[idx] {
			chosen = append(chosen, all[idx])
			used[idx] = true
		}
	}
	return chosen
}
