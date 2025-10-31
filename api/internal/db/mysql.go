package db

import (
	"fmt"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func NewGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:         "china_",
			SingularTable:       false,
			IdentifierMaxLength: 64,
		},
	}
}

func NewMysql(config *gorm.Config, appConfig *types.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(appConfig.MysqlDns), config)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(32)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping DB failed: %w", err)
	}
	models := []interface{}{
		&model.Config{},
		&model.Admin{},
		&model.Role{},
		&model.AdminRole{},
		&model.Permission{},
		&model.RolePermission{},
		&model.Merchant{},
		&model.MerchantPlatform{},
		&model.Partner{},
		&model.CityTax{},
		&model.MerchantTaxAccount{},
		&model.EmailVerification{},
	}
	for _, modelName := range models {
		if err := db.AutoMigrate(modelName); err != nil {
			panic(fmt.Sprintf("Failed to migrate model %T: %v", modelName, err))
		}
	}
	var count int64
	db.Model(&model.Admin{}).Count(&count)
	if count == 0 {
		salt := utils.RandString(16)
		password, _ := utils.HashPassword("Pj123456", salt)
		defaultAdmin := model.Admin{
			Avatar:   "static/images/default.jpg",
			Nickname: "超级管理员",
			Username: "admin",
			Password: password, // 建议使用加密
			Email:    "szpengjian@gmail.com",
			Salt:     salt,
			Status:   true,
			Super:    true,
		}
		if err := db.Create(&defaultAdmin).Error; err != nil {
			return nil, fmt.Errorf("failed to create default admin: %w", err)
		}
		fmt.Println("✅ 默认管理员账号已创建：用户名 admin / 密码 admin123")
	}
	return db, nil
}
