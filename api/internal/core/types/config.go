package types

import "net/http"

type AppConfig struct {
	Path          string `toml:"-"`
	Listen        string
	MinApp        MinAppConfig
	Session       Session
	Email         EmailConfig
	OcrServer     OcrServer
	NotAuth       NotFilter
	RedisConfig   RedisOption
	Logs          LogOptions
	Gama          GamaConfig
	MysqlDns      string
	StaticDir     string
	StaticUrl     string
	AesEncryptKey string
}

type MinAppConfig struct {
	AppName   string
	AppID     string
	AppKey    string
	AppSecret string
}

type NotFilter struct {
	Paths []string
}
type RedisOption struct {
	Host     string
	Password string
	DBName   int
}

type LogOptions struct {
	Mode       string
	FilePath   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

type GamaConfig struct {
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
}

type EmailConfig struct {
	PlatformName string // 平台昵称，如 "财务平台"
	SMTPHost     string // SMTP 服务器，如 "smtp.qq.com"
	SMTPPort     int    // SMTP 端口，如 465
	Username     string // 发件邮箱
	Password     string // 邮箱授权码
}

type OcrServer struct {
	Host     string // 服务器地址
	Username string // 平台账号
	Password string // 平台密码
}

type SessionDriver string

const (
	SessionDriverMem    = SessionDriver("mem")
	SessionDriverRedis  = SessionDriver("redis")
	SessionDriverCookie = SessionDriver("cookie")
)

type Session struct {
	Driver    SessionDriver
	SecretKey string
	Name      string
	Path      string
	Domain    string
	MaxAge    int
	Secure    bool
	HttpOnly  bool
	SameSite  http.SameSite
}
