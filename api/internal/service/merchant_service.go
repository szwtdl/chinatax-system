package service

import (
	"errors"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"gorm.io/gorm"
)

type MerchantService struct {
	BaseService[model.Merchant]
}

func NewMerchantService(db *gorm.DB, repo repository.BaseRepository[model.Merchant]) *MerchantService {
	return &MerchantService{
		BaseService: BaseService[model.Merchant]{
			repo: repo,
			db:   db,
		},
	}
}

func (s *MerchantService) GetByUsername(username string) (*model.Merchant, error) {
	return s.repo.FindOne(s.db, map[string]interface{}{"username": username})
}

func (s *MerchantService) GetToken(token string) (*model.Merchant, error) {
	return s.repo.FindOne(s.db, map[string]interface{}{"token": token})
}

func (s *MerchantService) Update(id uint, fields map[string]interface{}) error {
	rows, err := s.repo.UpdateFields(s.db, map[string]interface{}{"id": id}, fields)
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("未找到要更新的记录")
	}
	return nil
}
