package service

import (
	"fmt"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"gorm.io/gorm"
)

type RoleService struct {
	BaseService[model.Role]
}

func NewRoleService(db *gorm.DB, repo repository.BaseRepository[model.Role]) *RoleService {
	return &RoleService{
		BaseService: BaseService[model.Role]{
			repo: repo,
			db:   db,
		},
	}
}

func (s *RoleService) validatePermissions(tx *gorm.DB, ids []uint) error {
	if len(ids) == 0 {
		return fmt.Errorf("权限不能为空")
	}
	var count int64
	if err := tx.Model(&model.Permission{}).Where("id IN ?", ids).Count(&count).Error; err != nil {
		return err
	}
	if count != int64(len(ids)) {
		return fmt.Errorf("部分权限ID不存在")
	}
	return nil
}

func (s *RoleService) CreateWithPermissions(role *model.Role, permIDs []uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.validatePermissions(tx, permIDs); err != nil {
			return err
		}

		if err := tx.Create(role).Error; err != nil {
			return err
		}

		relations := make([]model.RolePermission, 0, len(permIDs))
		for _, pid := range permIDs {
			relations = append(relations, model.RolePermission{
				RoleID:       role.ID,
				PermissionID: pid,
			})
		}
		return tx.Create(&relations).Error
	})
}

func (s *RoleService) UpdateWithPermissions(role *model.Role, permIDs []uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.validatePermissions(tx, permIDs); err != nil {
			return err
		}

		if err := tx.Save(role).Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Where("role_id = ?", role.ID).Delete(&model.RolePermission{}).Error; err != nil {
			return err
		}

		relations := make([]model.RolePermission, 0, len(permIDs))
		for _, pid := range permIDs {
			relations = append(relations, model.RolePermission{
				RoleID:       role.ID,
				PermissionID: pid,
			})
		}
		return tx.Create(&relations).Error
	})
}

func (s *RoleService) DeleteWithRelations(roleID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("role_id = ?", roleID).Delete(&model.RolePermission{}).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("role_id = ?", roleID).Delete(&model.AdminRole{}).Error; err != nil {
			return err
		}
		return tx.Unscoped().Delete(&model.Role{}, roleID).Error
	})
}
