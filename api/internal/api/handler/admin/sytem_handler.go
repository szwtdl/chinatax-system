package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/szwtdl/chinatax-system/internal/api/handler"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"github.com/szwtdl/chinatax-system/internal/resp"
	"github.com/szwtdl/chinatax-system/internal/service"
	"github.com/szwtdl/chinatax-system/pkg/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SystemHandler struct {
	handler.BaseHandler
	configService *service.ConfigService
}

func NewSystemHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *SystemHandler {
	repo := repository.NewBaseRepository[model.Config]()
	return &SystemHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		configService: service.NewConfigService(db, repo),
	}
}

func (h *SystemHandler) GetValue(c *gin.Context) {
	var req struct {
		Marker string `json:"key"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	config, err := h.configService.GetSystem(req.Marker)
	if err != nil {
		resp.ERROR(c, "没有数据")
		return
	}
	resp.SUCCESS(c, utils.JsonDecode(config.Config, map[string]interface{}{
		"name": "xx",
	}))
}

func (h *SystemHandler) SetValue(c *gin.Context) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}

	resp.SUCCESS(c, req)
}
