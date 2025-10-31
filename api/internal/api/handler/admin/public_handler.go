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
	"time"
)

type PublicHandler struct {
	handler.BaseHandler
	service *service.AdminService
}

func NewPublicHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *PublicHandler {
	repo := repository.NewBaseRepository[model.Admin]()
	adminService := service.NewAdminService(db, repo)
	return &PublicHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		service: adminService,
	}
}

func (h *PublicHandler) Info(c *gin.Context) {
	token := utils.BearerToken(c)
	user, err := utils.ParseUserFromJWT(token)
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	_, err = h.service.GetByID(user.UserID)
	if err != nil {
		resp.ERROR(c, "用户不存在")
		return
	}
	resp.SUCCESS(c, map[string]interface{}{
		"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		"introduction": "超级管理员",
		"name":         "超级管理员",
		"roles":        []string{"admin"},
	})
}

func (h *PublicHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required,min=4,max=20"`
		Password string `json:"password" binding:"required,min=6,max=20"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	user, err := h.service.GetByUsername(req.Username)
	if err != nil {
		resp.ERROR(c, "用户名不存在")
		return
	}
	if !utils.ComparePassword(user.Password, req.Password, user.Salt) {
		resp.ERROR(c, "账号或密码错误")
		return
	}
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		resp.ERROR(c, "生成Token失败")
		return
	}
	_, err = h.service.UpdateFields(map[string]interface{}{"id": user.ID}, map[string]interface{}{
		"token":     token,
		"login_ip":  c.ClientIP(),
		"last_time": time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		resp.ERROR(c, "更新失败")
		return
	}
	resp.SUCCESS(c, map[string]interface{}{
		"token": token,
	})
}
