package vo

type TaxCity struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CityName  string `json:"city_name"`
	ClientID  string `json:"client_id"`
	Domain    string `json:"domain"`
	SortNum   int    `json:"sort_num"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CityTax struct {
	Name     string `json:"name"`
	ClientID string `json:"client_id"`
	Domain   string `json:"domain"`
	Status   string `json:"status"`
}

type Mobile struct {
	Mobile      string `json:"mobile"`
	DefaultFlag string `json:"defaultFlag"`
	MobileID    string `json:"mobileId"`
}
