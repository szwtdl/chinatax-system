package admin

import (
	"github.com/gin-gonic/gin"
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
)

type AccountHandler struct {
	handler.BaseHandler
	adminService *service.AdminService
}

func NewAccountHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *AccountHandler {
	Repo := repository.NewBaseRepository[model.Admin]()
	adminService := service.NewAdminService(db, Repo)
	return &AccountHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		adminService: adminService,
	}
}

func (h *AccountHandler) List(c *gin.Context) {
	page := h.GetInt(c, "page", 1)
	limit := h.GetInt(c, "limit", 10)
	keyword := c.Query("keyword")
	filter := c.Query("filter")
	where := buildAdminWhere(keyword, filter)
	var lists = make([]vo.Admin, 0)
	account, total, err := h.adminService.Pagination(where, page, limit, "id DESC")
	if err != nil {
		resp.ERROR(c, "查询失败")
		return
	}
	for _, v := range account {
		lists = append(lists, vo.Admin{
			ID:           v.ID,
			Avatar:       v.Avatar,
			Nickname:     v.Nickname,
			Username:     v.Username,
			Status:       v.Status,
			Super:        v.Super,
			RegisterTime: v.RegisterTime,
			RegisterIp:   v.RegisterIP,
			LoginIP:      v.LoginIp,
			LastTime:     v.LastTime,
			CreatedAt:    utils.Stamp2str(v.CreatedAt.Unix()),
			UpdatedAt:    utils.Stamp2str(v.UpdatedAt.Unix()),
		})
	}
	pageVo := vo.NewPage(total, page, limit, lists)
	resp.SUCCESS(c, pageVo)
}

func (h *AccountHandler) Create(c *gin.Context) {
	var req struct {
		Avatar   string `json:"avatar" binding:"required"`
		Nickname string `json:"nickname" binding:"required"`
		Username string `json:"username" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6,max=20"`
		RoleIds  []uint `json:"role_ids,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "参数错误")
		return
	}

	if !utils.IsStrongPassword(req.Password) {
		resp.ERROR(c, "账号太过简单")
		return
	}

	if exist, _ := h.adminService.GetByUsername(req.Username); exist != nil {
		resp.ERROR(c, "用户已经存在")
		return
	}

	token := utils.RandString(32)
	salt, _ := utils.GenerateSalt(16)
	password, err := utils.HashPassword(req.Password, salt)
	if err != nil {
		resp.ERROR(c, "密码加密失败")
		return
	}

	admin := &model.Admin{
		Avatar:   req.Avatar,
		Nickname: req.Nickname,
		Username: req.Username,
		Password: password,
		Token:    token,
		Salt:     salt,
	}
	if err = h.adminService.CreateWithRoles(admin, req.RoleIds); err != nil {
		resp.ERROR(c, "创建失败："+err.Error())
		return
	}

	resp.SUCCESS(c, "增加成功")
}

func (h *AccountHandler) Update(c *gin.Context) {
	var req struct {
		ID       uint   `json:"id" binding:"required"`
		Avatar   string `json:"avatar" binding:"required"`
		Nickname string `json:"nickname" binding:"required"`
		Password string `json:"password" binding:"omitempty,min=6,max=20"`
		RoleIds  []uint `json:"role_ids,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "参数错误")
		return
	}

	admin, err := h.adminService.GetByID(req.ID)
	if err != nil {
		resp.ERROR(c, "账号不存在")
		return
	}

	// 判断是否为超级管理员
	if admin.Super {
		resp.ERROR(c, "超级管理员不能修改")
		return
	}

	admin.Avatar = req.Avatar
	admin.Nickname = req.Nickname

	// 如果有密码，重新加密
	if req.Password != "" {
		if !utils.IsStrongPassword(req.Password) {
			resp.ERROR(c, "账号太过简单")
			return
		}
		password, err := utils.HashPassword(req.Password, admin.Salt)
		if err != nil {
			resp.ERROR(c, err.Error())
			return
		}
		admin.Password = password
	}

	if err = h.adminService.UpdateWithRoles(admin, req.RoleIds); err != nil {
		resp.ERROR(c, "修改失败："+err.Error())
		return
	}

	resp.SUCCESS(c, "修改成功")
}

func (h *AccountHandler) Delete(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, "参数错误")
		return
	}

	admin, err := h.adminService.GetByID(req.ID)
	if err != nil {
		resp.ERROR(c, "账号不存在")
		return
	}

	// 超级管理员保护
	if admin.Super {
		resp.ERROR(c, "超级管理员不能删除")
		return
	}
	if err = h.adminService.DeleteWithRoles(req.ID); err != nil {
		resp.ERROR(c, "删除失败："+err.Error())
		return
	}
	resp.SUCCESS(c, "删除成功")
}

func buildAdminWhere(keyword, filter string) map[string]interface{} {
	where := map[string]interface{}{}
	if keyword == "" {
		return where
	}
	allowedFilters := map[string]bool{"username": true, "nickname": true}

	if filter != "" && allowedFilters[filter] {
		where[filter+" LIKE ?"] = "%" + keyword + "%"
	} else {
		where["username LIKE ? OR nickname LIKE ?"] = []interface{}{"%" + keyword + "%", "%" + keyword + "%"}
	}
	return where
}
