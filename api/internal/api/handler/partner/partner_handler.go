package partner

import (
	"github.com/gin-gonic/gin"
	"github.com/szwtdl/chinatax-system/internal/api/handler"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/internal/resp"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PartnerHandler struct {
	handler.BaseHandler
	db *gorm.DB
}

func NewPartnerHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *PartnerHandler {
	return &PartnerHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		db: db,
	}
}

func (h *PartnerHandler) Info(c *gin.Context) {
	resp.SUCCESS(c, types.BizMsg[types.Success])
}

func (h *PartnerHandler) Update(c *gin.Context) {
	resp.SUCCESS(c, types.BizMsg[types.Success])
}

func (h *PartnerHandler) RePass(c *gin.Context) {
	resp.SUCCESS(c, types.BizMsg[types.Success])
}
