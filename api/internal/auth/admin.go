package auth

import (
	"github.com/szwtdl/chinatax-system/internal/service"
)

type AdminPermission struct {
	Service *service.AdminService
}

func NewAdminPermission(adminService *service.AdminService) *AdminPermission {
	return &AdminPermission{
		Service: adminService,
	}
}

func (a *AdminPermission) CheckPermission(userID uint, path, method string) bool {
	ok, err := a.Service.HasPermission(userID, path, method)
	return err == nil && ok
}
