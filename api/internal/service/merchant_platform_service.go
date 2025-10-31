package service

import (
	"errors"
	"fmt"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"github.com/szwtdl/chinatax-system/pkg/utils"
	"gorm.io/gorm"
	"strings"
	"time"
)

type MerchantPlatformService struct {
	BaseService[model.MerchantPlatform]
}

func NewMerchantPlatformService(db *gorm.DB, repo repository.BaseRepository[model.MerchantPlatform]) *MerchantPlatformService {
	return &MerchantPlatformService{
		BaseService[model.MerchantPlatform]{
			repo: repo,
			db:   db,
		},
	}
}

func (s *MerchantPlatformService) GetByPlatform(userID, platform string) (*model.MerchantPlatform, error) {
	return s.repo.FindOne(s.db, map[string]interface{}{"platform": platform, "platform_uid": userID})
}

func (s *MerchantPlatformService) GetOrCreatePlatform(
	userMobile, platform, platformUID, nickname, avatar, clientIP string,
) (*model.MerchantPlatform, *model.Merchant, error) {

	var platformAccount model.MerchantPlatform
	var user model.Merchant

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// 1. 查询平台账号
		if err := tx.Where("platform = ? AND platform_uid = ?", platform, platformUID).First(&platformAccount).Error; err == nil {
			// 平台账号存在，获取关联用户
			if err := tx.First(&user, platformAccount.MerchantID).Error; err != nil {
				return err
			}
			return nil
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// 2. 平台账号不存在，检查手机号 Merchant
		if err := tx.Where("username = ?", userMobile).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户不存在 → 创建
			salt, _ := utils.GenerateSalt(16)
			password := utils.RandString(8)
			hashedPwd, _ := utils.HashPassword(password, salt)
			userPtr := &model.Merchant{
				Username:     userMobile,
				Password:     hashedPwd,
				Salt:         salt,
				IsPassword:   false,
				RegisterTime: time.Now().Format("2006-01-02 15:04:05"),
				RegisterIP:   clientIP,
			}
			if err := tx.Create(userPtr).Error; err != nil {
				return err
			}
			user = *userPtr
		} else if err != nil {
			return err
		}

		// 3. 创建平台账号
		platformAccountPtr := &model.MerchantPlatform{
			MerchantID:  user.ID,
			Platform:    platform,
			PlatformUID: platformUID,
			Nickname:    nickname,
			Avatar:      avatar,
		}
		if err := tx.Create(platformAccountPtr).Error; err != nil {
			if strings.Contains(err.Error(), "Duplicate") || strings.Contains(err.Error(), "UNIQUE") {
				if err := tx.Where("platform = ? AND platform_uid = ?", platform, platformUID).First(&platformAccount).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			platformAccount = *platformAccountPtr
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return &platformAccount, &user, nil
}

func (s *MerchantPlatformService) GetOrCreate(mp *model.MerchantPlatform) (*model.MerchantPlatform, error) {
	var result *model.MerchantPlatform
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// 先查询
		existing, err := s.repo.FindOne(tx, map[string]interface{}{
			"platform":     mp.Platform,
			"platform_uid": mp.PlatformUID,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if existing != nil {
			// 存在直接返回
			result = existing
			return nil
		}

		// 不存在则创建
		if err := tx.Create(mp).Error; err != nil {
			return err
		}
		result = mp
		return nil
	})
	return result, err
}

func (s *MerchantPlatformService) CreateAccount(mp *model.MerchantPlatform) error {
	// 使用事务，确保安全
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 检查平台账号是否已存在
		existing, err := s.repo.FindOne(tx, map[string]interface{}{
			"platform":     mp.Platform,
			"platform_uid": mp.PlatformUID,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if existing != nil {
			return fmt.Errorf("平台账号已存在: %s", mp.PlatformUID)
		}

		// 插入新账号
		if err := tx.Create(mp).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *MerchantPlatformService) UpdateAccount(mp *model.MerchantPlatform) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 查询原始记录
		existing, err := s.repo.FindOne(tx, map[string]interface{}{
			"platform":     mp.Platform,
			"platform_uid": mp.PlatformUID,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("平台账号不存在: %s", mp.PlatformUID)
			}
			return err
		}

		// 更新字段
		existing.MerchantID = mp.MerchantID
		existing.Nickname = mp.Nickname
		existing.Avatar = mp.Avatar
		existing.UpdatedAt = time.Now()
		if err := tx.Save(existing).Error; err != nil {
			return err
		}
		return nil
	})
}
