package model

type AdminRole struct {
	BaseModel
	AdminID uint `gorm:"not null;index" json:"admin_id"`
	RoleID  uint `gorm:"not null;index" json:"role_id"`
}
