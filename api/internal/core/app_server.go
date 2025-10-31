package core

import (
	"github.com/gin-gonic/gin"
	"github.com/szwtdl/chinatax-system/internal/auth"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/internal/db"
	"github.com/szwtdl/chinatax-system/internal/resp"
	"github.com/szwtdl/chinatax-system/pkg/utils"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"runtime/debug"
	"strings"
)

type AppServer struct {
	Debug         bool
	Config        *types.AppConfig
	Engine        *gin.Engine
	db            *gorm.DB
	Redis         *db.RedisClient
	PermissionMap map[string]auth.PermissionChecker
}

func (s *AppServer) RegisterModule(name string, checker auth.PermissionChecker) {
	if s.PermissionMap == nil {
		s.PermissionMap = make(map[string]auth.PermissionChecker)
	}
	s.PermissionMap[name] = checker
}

func NewServer(appConfig *types.AppConfig, db *gorm.DB) *AppServer {
	return &AppServer{
		Debug:  true,
		Config: appConfig,
		db:     db,
		Engine: gin.Default(),
	}
}

func (s *AppServer) Init(debug bool) {
	if debug {
		s.Debug = debug
	}
	s.Engine.Use(authorizeMiddleware(s))
	s.Engine.Use(errorHandler)
	s.Engine.Static("/static", s.Config.StaticDir)
}

func (s *AppServer) Run() error {
	//var sysConfig model.Config
	//res := s.db.Where("marker", "system").Select("marker", "config_json").First(&sysConfig)
	//if res.Error != nil {
	//	return res.Error
	//}
	//err := utils.JsonDecode(sysConfig.Config, &s.SysConfig)
	//if err != nil {
	//	return err
	//}
	return s.Engine.Run(s.Config.Listen)
}

func errorHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			c.JSON(http.StatusOK, types.BizVo{Code: types.Failed, Message: "系统错误"})
			c.Abort()
		}
	}()
	c.Next()
}

func authorizeMiddleware(s *AppServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := strings.Trim(c.FullPath(), "/")
		// 白名单直接放行
		if IsWhiteListed(path, s.Config.NotAuth.Paths) {
			c.Next()
			return
		}
		// 获取并验证 token
		user, err := utils.ParseUserFromJWT(utils.BearerToken(c))
		if err != nil {
			resp.ERROR(c, types.BizMsg[types.NotAuthorized], types.AuthError)
			c.Abort()
			return
		}
		// 解析模块
		if path == "" {
			resp.ERROR(c, types.BizMsg[types.NotAuthorized], types.AuthError)
			c.Abort()
			return
		}
		parts := strings.Split(path, "/")
		module := parts[0]
		// 权限校验
		checker, exists := s.PermissionMap[module]
		if !exists || !checker.CheckPermission(user.UserID, path, c.Request.Method) {
			resp.ERROR(c, "没有权限", 403)
			c.Abort()
			return
		}
		c.Next()
	}
}

func IsWhiteListed(path string, patterns []string) bool {
	// 去掉路径开头的 /
	path = strings.TrimLeft(path, "/")
	for _, p := range patterns {
		// 同样去掉 pattern 开头的 /
		p = strings.TrimLeft(p, "/")
		matched, _ := regexp.MatchString(p, path)
		if matched {
			return true
		}
	}
	return false
}
