package service

import (
	"errors"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"gorm.io/gorm"
)

type PermissionService struct {
	BaseService[model.Permission]
}

func NewPermissionService(db *gorm.DB, repo repository.BaseRepository[model.Permission]) *PermissionService {
	return &PermissionService{
		BaseService: BaseService[model.Permission]{
			repo: repo,
			db:   db,
		},
	}
}

func (s *PermissionService) FindOne(cond map[string]interface{}, preloads ...string) (*model.Permission, error) {
	return s.repo.FindOne(s.db, cond, preloads...)
}

func (s *PermissionService) CreatePermission(p *model.Permission) error {
	if p.Name == "" {
		return errors.New("权限名称不能为空")
	}
	var count int64
	if err := s.db.Model(&model.Permission{}).Where("name = ?", p.Name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("权限名称已存在")
	}
	return s.repo.Create(s.db, p)
}

func (s *PermissionService) UpdatePermission(p *model.Permission) error {
	if p.ID == 0 {
		return errors.New("权限ID不能为空")
	}

	// 可选：检查名称是否重复
	var count int64
	if err := s.db.Model(&model.Permission{}).Where("name = ? AND id <> ?", p.Name, p.ID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("权限名称已存在")
	}

	return s.repo.Update(s.db, p)
}

func (s *PermissionService) DeletePermission(id uint) error {
	if id == 0 {
		return errors.New("权限ID不能为空")
	}
	return s.repo.Delete(s.db, map[string]interface{}{"id": id})
}
