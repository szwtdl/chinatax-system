package model

type Role struct {
	BaseModel
	Name        string `gorm:"size:50;uniqueIndex;not null" json:"name"`
	DisplayName string `gorm:"size:100" json:"display_name"`
	Description string `gorm:"size:255" json:"description"`
	Status      int    `gorm:"default:1" json:"status"`
}
