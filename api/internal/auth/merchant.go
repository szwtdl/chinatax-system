package auth

import (
	"github.com/szwtdl/chinatax-system/internal/service"
)

type MerchantPermission struct {
	Service *service.MerchantService
}

func NewMerchantPermission(s *service.MerchantService) *MerchantPermission {
	return &MerchantPermission{
		Service: s,
	}
}

func (m *MerchantPermission) CheckPermission(userID uint, path, method string) bool {
	merchant, err := m.Service.GetByID(userID)
	if err != nil || merchant == nil {
		return false
	}
	return true
}
