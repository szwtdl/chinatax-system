package model

type CityTax struct {
	BaseModel
	Name        string `gorm:"index:name,unique;not null" json:"name"`
	ClientID    string `gorm:"index:client_id,unique;not null" json:"client_id"`
	CityName    string `json:"city_name"`
	Description string `json:"description"`
	IsDefault   string `gorm:"default:'0'" json:"default"` // 默认值为 "0"
	Status      string `gorm:"default:'0'" json:"status"`  // 状态0
}
