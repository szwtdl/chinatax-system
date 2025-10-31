package model

type TaxAccount struct {
	BaseModel
	UserID       uint                 `json:"user_id"`
	Username     string               `json:"username"`
	FullName     string               `json:"full_name"`
	UserType     string               `json:"user_type"`
	IDCard       string               `json:"id_card"`
	ClientID     string               `json:"client_id"`
	ClientType   string               `json:"client_type"`
	TaxID        string               `json:"tax_id"`
	Phone        string               `json:"phone"`
	Mobile       string               `json:"mobile"`
	Gender       string               `json:"gender"`
	ExpiresIn    int                  `json:"expires_in"`
	TaxerCode    string               `json:"taxer_code"`
	TaxerName    string               `json:"taxer_name"`
	AreaPrefix   string               `json:"area_prefix"`
	AreaName     string               `json:"area_name"`
	AreaPreName  string               `json:"area_pre_name"`
	RegisterTime string               `json:"register_time"`
	Bindings     []MerchantTaxAccount `gorm:"foreignKey:TaxAccountID" json:"bindings"`
}
