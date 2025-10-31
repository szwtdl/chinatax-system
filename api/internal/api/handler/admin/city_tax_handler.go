package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/szwtdl/chinatax-system/internal/api/handler"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"github.com/szwtdl/chinatax-system/internal/resp"
	"github.com/szwtdl/chinatax-system/internal/service"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CityTaxHandler struct {
	handler.BaseHandler
	cityTaxService *service.CityTaxService
}

func NewCityTaxHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *CityTaxHandler {
	Repo := repository.NewBaseRepository[model.CityTax]()
	CityTax := service.NewCityTaxService(db, Repo)
	return &CityTaxHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		cityTaxService: CityTax,
	}
}

func (h *CityTaxHandler) List(c *gin.Context) {
	var req struct {
		Keyword string `json:"keyword" binding:"required"`
		Filter  string `json:"filter,omitempty"`
		Page    int    `json:"page,omitempty"`
		Limit   int    `json:"limit,omitempty"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.ERROR(c, "参数错误")
		return
	}
	resp.SUCCESS(c, "获取成功")
}

func (h *CityTaxHandler) Create(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		ClientID    string `json:"client_id" binding:"required,len=32"`
		CityName    string `json:"city_name" binding:"required"`
		Description string `json:"description,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "参数不能为空")
		return
	}
	if exist, _ := h.cityTaxService.GetClient(req.ClientID); exist != nil {
		resp.ERROR(c, "城市已经存在")
		return
	}
	if exist, _ := h.cityTaxService.GetCityName(req.CityName); exist != nil {
		resp.ERROR(c, "城市CityName已经存在")
		return
	}
	cityTax := &model.CityTax{
		Name:        req.Name,
		ClientID:    req.ClientID,
		CityName:    req.CityName,
		Description: req.Description,
	}
	err := h.cityTaxService.Create(cityTax)
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	resp.SUCCESS(c, "添加成功")
}

func (h *CityTaxHandler) Update(c *gin.Context) {
	var req struct {
		ID          uint   `json:"id" binding:"required"`
		Name        string `json:"name" binding:"required"`
		ClientID    string `json:"client_id" binding:"required,len=32"`
		CityName    string `json:"city_name" binding:"required"`
		Status      string `json:"status" binding:"required,oneof=0 1"`
		Description string `json:"description,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "参数错误")
		return
	}
	client, err := h.cityTaxService.GetByID(req.ID)
	if err != nil {
		resp.ERROR(c, "不存在")
		return
	}
	client.Name = req.Name
	client.ClientID = req.ClientID
	client.CityName = req.CityName
	client.Description = req.Description
	client.Status = req.Status
	if err = h.cityTaxService.Update(client); err != nil {
		resp.ERROR(c, "更新失败")
		return
	}
	resp.SUCCESS(c, "更新成功")
}

func (h *CityTaxHandler) Delete(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	_, err := h.cityTaxService.GetByID(req.ID)
	if err != nil {
		resp.ERROR(c, "城市不存在")
		return
	}
	if err = h.cityTaxService.DeleteByID(req.ID); err != nil {
		resp.ERROR(c, "删除失败")
		return
	}
	resp.SUCCESS(c, "删除成功")
}
