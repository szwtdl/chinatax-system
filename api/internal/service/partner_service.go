package service

import (
	"errors"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"gorm.io/gorm"
)

type PartnerService struct {
	BaseService[model.Partner]
}

func NewPartnerService(db *gorm.DB, repo repository.BaseRepository[model.Partner]) *PartnerService {
	return &PartnerService{
		BaseService: BaseService[model.Partner]{
			Repo: repo,
			db:   db,
		},
	}
}

func (s *PartnerService) GetByUsername(username string) (*model.Partner, error) {
	return s.Repo.FindOne(s.db, map[string]interface{}{"username": username})
}

func (s *PartnerService) GetToken(token string) (*model.Partner, error) {
	return s.Repo.FindOne(s.db, map[string]interface{}{"token": token})
}

func (s *PartnerService) Update(id uint, fields map[string]interface{}) error {
	rows, err := s.Repo.UpdateFields(s.db, map[string]interface{}{"id": id}, fields)
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("未找到要更新的记录")
	}
	return nil
}
