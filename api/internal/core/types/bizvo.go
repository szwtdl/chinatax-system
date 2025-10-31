package types

type BizVo struct {
	Code     BizCode     `json:"code"`
	PageNum  int         `json:"page_num,omitempty"`
	PageSize int         `json:"page_size,omitempty"`
	Total    int64       `json:"total,omitempty"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
