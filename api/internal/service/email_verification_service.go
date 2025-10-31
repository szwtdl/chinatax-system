package service

import (
	"errors"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type EmailVerificationService struct {
	db *gorm.DB
}

func NewEmailVerificationService(db *gorm.DB) *EmailVerificationService {
	return &EmailVerificationService{
		db: db,
	}
}

func (s *EmailVerificationService) GenerateCode(username, email string, expireMinutes int, intervalSeconds int) (string, error) {
	var last model.EmailVerification
	// 查询最后一条未使用且未过期的验证码
	err := s.db.
		Where("username = ? AND email = ? AND status = 0 AND expires_at > ?", username, email, time.Now()).
		Order("created_at DESC").
		First(&last).Error
	// 如果存在，检查发送间隔
	if err == nil {
		if time.Since(last.CreatedAt) < time.Duration(intervalSeconds)*time.Second {
			return "", errors.New("验证码发送过于频繁，请稍后再试")
		}
	}
	// 生成新验证码
	code := utils.GenerateRandomCode(6)
	verification := model.EmailVerification{
		Username:  username,
		Email:     email,
		Code:      code,
		Status:    0, // 未使用
		ExpiresAt: time.Now().Add(time.Duration(expireMinutes) * time.Minute),
	}
	if err := s.db.Create(&verification).Error; err != nil {
		return "", err
	}

	return code, nil
}

func (s *EmailVerificationService) VerifyCode(username, email, code string) (bool, error) {
	var record model.EmailVerification
	err := s.db.
		Where("username = ? AND email = ? AND code = ? AND status = 0 AND expires_at > ?", username, email, code, time.Now()).
		First(&record).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("验证码无效或已过期")
		}
		return false, err
	}

	// 更新状态为已验证
	record.Status = 1
	if err := s.db.Save(&record).Error; err != nil {
		return false, err
	}

	return true, nil
}
