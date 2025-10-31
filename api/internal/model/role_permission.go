package model

type RolePermission struct {
	BaseModel
	RoleID       uint `gorm:"not null;index" json:"role_id"`
	PermissionID uint `gorm:"not null;index" json:"permission_id"`
}
