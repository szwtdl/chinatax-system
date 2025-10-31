package types

type BizCode int

const (
	// 通用状态码
	Success       BizCode = 0
	Failed        BizCode = 1
	NotAuthorized BizCode = 400

	// 自定义错误码可以在这里扩展
	InvalidParam  BizCode = 1001
	DatabaseError BizCode = 1002
	AuthError     BizCode = 1003
)

var BizMsg = map[BizCode]string{
	Success:       "Success",
	Failed:        "系统开小差了",
	NotAuthorized: "未授权访问",
	InvalidParam:  "非法参数或参数解析失败",
	DatabaseError: "数据库操作失败",
	AuthError:     "认证失败",
}
