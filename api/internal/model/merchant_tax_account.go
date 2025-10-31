package model

import "time"

type MerchantTaxAccount struct {
	MerchantID   uint       `gorm:"index"`
	TaxAccountID uint       `gorm:"index"`
	Status       string     `gorm:"default:'0'" json:"status"` // 1=绑定
	BindTime     time.Time  `json:"bind_time"`
	Merchant     Merchant   `gorm:"foreignKey:MerchantID"`
	TaxAccount   TaxAccount `gorm:"foreignKey:TaxAccountID"`
}
