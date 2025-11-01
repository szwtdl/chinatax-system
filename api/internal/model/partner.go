package model

type Partner struct {
	BaseModel
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	Username     string `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password     string `json:"password"`
	Token        string `json:"token"`
	Salt         string `json:"salt"`
	Email        string `json:"email"`
	InviteCode   string `json:"invite_code"`
	LoginCount   int    `gorm:"default:0" json:"login_count"`
	Status       bool   `gorm:"default:true" json:"status"`
	RegisterTime string `json:"register_time"`
	RegisterIP   string `json:"register_ip"`
	LoginIp      string `json:"login_ip"`
	LastTime     string `json:"last_time"`
}
