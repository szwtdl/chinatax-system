package model

type Company struct {
	BaseModel
	CompanyName string `json:"company_name"`
	CompanyLogo string `json:"company_logo"`
	Address     string `json:"address"`
	Contact     string `json:"contact"`
	Phone       string `json:"phone"`
}
