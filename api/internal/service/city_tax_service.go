package service

import (
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"gorm.io/gorm"
)

type CityTaxService struct {
	BaseService[model.CityTax]
}

func NewCityTaxService(db *gorm.DB, repo repository.BaseRepository[model.CityTax]) *CityTaxService {
	return &CityTaxService{
		BaseService: BaseService[model.CityTax]{
			repo: repo,
			db:   db,
		},
	}
}

func (h *CityTaxService) GetClient(clientID string) (*model.CityTax, error) {
	return h.repo.FindOne(h.db, map[string]interface{}{"client_id": clientID, "status": "0"})
}

func (h *CityTaxService) GetCityName(cityName string) (*model.CityTax, error) {
	return h.repo.FindOne(h.db, map[string]interface{}{"city_name": cityName})
}
