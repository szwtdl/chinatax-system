package vo

type Admin struct {
	ID           uint   `json:"id"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Status       bool   `json:"status"`
	Super        bool   `json:"super"`
	RegisterTime string `json:"register_time"`
	RegisterIp   string `json:"register_ip"`
	LoginIP      string `json:"login_ip"`
	LastTime     string `json:"last_time"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
