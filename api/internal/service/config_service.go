package service

import (
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"gorm.io/gorm"
)

type ConfigService struct {
	BaseService[model.Config]
}

func NewConfigService(db *gorm.DB, repo repository.BaseRepository[model.Config]) *ConfigService {
	return &ConfigService{
		BaseService[model.Config]{
			Repo: repo,
			db:   db,
		},
	}
}

func (s *ConfigService) GetSystem(name string) (*model.Config, error) {
	return s.Repo.FindOne(s.db, map[string]interface{}{"marker": name})
}
