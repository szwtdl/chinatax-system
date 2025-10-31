package merchant

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/szwtdl/chinatax-system/internal/api/handler"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"github.com/szwtdl/chinatax-system/internal/resp"
	"github.com/szwtdl/chinatax-system/internal/service"
	"github.com/szwtdl/chinatax-system/internal/vo"
	"github.com/szwtdl/chinatax_api/chinatax"
	"github.com/szwtdl/chinatax_api/platforms"
	"github.com/szwtdl/chinatax_api/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ChinaTaxHandler struct {
	handler.BaseHandler
	ApiClient      chinatax.ApiClient
	cityTaxService *service.CityTaxService
}

func NewChinaTaxHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *ChinaTaxHandler {
	repo := repository.NewBaseRepository[model.CityTax]()
	cityTaxService := service.NewCityTaxService(db, repo)
	return &ChinaTaxHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		ApiClient: platforms.NewClient(map[string]string{
			"app_id":    "test",
			"secret":    "test",
			"city_name": "shenzhen",
			"debug":     "true",
		}),
		cityTaxService: cityTaxService,
	}
}

func (h *ChinaTaxHandler) SendSms(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		ClientID string `json:"client_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "参数错误")
		return
	}
	city, err := h.cityTaxService.GetClient(req.ClientID)
	if err != nil {
		resp.ERROR(c, "未开通地区")
		return
	}
	response, err := h.ApiClient.SendSmsCode(request.SendSmsCodeRequest{
		Username: req.Username,
		Password: req.Password,
		CityName: city.CityName,
		ClientID: req.ClientID,
		OcrHost:  h.App.OcrServer.Host,
		OcrUser:  h.App.OcrServer.Username,
		OcrPass:  h.App.OcrServer.Password,
	})
	if err != nil {
		h.Log.Error(err)
		resp.ERROR(c, err.Error())
		return
	}
	if response.Code != 0 {
		resp.ERROR(c, response.Msg)
		return
	}
	resp.SUCCESS(c, response.Data)
}

func (h *ChinaTaxHandler) Login(c *gin.Context) {
	var req struct {
		ClientID     string `json:"client_id" binding:"required"`
		PublicUUID   string `json:"public_uuid" binding:"required"`
		AccountUUID  string `json:"account_uuid" binding:"required"`
		SmsCodeID    string `json:"sms_code_id" binding:"required"`
		SmsCode      string `json:"sms_code" binding:"required,len=6"`
		SignatureKey string `json:"signature_key" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "参数错误")
		return
	}
	city, err := h.cityTaxService.GetClient(req.ClientID)
	if err != nil {
		resp.ERROR(c, "未开通地区")
		return
	}
	response, err := h.ApiClient.SmsLogin(request.SmsLoginRequest{
		CityName:     city.CityName,
		ClientID:     req.ClientID,
		PublicUUID:   req.PublicUUID,
		AccountUUID:  req.AccountUUID,
		SignatureKey: req.SignatureKey,
		SmsCodeID:    req.SmsCodeID,
		SmsCode:      req.SmsCode,
	})
	if err != nil {
		h.Log.Error(err)
		resp.ERROR(c, err.Error())
		return
	}
	if response.Code != 0 {
		h.Log.Error(response.Msg)
		resp.ERROR(c, response.Msg)
		return
	}
	resp.SUCCESS(c, response.Data)
}

func (h *ChinaTaxHandler) Info(c *gin.Context) {
	var req struct {
		ClientID     string `json:"client_id" binding:"required"`
		PublicUUID   string `json:"public_uuid" binding:"required"`
		AccessToken  string `json:"access_token" binding:"required"`
		SignatureKey string `json:"signature_key" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	city, err := h.cityTaxService.GetClient(req.ClientID)
	if err != nil {
		resp.ERROR(c, "未开通地区")
		return
	}
	info, err := h.ApiClient.UserInfo(request.UserInfoRequest{
		ClientID:      city.ClientID,
		PublicUUID:    req.PublicUUID,
		Authorization: req.AccessToken,
		SignatureKey:  req.SignatureKey,
	})
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	if info.Code != 0 {
		resp.ERROR(c, info.Msg)
		return
	}
	resp.SUCCESS(c, info.Data)
}

func (h *ChinaTaxHandler) Bind(c *gin.Context) {
	var req struct {
		UserID    string `json:"user_id" binding:"required"`
		AccountID string `json:"account_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	resp.SUCCESS(c, types.BizMsg[types.Success])
}

func (h *ChinaTaxHandler) MobileList(c *gin.Context) {
	var req struct {
		UserID       string `json:"user_id" binding:"required"`
		ClientID     string `json:"client_id" binding:"required"`
		PublicUUID   string `json:"public_uuid" binding:"required"`
		AccessToken  string `json:"access_token" binding:"required"`
		SignatureKey string `json:"signature_key" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	list, err := h.ApiClient.GetMobileList(request.GetMobileListRequest{
		ClientID:      req.ClientID,
		PublicUUID:    req.PublicUUID,
		Authorization: req.AccessToken,
		SignatureKey:  req.SignatureKey,
		UserID:        req.UserID,
	})
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	resp.SUCCESS(c, list.Data)
}

func (h *ChinaTaxHandler) DefaultMobile(c *gin.Context) {
	var req struct {
		UserID       string `json:"user_id" binding:"required"`
		ClientID     string `json:"client_id" binding:"required"`
		PublicUUID   string `json:"public_uuid" binding:"required"`
		AccessToken  string `json:"access_token" binding:"required"`
		SignatureKey string `json:"signature_key" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	list, err := h.ApiClient.GetMobileList(request.GetMobileListRequest{
		ClientID:      req.ClientID,
		PublicUUID:    req.PublicUUID,
		Authorization: req.AccessToken,
		SignatureKey:  req.SignatureKey,
		UserID:        req.UserID,
	})
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	var mobiles []vo.Mobile
	dataBytes, _ := json.Marshal(list.Data)
	if err := json.Unmarshal(dataBytes, &mobiles); err != nil {
		resp.ERROR(c, "数据格式错误")
		return
	}
	if len(mobiles) == 0 {
		resp.ERROR(c, "没有手机号码")
		return
	}
	for _, item := range mobiles {
		if item.DefaultFlag == "1" {
			resp.SUCCESS(c, item)
			return
		}
	}
	resp.ERROR(c, types.BizMsg[types.InvalidParam])
}

func (h *ChinaTaxHandler) GetSmsCode(c *gin.Context) {
	var req struct {
		ClientID     string `json:"client_id" binding:"required"`
		UserID       string `json:"user_id" binding:"required"`
		SignatureKey string `json:"signature_key" binding:"required"`
		MobileID     string `json:"mobile_id" binding:"required"`
		AccessToken  string `json:"access_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	city, err := h.cityTaxService.GetClient(req.ClientID)
	if err != nil {
		resp.ERROR(c, "地区未开发，或维护中")
		return
	}
	response, err := h.ApiClient.SendMobileSendCode(request.SendMobileSendCodeRequest{
		ClientID:      city.ClientID,
		CityName:      city.CityName,
		SignatureKey:  req.SignatureKey,
		Authorization: req.AccessToken,
		UserID:        req.UserID,
		MobileID:      req.MobileID,
		OcrHost:       h.App.OcrServer.Host,
		OcrUser:       h.App.OcrServer.Username,
		OcrPass:       h.App.OcrServer.Password,
	})
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	h.Log.Info("调试信息", response)
	resp.SUCCESS(c, response, types.Success)
}

func (h *ChinaTaxHandler) VerifyOldSmsCode(c *gin.Context) {
	var req struct {
		UserID     string `json:"user_id" binding:"required"`
		MobileID   string `json:"mobile_id" binding:"required"`
		SendCodeID string `json:"send_code_id" binding:"required"`
		SmsCode    string `json:"sms_code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	response, err := h.ApiClient.VerifyOldSmsCode(request.VerifyOldSmsCodeRequest{
		UserID:    req.UserID,
		MobileID:  req.MobileID,
		SmsCodeID: req.SendCodeID,
		SmsCode:   req.SmsCode,
	})
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	resp.SUCCESS(c, response.Data)
}

func (h *ChinaTaxHandler) SendSmsCodeMobileChange(c *gin.Context) {
	var req struct {
		ClientID     string `json:"client_id" binding:"required"`
		PublicUUID   string `json:"public_uuid" binding:"required"`
		SignatureKey string `json:"signature_key" binding:"required"`
		AccessToken  string `json:"access_token" binding:"required"`
		Mobile       string `json:"mobile" binding:"required"`
		Name         string `json:"name" binding:"required"`
		IDType       string `json:"id_type" binding:"required"`
		IDCard       string `json:"id_card" binding:"required"`
		UserID       string `json:"user_id" binding:"required"`
		Scene        string `json:"scene" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	response, err := h.ApiClient.SendSmsCodeMobileChange(request.SendSmsCodeMobileChangeRequest{
		ClientID:      req.ClientID,
		PublicUUID:    req.PublicUUID,
		SignatureKey:  req.SignatureKey,
		Authorization: req.AccessToken,
		Mobile:        req.Mobile,
		Name:          req.Name,
		IDType:        req.IDType,
		IDCard:        req.IDCard,
		UserID:        req.UserID,
		Scene:         req.Scene,
	})
	if err != nil {
		resp.ERROR(c, err.Error())
		return
	}
	resp.SUCCESS(c, response.Data)
}

func (h *ChinaTaxHandler) ModifyDefaultMobileByMessage(c *gin.Context) {
	var req struct {
		MobileID         string `json:"mobile_id" binding:"required"`
		NewMobile        string `json:"new_mobile" binding:"required"`
		SmsCodeID        string `json:"sms_code_id" binding:"required"`
		SmsCode          string `json:"sms_code" binding:"required"`
		UserID           string `json:"user_id" binding:"required"`
		DefaultSmsCodeID string `json:"default_sms_code_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	response, ok := h.ApiClient.ModifyDefaultMobileByMessage(request.ModifyDefaultMobileByMessageRequest{
		MobileID:         req.MobileID,
		NewMobile:        req.NewMobile,
		SmsCodeID:        req.SmsCodeID,
		SmsCode:          req.SmsCode,
		UserID:           req.UserID,
		DefaultSmsCodeID: req.DefaultSmsCodeID,
	})
	if ok != nil {
		resp.ERROR(c, ok.Error())
		return
	}
	resp.SUCCESS(c, response.Data)
}
