package auth

import (
	"github.com/szwtdl/chinatax-system/internal/service"
)

type PartnerPermission struct {
	Service *service.PartnerService
}

func NewPartnerPermission(s *service.PartnerService) *PartnerPermission {
	return &PartnerPermission{
		Service: s,
	}
}

func (p *PartnerPermission) CheckPermission(userID uint, path, method string) bool {
	partner, err := p.Service.GetByID(userID)
	if err != nil || partner == nil {
		return false
	}
	return true
}
