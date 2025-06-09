// Package utils chứa các hàm tiện ích cho việc xử lý JSON và tìm kiếm
package utils

import (
	"encoding/json"
	"io"
	"net"
	"os"

	"tcr_project/models"
)

// LoadTroops đọc và parse danh sách quân từ file JSON
// Trả về slice chứa thông tin các quân và error nếu có
func LoadTroops(filename string) ([]models.Troop, error) {
	// Mở file JSON
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Đọc toàn bộ nội dung file
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse JSON vào slice troops
	var troops []models.Troop
	err = json.Unmarshal(bytes, &troops)
	if err != nil {
		return nil, err
	}

	return troops, nil
}

// LoadTowers đọc và parse danh sách tháp từ file JSON
// Trả về slice chứa thông tin các tháp và error nếu có
func LoadTowers(filename string) ([]models.Tower, error) {
	// Mở file JSON
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Đọc toàn bộ nội dung file
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse JSON vào slice towers
	var towers []models.Tower
	err = json.Unmarshal(bytes, &towers)
	if err != nil {
		return nil, err
	}

	return towers, nil
}

// LoadPlayers đọc và parse danh sách người chơi từ file JSON
// Trả về slice chứa thông tin người chơi và error nếu có
func LoadPlayers(filename string) ([]models.Player, error) {
	// Mở file JSON
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Parse JSON trực tiếp từ file vào slice players
	var players []models.Player
	if err := json.NewDecoder(file).Decode(&players); err != nil {
		return nil, err
	}
	return players, nil
}

// SavePlayers lưu danh sách người chơi vào file JSON
// Dữ liệu được format đẹp với indent để dễ đọc
func SavePlayers(filename string, players []models.Player) error {
	// Chuyển đổi slice players thành JSON với định dạng đẹp
	data, err := json.MarshalIndent(players, "", "  ")
	if err != nil {
		return err
	}
	// Ghi file với quyền 0644 (rw-r--r--)
	return os.WriteFile(filename, data, 0644)
}

// FindPlayer tìm người chơi dựa trên kết nối
// Trong phiên bản hiện tại, chỉ hỗ trợ 2 người chơi
// Trả về con trỏ đến người chơi tương ứng hoặc nil nếu không tìm thấy
func FindPlayer(conn net.Conn, players []models.Player) *models.Player {
	// Giả sử chỉ có 2 người, lấy tạm người đầu và sau
	if len(players) >= 2 {
		if conn.RemoteAddr().String() < players[0].Username {
			return &players[0]
		}
		return &players[1]
	}
	return nil
}
