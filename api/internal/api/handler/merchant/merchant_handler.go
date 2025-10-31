package merchant

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

type MerchantHandler struct {
	handler.BaseHandler
	merchantService *service.MerchantService
}

func NewMerchantHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *MerchantHandler {
	repo := repository.NewBaseRepository[model.Merchant]()
	return &MerchantHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		merchantService: service.NewMerchantService(db, repo),
	}
}

func (h *MerchantHandler) Info(c *gin.Context) {
	resp.ERROR(c, types.BizMsg[types.Success])
}

func (h *MerchantHandler) Update(c *gin.Context) {
	resp.ERROR(c, types.BizMsg[types.Success])
}

func (h *MerchantHandler) RePass(c *gin.Context) {
	resp.ERROR(c, types.BizMsg[types.Success])
}
