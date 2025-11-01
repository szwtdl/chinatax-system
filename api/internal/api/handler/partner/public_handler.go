package partner

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
	PartnerService *service.PartnerService
}

func NewPublicHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *PublicHandler {
	repo := repository.NewBaseRepository[model.Partner]() // ✅ 创建仓储实例
	partnerService := service.NewPartnerService(db, repo)
	return &PublicHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		PartnerService: partnerService,
	}
}

func (h *PublicHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "参数错误")
		return
	}

	if !utils.IsPhoneNumber(req.Username) {
		resp.ERROR(c, "手机号码不合格")
		return
	}

	user, err := h.PartnerService.GetByUsername(req.Username)
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
	err = h.PartnerService.Update(user.ID, map[string]interface{}{
		"token":     token,
		"login_ip":  c.ClientIP(),
		"last_time": time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		resp.ERROR(c, "更新令牌失败")
		return
	}
	resp.SUCCESS(c, map[string]interface{}{
		"token": token,
	})
}

func (h *PublicHandler) Logout(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "令牌不能为空")
		return
	}
	user, err := h.PartnerService.GetToken(req.Token)
	if err != nil {
		resp.ERROR(c, "令牌不存在")
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		resp.ERROR(c, "生成Token失败")
		return
	}
	err = h.PartnerService.Update(user.ID, map[string]interface{}{
		"token": token,
	})
	if err != nil {
		resp.ERROR(c, "退出失败")
		return
	}
	resp.SUCCESS(c, "退出成功")
}

func (h *PublicHandler) Register(c *gin.Context) {
	var req struct {
		Nickname   string `json:"nickname" binding:"required"`
		Username   string `json:"username" binding:"required"`
		Password   string `json:"password" binding:"required,min=6,max=20"`
		InviteCode string `json:"invite_code,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "参数错误")
		return
	}
	if !utils.IsPhoneNumber(req.Username) {
		resp.ERROR(c, "手机号码不合格")
		return
	}
	if exist, _ := h.PartnerService.GetByUsername(req.Username); exist != nil {
		resp.ERROR(c, "账号已经存在")
		return
	}
	salt, _ := utils.GenerateSalt(16)
	hashedPwd, _ := utils.HashPassword(req.Password, salt)
	merchant := &model.Partner{
		Username:     req.Username,
		Password:     hashedPwd,
		Salt:         salt,
		InviteCode:   req.InviteCode,
		RegisterTime: time.Now().Format("2006-01-02 15:04:05"),
		RegisterIP:   c.ClientIP(),
	}
	if err := h.PartnerService.Create(merchant); err != nil {
		resp.ERROR(c, "注册失败:"+err.Error())
		return
	}
	resp.SUCCESS(c, "注册成功")
}
