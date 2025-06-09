package models

type Tower struct {
	Name string  `json:"name"`
	HP   int     `json:"hp"`
	ATK  int     `json:"atk"`
	DEF  int     `json:"def"`
	CRIT float64 `json:"crit"` // tỉ lệ phần trăm, ví dụ: 0.1 cho 10%
	EXP  int     `json:"exp"`
}

type Troop struct {
	Name    string `json:"name"`
	HP      int    `json:"hp"`
	ATK     int    `json:"atk"`
	DEF     int    `json:"def"`
	Mana    int    `json:"mana"`
	EXP     int    `json:"exp"`
	Special string `json:"special"` // "heal" hoặc "" nếu không có kỹ năng đặc biệt
}

type Player struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	EXP      int     `json:"exp"`
	Level    int     `json:"level"`
	Towers   []Tower `json:"towers"`
	Troops   []Troop `json:"troops"`
	Mana     int     `json:"mana"`
}
