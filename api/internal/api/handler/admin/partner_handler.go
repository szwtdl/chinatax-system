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

type PartnerHandler struct {
	handler.BaseHandler
	partnerService *service.PartnerService
}

func NewPartnerHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *PartnerHandler {
	repo := repository.NewBaseRepository[model.Partner]()
	partnerService := service.NewPartnerService(db, repo)
	return &PartnerHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		partnerService: partnerService,
	}
}

func (h *PartnerHandler) List(c *gin.Context) {
	page := h.GetInt(c, "page", 1)
	limit := h.GetInt(c, "limit", 10)
	keyword := c.Query("keywords")
	filter := c.Query("filter")
	where := map[string]interface{}{}
	if keyword != "" && filter != "" {
		allowedFilters := map[string]bool{"username": true, "nickname": true, "email": true} // 允许的字段
		if allowedFilters[filter] {
			where[filter+" LIKE ?"] = "%" + utils.EscapeLikeKeyword(keyword) + "%"
		}
	}
	var lists = make([]vo.Partner, 0)
	items, total, err := h.partnerService.Pagination(where, page, limit, "id DESC")
	if err != nil {
		resp.SUCCESS(c, "获取失败")
		return
	}
	for _, item := range items {
		lists = append(lists, vo.Partner{
			ID:           item.ID,
			Avatar:       item.Avatar,
			Nickname:     item.Nickname,
			Username:     item.Username,
			Token:        item.Token,
			RegisterTime: item.RegisterTime,
			RegisterIP:   item.RegisterIP,
			LastTime:     item.LastTime,
			LoginIp:      item.LoginIp,
			Status:       item.Status,
			LoginCount:   item.LoginCount,
			CreatedAt:    utils.Stamp2str(item.CreatedAt.Unix()),
			UpdatedAt:    utils.Stamp2str(item.UpdatedAt.Unix()),
		})
	}
	pageVo := vo.NewPage(total, page, limit, lists)
	resp.SUCCESS(c, pageVo)
}

func (h *PartnerHandler) Create(c *gin.Context) {
	var req struct {
		Nickname string `json:"nickname" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}

}

func (h *PartnerHandler) Update(c *gin.Context) {}

func (h *PartnerHandler) Delete(c *gin.Context) {}
