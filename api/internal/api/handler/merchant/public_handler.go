package merchant

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/config"
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
	"time"
)

type PublicHandler struct {
	handler.BaseHandler
	merchantService *service.MerchantService
	cityTaxService  *service.CityTaxService
	emailService    *service.EmailVerificationService
	platformService *service.MerchantPlatformService
	minApp          *miniprogram.MiniProgram
}

func NewPublicHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *PublicHandler {
	merchantService := service.NewMerchantService(db, repository.NewBaseRepository[model.Merchant]())
	platformService := service.NewMerchantPlatformService(db, repository.NewBaseRepository[model.MerchantPlatform]())
	cityTaxService := service.NewCityTaxService(db, repository.NewBaseRepository[model.CityTax]())
	emailService := service.NewEmailVerificationService(db)
	wc := wechat.NewWechat()
	minApp := wc.GetMiniProgram(&config.Config{
		AppID:     app.MinApp.AppID,
		AppKey:    app.MinApp.AppKey,
		AppSecret: app.MinApp.AppSecret,
		Cache: cache.NewRedis(context.Background(), &cache.RedisOpts{
			Host:        app.RedisConfig.Host,
			Password:    app.RedisConfig.Password,
			Database:    app.RedisConfig.DBName,
			MaxActive:   10,
			MaxIdle:     10,
			IdleTimeout: 60,
		}),
	})
	return &PublicHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		merchantService: merchantService,
		platformService: platformService,
		cityTaxService:  cityTaxService,
		emailService:    emailService,
		minApp:          minApp,
	}
}

func (h *PublicHandler) CityList(c *gin.Context) {
	list, err := h.cityTaxService.Find(map[string]interface{}{})
	if err != nil {
		resp.ERROR(c, "没有数据")
		return
	}
	StatusMap := map[string]string{
		"0": "正常",
		"1": "维护中",
		"2": "下架",
	}
	items := make([]vo.CityTax, 0)
	for _, item := range list {
		items = append(items, vo.CityTax{
			Name:     item.Name,
			ClientID: item.ClientID,
			Status:   StatusMap[item.Status],
			Domain:   fmt.Sprintf("https://etax.%s.chinatax.gov.cn", item.CityName),
		})
	}
	resp.SUCCESS(c, items)
}

func (h *PublicHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}

	if !utils.IsPhoneNumber(req.Username) {
		resp.ERROR(c, "手机号格式错误")
		return
	}

	user, err := h.merchantService.GetByUsername(req.Username)
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
	err = h.merchantService.Update(user.ID, map[string]interface{}{
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

func (h *PublicHandler) MobileLogin(c *gin.Context) {
	var req struct {
		Code          string `json:"code" binding:"required"`
		EncryptedData string `json:"encrypted_data" binding:"required"`
		Iv            string `json:"iv" binding:"required"`
		Channel       string `json:"channel,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}

	session, err := h.minApp.GetAuth().Code2Session(req.Code)
	if err != nil {
		h.Log.Error(err)
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	decrypt, err := h.minApp.GetEncryptor().Decrypt(session.SessionKey, req.EncryptedData, req.Iv)
	if err != nil {
		h.Log.Error(err)
		resp.ERROR(c, "数据解密失败")
		return
	}
	mobile := decrypt.PhoneNumber
	platformAccount, merchant, err := h.platformService.GetOrCreatePlatform(
		mobile,
		"minapp",
		session.OpenID,
		"微信用户",
		"https://tax.szwtdl.cn/static/images/default.png",
		c.ClientIP(),
	)
	if err != nil {
		h.Log.Error(err)
		resp.ERROR(c, "登录失败")
		return
	}
	// 统一生成 JWT
	token, err := utils.GenerateJWT(merchant.ID, merchant.Username)
	if err != nil {
		resp.ERROR(c, "生成Token失败")
		return
	}

	// 更新 token 信息
	if err = h.merchantService.Update(merchant.ID, map[string]interface{}{
		"token":     token,
		"login_ip":  c.ClientIP(),
		"last_time": time.Now().Format("2006-01-02 15:04:05"),
	}); err != nil {
		resp.ERROR(c, "更新令牌失败")
		return
	}

	resp.SUCCESS(c, map[string]interface{}{
		"token":  token,
		"openid": platformAccount.PlatformUID,
		"mobile": mobile,
	})
}

func (h *PublicHandler) Register(c *gin.Context) {
	var req struct {
		Username   string `json:"username" binding:"required"`
		Password   string `json:"password" binding:"required,min=8,max=20"`
		InviteCode string `json:"invite_code,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}

	if !utils.IsPhoneNumber(req.Username) {
		resp.ERROR(c, "手机号格式错误")
		return
	}

	if !utils.IsStrongPassword(req.Password) {
		resp.ERROR(c, "账号太过简单")
		return
	}

	if exist, _ := h.merchantService.GetByUsername(req.Username); exist != nil {
		resp.ERROR(c, "账号已经存在")
		return
	}
	salt, _ := utils.GenerateSalt(16)
	hashedPwd, _ := utils.HashPassword(req.Password, salt)
	merchant := &model.Merchant{
		Username:     req.Username,
		Password:     hashedPwd,
		Salt:         salt,
		InviteCode:   req.InviteCode,
		IsPassword:   true,
		RegisterTime: time.Now().Format("2006-01-02 15:04:05"),
		RegisterIP:   c.ClientIP(),
	}
	if err := h.merchantService.Create(merchant); err != nil {
		resp.ERROR(c, "注册失败:"+err.Error())
	}
	resp.SUCCESS(c, "注册成功")
}

func (h *PublicHandler) Logout(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "登录令牌错误")
	}
	user, err := h.merchantService.GetToken(req.Token)
	if err != nil {
		resp.ERROR(c, "令牌不存在")
		return
	}
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		resp.ERROR(c, "令牌失败")
		return
	}
	err = h.merchantService.Update(user.ID, map[string]interface{}{
		"token": token,
	})
	if err != nil {
		resp.ERROR(c, "退出失败")
		return
	}
	resp.SUCCESS(c, "退出成功")
}

func (h *PublicHandler) SmsCode(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "账号参数错误")
		return
	}

	if !utils.IsPhoneNumber(req.Username) {
		resp.ERROR(c, "手机号格式错误")
		return
	}

	user, err := h.merchantService.GetByUsername(req.Username)
	if err != nil || user == nil {
		resp.ERROR(c, "账号不存在")
		return
	}
	resp.SUCCESS(c, map[string]interface{}{
		"code":  utils.GenerateRandomCode(6),
		"phone": req.Username,
	})
}

func (h *PublicHandler) SendEmail(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "账号参数错误")
		return
	}

	if !utils.IsPhoneNumber(req.Username) {
		resp.ERROR(c, "手机号格式错误")
		return
	}

	user, err := h.merchantService.GetByUsername(req.Username)
	if err != nil || user == nil {
		resp.ERROR(c, "账号不存在")
		return
	}

	if user.Email == "" {
		resp.ERROR(c, "您没有绑定邮箱")
		return
	}

	// 检查绑定邮箱是否正确
	if user.Email != req.Email {
		resp.ERROR(c, "绑定邮箱不匹配")
		return
	}
	// 生成验证码
	code, err := h.emailService.GenerateCode(req.Username, req.Email, 5, 60)
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	subject := fmt.Sprintf("[%s] 重置密码验证码", h.App.Email.PlatformName)
	body := fmt.Sprintf("您好，您的验证码是：<b>%s</b>，5分钟内有效", code)
	err = utils.SendEmailTLS(
		req.Email,
		subject,
		body,
		h.App.Email.SMTPHost,
		h.App.Email.Username,
		h.App.Email.Password,
		h.App.Email.SMTPPort)
	if err != nil {
		resp.ERROR(c, "发送邮件失败："+err.Error())
		return
	}
	resp.SUCCESS(c, "验证码已发送到绑定邮箱，请注意查收")
}

func (h *PublicHandler) RePassword(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Type     string `json:"type" binding:"required,oneof=email sms"`
		Code     string `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}

	if !utils.IsPhoneNumber(req.Username) {
		resp.ERROR(c, "手机号格式错误")
		return
	}

	user, err := h.merchantService.GetByUsername(req.Username)
	if err != nil || user == nil {
		resp.ERROR(c, "账号不存在")
		return
	}
	switch req.Type {
	case "email":
		if user.Email == "" {
			resp.ERROR(c, "您没有绑定邮箱")
			return
		}
		result, ok := h.emailService.VerifyCode(req.Username, user.Email, req.Code)
		if ok != nil || !result {
			resp.ERROR(c, "验证码无效或已过期")
			return
		}
	case "sms":
		resp.ERROR(c, "短信验证暂未实现")
		return
	default:
		resp.ERROR(c, "验证类型不支持")
		return
	}
	newPassword, err := utils.HashPassword(req.Password, user.Salt)
	if err != nil {
		resp.ERROR(c, "生成密码失败")
		return
	}
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		resp.ERROR(c, "生成 Token 失败")
		return
	}
	err = h.merchantService.Update(user.ID, map[string]interface{}{
		"password": newPassword,
		"token":    token,
	})
	if err != nil {
		resp.ERROR(c, "修改密码失败")
		return
	}
	resp.SUCCESS(c, "密码重置成功")
}
