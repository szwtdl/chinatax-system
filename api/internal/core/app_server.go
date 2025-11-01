package core

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
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
	s.Engine.Use(corsMiddleware())
	s.Engine.Use(sessionMiddleware(s.Config))
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

func corsMiddleware() gin.HandlerFunc {
	allowedOrigins := []string{
		"https://tax.szwtdl.cn",
		"http://localhost:9527",
	}
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" && utils.Contains(allowedOrigins, origin) {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, Token, Session, Content-Type, X-Requested-With")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "172800")
		}
		// OPTIONS 预检请求直接返回
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("CORS middleware panic:", err)
			}
		}()
		c.Next()
	}
}

func sessionMiddleware(config *types.AppConfig) gin.HandlerFunc {
	var store sessions.Store
	switch config.Session.Driver {
	case types.SessionDriverMem:
		store = memstore.NewStore([]byte(config.Session.SecretKey))
		break
	case types.SessionDriverCookie:
		store = cookie.NewStore([]byte(config.Session.SecretKey))
		break
	default:
		config.Session.Driver = types.SessionDriverCookie
		store = cookie.NewStore([]byte(config.Session.SecretKey))
	}

	store.Options(sessions.Options{
		Path:     config.Session.Path,
		Domain:   config.Session.Domain,
		MaxAge:   config.Session.MaxAge,
		Secure:   config.Session.Secure,
		HttpOnly: config.Session.HttpOnly,
		SameSite: config.Session.SameSite,
	})
	return sessions.Sessions(config.Session.Name, store)
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
			resp.ERROR(c, types.BizMsg[types.NotFound], types.NotFound)
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
