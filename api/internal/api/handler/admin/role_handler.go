package admin

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/szwtdl/chinatax-system/internal/api/handler"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"github.com/szwtdl/chinatax-system/internal/resp"
	"github.com/szwtdl/chinatax-system/internal/service"
	"github.com/szwtdl/chinatax-system/internal/vo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RoleHandler struct {
	handler.BaseHandler
	service *service.RoleService
}

func NewRoleHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *RoleHandler {
	repo := repository.NewBaseRepository[model.Role]()
	roleService := service.NewRoleService(db, repo)
	return &RoleHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		service: roleService,
	}
}

func (h *RoleHandler) List(c *gin.Context) {
	page := h.GetInt(c, "page", 1)
	limit := h.GetInt(c, "limit", 20)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	var lists = make([]vo.Role, 0)
	items, total, err := h.service.Pagination(map[string]interface{}{}, page, limit, "id DESC")
	if err != nil {
		resp.SUCCESS(c, "获取失败")
		return
	}
	for _, item := range items {
		lists = append(lists, vo.Role{
			ID:          item.ID,
			Name:        item.Name,
			DisplayName: item.DisplayName,
			Description: item.Description,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   item.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	pageVo := vo.NewPage(total, page, limit, lists)
	resp.SUCCESS(c, pageVo)
}

func (h *RoleHandler) Create(c *gin.Context) {
	var req struct {
		Name          string `json:"name" binding:"required"`
		DisplayName   string `json:"display_name" binding:"required"`
		Description   string `json:"description,omitempty"`
		Status        int    `json:"status" binding:"required"`
		PermissionIDs []uint `json:"permission_ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}

	if len(req.PermissionIDs) == 0 {
		resp.ERROR(c, "权限不能为空")
		return
	}

	// 检查角色名是否已存在
	existing, err := h.service.FindOne(map[string]interface{}{"name": req.Name})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		resp.ERROR(c, "查询失败："+err.Error())
		return
	}
	if existing != nil {
		resp.ERROR(c, "角色名称已存在")
		return
	}

	role := &model.Role{
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Description: req.Description,
		Status:      req.Status,
	}

	// 事务创建角色和权限关系
	if err := h.service.CreateWithPermissions(role, req.PermissionIDs); err != nil {
		resp.ERROR(c, "创建失败："+err.Error())
		return
	}
	resp.SUCCESS(c, "创建成功")
}

func (h *RoleHandler) Update(c *gin.Context) {
	var req struct {
		ID            uint   `json:"id" binding:"required"`
		Name          string `json:"name" binding:"required"`
		DisplayName   string `json:"display_name" binding:"required"`
		Description   string `json:"description,omitempty"`
		Status        int    `json:"status" binding:"required"`
		PermissionIDs []uint `json:"permission_ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}

	if len(req.PermissionIDs) == 0 {
		resp.ERROR(c, "权限不能为空")
		return
	}

	// 查询角色是否存在
	role, err := h.service.GetByID(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ERROR(c, "角色不存在")
			return
		}
		resp.ERROR(c, "查询失败："+err.Error())
		return
	}

	// 检查角色名是否重复（排除自己）
	existing, err := h.service.FindOne(map[string]interface{}{"name": req.Name})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		resp.ERROR(c, "查询失败："+err.Error())
		return
	}
	if existing != nil && existing.ID != req.ID {
		resp.ERROR(c, "角色名称已存在")
		return
	}
	// 更新角色信息 + 权限，使用事务
	if err = h.service.UpdateWithPermissions(role, req.PermissionIDs); err != nil {
		resp.ERROR(c, "更新失败："+err.Error())
		return
	}
	resp.SUCCESS(c, "更新成功")
}

func (h *RoleHandler) Delete(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	// 查询角色是否存在
	role, err := h.service.GetByID(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ERROR(c, "角色不存在")
			return
		}
		resp.ERROR(c, "查询失败："+err.Error())
		return
	}
	if err = h.service.DeleteWithRelations(req.ID); err != nil {
		resp.ERROR(c, "删除失败："+err.Error())
		return
	}
	resp.SUCCESS(c, fmt.Sprintf("角色 %s 删除成功", role.Name))
}
