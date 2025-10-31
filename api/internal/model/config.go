package model

type Config struct {
	ID     int    `gorm:"primaryKey;autoIncrement;column:id"`
	Key    string `gorm:"column:marker;unique"`
	Config string `gorm:"column:config_json"`
}
