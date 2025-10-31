package vo

type Merchant struct {
	ID           uint   `json:"id"`
	Avatar       string `json:"avatar"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RegisterTime string `json:"register_time"`
	RegisterIP   string `json:"register_ip"`
	LoginIp      string `json:"login_ip"`
	LastTime     string `json:"last_time"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
