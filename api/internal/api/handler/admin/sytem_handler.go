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
	db            *gorm.DB
	configService *service.ConfigService
}

func NewSystemHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *SystemHandler {
	repo := repository.NewBaseRepository[model.Config]()
	return &SystemHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		db:            db,
		configService: service.NewConfigService(db, repo),
	}
}

func (h *SystemHandler) GetValue(c *gin.Context) {
	key := c.Query("key")
	var config model.Config
	res := h.db.Where("marker", key).First(&config)
	if res.Error != nil {
		resp.ERROR(c, res.Error.Error())
		return
	}
	var m map[string]interface{}
	err := utils.JsonDecode(config.Config, &m)
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	resp.SUCCESS(c, m)
}

func (h *SystemHandler) SetValue(c *gin.Context) {
	var data struct {
		Key    string                 `json:"key"`
		Config map[string]interface{} `json:"config"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	str := utils.JsonEncode(&data.Config)
	config := model.Config{Key: data.Key, Config: str}
	res := h.db.FirstOrCreate(&config, model.Config{Key: data.Key})
	if res.Error != nil {
		resp.ERROR(c, res.Error.Error())
		return
	}
	if config.ID > 0 {
		config.Config = str
		res = h.db.Updates(&config)
		if res.Error != nil {
			resp.ERROR(c, res.Error.Error())
			return
		}
		var cfg model.Config
		h.db.Where("marker", data.Key).First(&cfg)
		var err error
		if err != nil {
			resp.ERROR(c, "Failed to update config cache: "+err.Error())
			return
		}
	}
	resp.SUCCESS(c, config)
}
