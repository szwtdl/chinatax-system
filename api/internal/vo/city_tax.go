package vo

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
