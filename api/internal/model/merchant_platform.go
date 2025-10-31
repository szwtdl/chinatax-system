package model

type MerchantPlatform struct {
	BaseModel
	MerchantID  uint   `json:"merchant_id"`
	Platform    string `json:"platform" gorm:"size:50;not null;index:idx_platform_uid,unique"`
	PlatformUID string `json:"platform_uid" gorm:"size:100;not null;index:idx_platform_uid,unique"`
	Nickname    string `json:"nickname" gorm:"size:100"`
	Avatar      string `json:"avatar" gorm:"size:255"`
}
