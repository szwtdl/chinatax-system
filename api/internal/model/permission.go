package model

type Permission struct {
	BaseModel
	ParentID    uint   `gorm:"default:0;index" json:"parent_id"`
	Name        string `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Path        string `gorm:"size:255" json:"path"`
	Method      string `gorm:"size:10" json:"method"`
	Description string `gorm:"size:255" json:"description"`
}
