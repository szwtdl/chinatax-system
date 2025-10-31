package vo

type Partner struct {
	ID           uint   `json:"id"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RegisterTime string `json:"register_time"`
	RegisterIP   string `json:"register_ip"`
	LoginIp      string `json:"login_ip"`
	LastTime     string `json:"last_time"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
