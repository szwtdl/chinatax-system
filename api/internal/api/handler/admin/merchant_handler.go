package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/szwtdl/chinatax-system/internal/api/handler"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"github.com/szwtdl/chinatax-system/internal/resp"
	"github.com/szwtdl/chinatax-system/internal/service"
	"github.com/szwtdl/chinatax-system/internal/vo"
	"github.com/szwtdl/chinatax-system/pkg/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MerchantHandler struct {
	handler.BaseHandler
	merchantService *service.MerchantService
}

func NewMerchantHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *MerchantHandler {
	repo := repository.NewBaseRepository[model.Merchant]()
	merchantService := service.NewMerchantService(db, repo)
	return &MerchantHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		merchantService: merchantService,
	}
}

func (h *MerchantHandler) List(c *gin.Context) {
	var req struct {
		Keyword string `json:"keyword" binding:"required"`
		Filter  string `json:"filter,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "参数错误")
		return
	}
	page := h.GetInt(c, "page", 1)
	limit := h.GetInt(c, "limit", 10)
	var lists = make([]vo.Merchant, 0)
	items, total, err := h.merchantService.Pagination(map[string]interface{}{}, page, limit, "id DESC")
	if err != nil {
		resp.SUCCESS(c, "获取失败")
		return
	}
	for _, item := range items {
		lists = append(lists, vo.Merchant{
			ID:           item.ID,
			Avatar:       item.Avatar,
			Nickname:     item.Nickname,
			Username:     item.Username,
			Token:        item.Token,
			RegisterTime: item.RegisterTime,
			RegisterIP:   item.RegisterIP,
			CreatedAt:    utils.Stamp2str(item.CreatedAt.Unix()),
			UpdatedAt:    utils.Stamp2str(item.UpdatedAt.Unix()),
		})
	}
	pageVo := vo.NewPage(total, page, limit, lists)
	resp.SUCCESS(c, pageVo)
}

func (h *MerchantHandler) Update(c *gin.Context) {}

func (h *MerchantHandler) Delete(c *gin.Context) {}
