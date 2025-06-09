// Package utils chứa các hàm tiện ích cho việc tính toán level và EXP
package utils

import "tcr_project/models"

// RequiredEXP tính số EXP cần thiết để lên level tiếp theo
// Công thức: Mỗi level tăng 10% EXP so với level trước
// Ví dụ:
//   - Level 1: 100 EXP
//   - Level 2: 110 EXP (100 + 10%)
//   - Level 3: 121 EXP (110 + 10%)
//   - ...
func RequiredEXP(level int) int {
	baseEXP := 100 // EXP cơ bản cho level 1
	totalEXP := baseEXP

	// Tính EXP cho các level tiếp theo
	// Level 1 → cần 100
	// Level 2 → cần 100 + 10%
	// Level 3 → 100 + 10% + 10% ...
	for i := 2; i <= level; i++ {
		totalEXP = int(float64(totalEXP) * 1.1) // tăng 10%
	}
	return totalEXP
}

// CheckLevelUp kiểm tra và nâng cấp level cho người chơi
// Trả về số level đã tăng
// Quy trình:
//   1. Kiểm tra EXP hiện tại có đủ để lên level không
//   2. Nếu đủ, trừ EXP và tăng level
//   3. Lặp lại cho đến khi không đủ EXP
func CheckLevelUp(player *models.Player) int {
	levelsGained := 0 // Số level đã tăng
	for {
		expNeed := RequiredEXP(player.Level) // EXP cần để lên level tiếp
		if player.EXP >= expNeed {
			player.EXP -= expNeed // Trừ EXP đã dùng
			player.Level++        // Tăng level
			levelsGained++        // Đếm số level đã tăng
		} else {
			break // Không đủ EXP để lên level tiếp
		}
	}
	return levelsGained
}
