package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/szwtdl/chinatax-system/internal/api/handler"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/internal/resp"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DashboardHandler struct {
	handler.BaseHandler
}

func NewDashboardHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *DashboardHandler {
	return &DashboardHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
	}
}

func (h *DashboardHandler) Welcome(c *gin.Context) {
	resp.SUCCESS(c, map[string]interface{}{
		"order":    100,
		"merchant": 100,
		"partner":  100,
	})
}
