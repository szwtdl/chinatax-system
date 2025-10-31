package vo

type Permission struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ParentID  uint   `json:"parent_id"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	Desc      string `json:"desc"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
