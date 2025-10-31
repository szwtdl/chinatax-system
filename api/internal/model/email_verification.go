package model

import "time"

type EmailVerification struct {
	BaseModel
	Username  string    `gorm:"size:50;index:idx_username_email" json:"username"`
	Email     string    `gorm:"size:100" json:"email"`
	Code      string    `gorm:"size:10" json:"code"`
	Status    int       `gorm:"default:0" json:"status"` // 0-未使用,1-已验证,2-过期
	ExpiresAt time.Time `json:"expires_at"`
}
