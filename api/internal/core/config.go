package core

import (
	"bytes"
	"github.com/BurntSushi/toml"
	types2 "github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/pkg/utils"
	"net/http"
	"os"
)

func NewDefaultConfig() *types2.AppConfig {
	return &types2.AppConfig{
		Listen:        "0.0.0.0:9977",
		StaticDir:     "./static",
		StaticUrl:     "http://localhost/9977/static",
		AesEncryptKey: utils.RandString(24),
		RedisConfig: types2.RedisOption{
			Host:     "localhost:6379",
			Password: "",
			DBName:   0,
		},
		MinApp: types2.MinAppConfig{
			AppName:   "内测小程序",
			AppID:     "wxc7000dc3da7e2ee5",
			AppKey:    "",
			AppSecret: "5499110ef804c5670b3f95292e91733e",
		},
		Logs: types2.LogOptions{
			FilePath:   "logs/app.log",
			MaxSize:    10,
			MaxBackups: 3,
			MaxAge:     7,
			Mode:       "development", // development or production
		},
		Gama: types2.GamaConfig{
			AppID:     "10000000",
			AppSecret: "970aa04c033bfa8c3be3e019c26af9c0",
		},
		NotAuth: types2.NotFilter{
			Paths: []string{
				"merchant/api/v1/public/login",
				"merchant/api/v1/public/register",
				"merchant/api/v1/public/sendSms",
				"merchant/api/v1/public/sendEmail",
				"merchant/api/v1/public/repass",
				"partner/api/v1/public/login",
				"partner/api/v1/public/register",
				"admin/api/v1/public/login",
				"admin/api/v1/public/logout",
				"admin/api/v1/public/captcha",
			},
		},
		Email: types2.EmailConfig{
			PlatformName: "网通动力",
			SMTPHost:     "smtp.qq.com",
			SMTPPort:     465,
			Username:     "718349985@qq.com",
			Password:     "uacmnntcajambcdf",
		},
		OcrServer: types2.OcrServer{
			Host:     "http://localhost:8000",
			Username: "",
			Password: "",
		},
		Session: types2.Session{
			Driver:    types2.SessionDriverCookie,
			SecretKey: utils.RandString(32),
			Name:      types2.SessionName,
			Domain:    "szwtdl.cn",
			Path:      "/",
			MaxAge:    3600,
			Secure:    true,
			HttpOnly:  true,
			SameSite:  http.SameSiteLaxMode,
		},
	}
}

func LoadConfig(configFile string) (*types2.AppConfig, error) {
	var config *types2.AppConfig
	_, err := os.Stat(configFile)
	if err != nil {
		config = NewDefaultConfig()
		config.Path = configFile
		err := SaveConfig(config)
		if err != nil {
			return nil, err
		}
		return config, nil
	}
	_, err = toml.DecodeFile(configFile, &config)
	if err != nil {
		return nil, err
	}
	return config, err
}

func SaveConfig(config *types2.AppConfig) error {
	buf := new(bytes.Buffer)
	encoder := toml.NewEncoder(buf)
	if err := encoder.Encode(&config); err != nil {
		return err
	}
	return os.WriteFile(config.Path, buf.Bytes(), 0644)
}
