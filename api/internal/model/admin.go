package model

type Admin struct {
	BaseModel
	Username     string `gorm:"index:username,unique"`
	Email        string `gorm:"index:email,unique"`
	Password     string
	Nickname     string
	Avatar       string
	Token        string `gorm:"index:token,unique"`
	Salt         string
	Status       bool   `gorm:"default:true"`
	Super        bool   `gorm:"default:false"`
	RegisterTime string `json:"register_time"`
	RegisterIP   string `json:"register_ip"`
	LoginIp      string `json:"login_ip"`
	LastTime     string `json:"last_time"`
}
