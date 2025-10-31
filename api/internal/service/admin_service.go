package service

import (
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"gorm.io/gorm"
	"strings"
)

type AdminService struct {
	BaseService[model.Admin]
}

func NewAdminService(db *gorm.DB, repo repository.BaseRepository[model.Admin]) *AdminService {
	return &AdminService{
		BaseService: BaseService[model.Admin]{
			repo: repo,
			db:   db,
		},
	}
}

func (s *AdminService) GetByUsername(username string) (*model.Admin, error) {
	return s.repo.FindOne(s.db, map[string]interface{}{"username": username})
}

func (s *AdminService) CreateWithRoles(admin *model.Admin, roleIDs []uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(admin).Error; err != nil {
			return err
		}

		if len(roleIDs) > 0 {
			relations := make([]model.AdminRole, 0, len(roleIDs))
			for _, rid := range roleIDs {
				relations = append(relations, model.AdminRole{
					AdminID: admin.ID,
					RoleID:  rid,
				})
			}
			if err := tx.Create(&relations).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *AdminService) UpdateWithRoles(admin *model.Admin, roleIDs []uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 先更新基础信息
		if err := tx.Save(admin).Error; err != nil {
			return err
		}

		// 删除旧关联
		if err := tx.Where("admin_id = ?", admin.ID).Delete(&model.AdminRole{}).Error; err != nil {
			return err
		}

		// 添加新关联
		if len(roleIDs) > 0 {
			relations := make([]model.AdminRole, 0, len(roleIDs))
			for _, rid := range roleIDs {
				relations = append(relations, model.AdminRole{
					AdminID: admin.ID,
					RoleID:  rid,
				})
			}
			if err := tx.Create(&relations).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *AdminService) DeleteWithRoles(adminID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 删除角色关联
		if err := tx.Where("admin_id = ?", adminID).Delete(&model.AdminRole{}).Error; err != nil {
			return err
		}
		// 删除管理员
		if err := tx.Delete(&model.Admin{}, adminID).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *AdminService) HasPermission(userID uint, path, method string) (bool, error) {
	admin, err := s.repo.FindOne(s.db, map[string]interface{}{"id": userID})
	if err != nil {
		return false, err
	}
	// 超级管理员直接放行
	if admin.Super {
		return true, nil
	}

	// 查询角色
	var adminRoles []model.AdminRole
	if err = s.db.Where("admin_id = ?", admin.ID).Find(&adminRoles).Error; err != nil {
		return false, err
	}
	if len(adminRoles) == 0 {
		return false, nil
	}

	roleIDs := make([]uint, 0, len(adminRoles))
	for _, ar := range adminRoles {
		roleIDs = append(roleIDs, ar.RoleID)
	}

	// 查询角色对应权限ID
	var rolePerms []model.RolePermission
	if err = s.db.Where("role_id IN ?", roleIDs).Find(&rolePerms).Error; err != nil {
		return false, err
	}
	if len(rolePerms) == 0 {
		return false, nil
	}

	permIDs := make([]uint, 0, len(rolePerms))
	for _, rp := range rolePerms {
		permIDs = append(permIDs, rp.PermissionID)
	}

	// 查询权限表并匹配
	var perms []model.Permission
	if err = s.db.Where("id IN ?", permIDs).Find(&perms).Error; err != nil {
		return false, err
	}

	for _, p := range perms {
		if strings.EqualFold(p.Path, path) && strings.EqualFold(p.Method, method) {
			return true, nil
		}
	}

	return false, nil
}
