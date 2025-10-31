package model

type UserPlatform struct {
	BaseModel
	Platform  string `json:"platform"`
	OpenID    string `gorm:"type:varchar(64);not null;uniqueIndex:idx_platform_openid" json:"openid"` // 平台内唯一
	UnionID   string `gorm:"type:varchar(64);index" json:"unionid,omitempty"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatar_url"`
	Gender    uint8  `gorm:"type:tinyint;default:0" json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
}
