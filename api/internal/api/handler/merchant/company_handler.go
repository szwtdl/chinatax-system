package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/szwtdl/chinatax-system/internal/api/handler"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/internal/resp"
	"github.com/szwtdl/chinatax/service"
	"github.com/szwtdl/chinatax_api/chinatax"
	"github.com/szwtdl/chinatax_api/platforms"
	"github.com/szwtdl/chinatax_api/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CompanyHandler struct {
	handler.BaseHandler
	db        *gorm.DB
	ApiClient chinatax.ApiClient
}

func NewCompanyHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *CompanyHandler {
	return &CompanyHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		db: db,
		ApiClient: platforms.NewClient(map[string]string{
			"app_id":    "test",
			"secret":    "test",
			"debug":     "false",
			"city_name": "shenzhen",
		}),
	}
}

func (h *CompanyHandler) Search(c *gin.Context) {
	var req struct {
		Keywords string `json:"keywords" binding:"required"`
		Type     string `json:"type" binding:"required,oneof=1 2"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "参数不能为空")
		return
	}
	searchType := service.Gama
	token := h.App.Gama.AppSecret
	if req.Type == "2" {
		searchType = service.Tianyancha
		token = "" // 天眼查可能不需要 token
	}
	company, err := h.ApiClient.CompanyList(request.CompanyListRequest{
		Keyword: req.Keywords,
		Type:    searchType,
		Token:   token,
	})
	if err != nil {
		h.Log.Error(err)
		resp.ERROR(c, "搜索失败:")
		return
	}
	if company.Code != 0 {
		resp.ERROR(c, company.Msg)
		return
	}
	resp.SUCCESS(c, company.Data)
}

func (h *CompanyHandler) Detail(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	response, err := h.ApiClient.CompanyDetail(request.CompanyDetailRequest{
		Name:  req.Name,
		Token: h.App.Gama.AppSecret,
	})
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	if response.Code != 0 {
		resp.ERROR(c, response.Msg)
		return
	}
	resp.SUCCESS(c, response.Data)
}

func (h *CompanyHandler) Ocr(c *gin.Context) {
	var req struct {
		Image string `json:"image" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	ocr, err := h.ApiClient.CompanyOcr(request.CompanyOcrRequest{
		Image: req.Image,
		Token: h.App.Gama.AppSecret,
	})
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	if ocr.Code != 0 {
		resp.ERROR(c, ocr.Msg)
		return
	}
	resp.SUCCESS(c, ocr.Data)
}
