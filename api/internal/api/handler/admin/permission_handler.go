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

type PermissionHandler struct {
	handler.BaseHandler
	service *service.PermissionService
}

func NewPermissionHandler(app *types.AppConfig, db *gorm.DB, log *zap.SugaredLogger) *PermissionHandler {
	repo := repository.NewBaseRepository[model.Permission]()
	servicePermission := service.NewPermissionService(db, repo)
	return &PermissionHandler{
		BaseHandler: handler.BaseHandler{
			App: app,
			Log: log,
		},
		service: servicePermission,
	}
}

func (h *PermissionHandler) List(c *gin.Context) {
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
	var lists = make([]vo.Permission, 0)
	items, total, err := h.service.Pagination(map[string]interface{}{}, page, limit, "id DESC")
	if err != nil {
		resp.SUCCESS(c, "获取失败")
		return
	}
	for _, item := range items {
		lists = append(lists, vo.Permission{
			ID:        item.ID,
			Name:      item.Name,
			ParentID:  item.ParentID,
			Path:      item.Path,
			Method:    item.Method,
			Desc:      item.Description,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: item.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	pageVo := vo.NewPage(total, page, limit, lists)
	resp.SUCCESS(c, pageVo)
}

func (h *PermissionHandler) Create(c *gin.Context) {
	var req struct {
		ParentID    uint   `json:"parent_id,omitempty"`
		Name        string `json:"name" binding:"required"`
		Path        string `json:"path" binding:"required"`
		Method      string `json:"method" binding:"required"`
		Description string `json:"description,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	existingPerm, err := h.service.FindOne(map[string]interface{}{"name": req.Name})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err.Error())
		resp.ERROR(c, "查询失败")
		return
	}
	if existingPerm != nil {
		resp.ERROR(c, "权限已存在")
		return
	}

	perm := &model.Permission{
		ParentID:    req.ParentID, // 默认为0
		Name:        req.Name,
		Path:        req.Path,
		Method:      req.Method,
		Description: req.Description,
	}
	if err := h.service.CreatePermission(perm); err != nil {
		resp.ERROR(c, types.BizMsg[types.DatabaseError])
		return
	}
	resp.SUCCESS(c, gin.H{
		"id":   perm.ID,
		"name": perm.Name,
	})
}

func (h *PermissionHandler) Update(c *gin.Context) {
	var req struct {
		ID          uint   `json:"id" binding:"required"`
		ParentID    uint   `json:"parent_id,omitempty"`
		Name        string `json:"name" binding:"required"`
		Path        string `json:"path" binding:"required"`
		Method      string `json:"method" binding:"required"`
		Description string `json:"description,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	perm, err := h.service.FindOne(map[string]interface{}{"id": req.ID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ERROR(c, "权限不存在")
			return
		}
		resp.ERROR(c, types.BizMsg[types.DatabaseError])
		return
	}
	existingPerm, err := h.service.FindOne(map[string]interface{}{"name": req.Name})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	if existingPerm != nil && existingPerm.ID != req.ID {
		resp.ERROR(c, "权限名称已存在")
		return
	}
	perm.ParentID = req.ParentID
	perm.Name = req.Name
	perm.Path = req.Path
	perm.Method = req.Method
	perm.Description = req.Description
	err = h.service.UpdatePermission(perm)
	if err != nil {
		resp.ERROR(c, types.BizMsg[types.DatabaseError])
		return
	}
	resp.SUCCESS(c, "更新成功")
}

func (h *PermissionHandler) Delete(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ERROR(c, types.BizMsg[types.InvalidParam])
		return
	}
	perm, err := h.service.FindOne(map[string]interface{}{"id": req.ID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ERROR(c, "权限不存在")
			return
		}
		resp.ERROR(c, types.BizMsg[types.DatabaseError])
		return
	}
	children, err := h.service.Find(map[string]interface{}{"parent_id": perm.ID})
	if err != nil {
		resp.ERROR(c, types.BizMsg[types.DatabaseError])
		return
	}
	if len(children) > 0 {
		resp.ERROR(c, "存在子权限，无法删除")
		return
	}
	if err := h.service.DeleteByID(perm.ID); err != nil {
		resp.ERROR(c, types.BizMsg[types.DatabaseError])
		return
	}
	resp.SUCCESS(c, "删除成功")
}
