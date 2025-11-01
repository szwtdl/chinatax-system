package main

import (
	"context"
	"fmt"
	"github.com/szwtdl/chinatax-system/internal/api/handler/admin"
	"github.com/szwtdl/chinatax-system/internal/api/handler/merchant"
	"github.com/szwtdl/chinatax-system/internal/api/handler/partner"
	"github.com/szwtdl/chinatax-system/internal/auth"
	"github.com/szwtdl/chinatax-system/internal/core"
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"github.com/szwtdl/chinatax-system/internal/db"
	"github.com/szwtdl/chinatax-system/internal/model"
	"github.com/szwtdl/chinatax-system/internal/repository"
	"github.com/szwtdl/chinatax-system/internal/service"
	myLogger "github.com/szwtdl/chinatax-system/pkg/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type AppLifecycle struct {
}

func (l *AppLifecycle) OnStart(context.Context) error {
	fmt.Println("start")
	return nil
}

func (l *AppLifecycle) OnStop(context.Context) error {
	fmt.Println("stop")
	return nil
}

func main() {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "config.toml"
	}
	var debug bool
	debugEnv := os.Getenv("DEBUG")
	if debugEnv == "" {
		debug = true
	} else {
		debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	}

	app := fx.New(
		// 初始化配置应用配置
		fx.Provide(func() *types.AppConfig {
			config, err := core.LoadConfig(configFile)
			if err != nil {
				log.Fatal(err)
			}
			config.Path = configFile
			if debug {
				_ = core.SaveConfig(config)
			}
			return config
		}),
		// 注册Logs
		fx.Provide(func(s *core.AppServer) *zap.SugaredLogger {
			return myLogger.GetLogger(s.Config.Logs)
		}),
		fx.Provide(db.NewGormConfig),
		fx.Provide(db.NewMysql),
		fx.Provide(func(s *core.AppServer) *db.RedisClient {
			var Options = s.Config.RedisConfig
			return db.NewRedis(Options.Host, Options.Password, Options.DBName)
		}),
		// 创建应用服务
		fx.Provide(core.NewServer),
		fx.Provide(
			func(db *gorm.DB) repository.BaseRepository[model.Admin] {
				return repository.NewBaseRepository[model.Admin]()
			},
			func(db *gorm.DB) repository.BaseRepository[model.Merchant] {
				return repository.NewBaseRepository[model.Merchant]()
			},
			func(db *gorm.DB) repository.BaseRepository[model.Partner] {
				return repository.NewBaseRepository[model.Partner]()
			},
		),
		fx.Provide(
			func(db *gorm.DB, repo repository.BaseRepository[model.Admin]) *service.AdminService {
				return service.NewAdminService(db, repo)
			},
			func(db *gorm.DB, repo repository.BaseRepository[model.Merchant]) *service.MerchantService {
				return service.NewMerchantService(db, repo)
			},
			func(db *gorm.DB, repo repository.BaseRepository[model.Partner]) *service.PartnerService {
				return service.NewPartnerService(db, repo)
			},
		),
		fx.Provide(
			func(s *service.AdminService) *auth.AdminPermission { return auth.NewAdminPermission(s) },
			func(s *service.MerchantService) *auth.MerchantPermission {
				return auth.NewMerchantPermission(s)
			},
			func(s *service.PartnerService) *auth.PartnerPermission {
				return auth.NewPartnerPermission(s)
			},
		),
		// 初始化应用服务
		fx.Invoke(func(s *core.AppServer, r *db.RedisClient, admin *auth.AdminPermission, partner *auth.PartnerPermission, merchant *auth.MerchantPermission) {
			s.RegisterModule("admin", admin)
			s.RegisterModule("merchant", merchant)
			s.RegisterModule("partner", partner)
			s.Init(debug)
			s.Redis = r
		}),

		// 用户端(企业)
		fx.Provide(merchant.NewPublicHandler),
		fx.Provide(merchant.NewMerchantHandler),
		fx.Provide(merchant.NewChinaTaxHandler),
		fx.Provide(merchant.NewCompanyHandler),

		// 财务端(财务公司)
		fx.Provide(partner.NewPublicHandler),
		fx.Provide(partner.NewPartnerHandler),

		// 平台端(平台运营)
		fx.Provide(admin.NewPublicHandler),
		fx.Provide(admin.NewCityTaxHandler),
		fx.Provide(admin.NewPartnerHandler),
		fx.Provide(admin.NewMerchantHandler),
		fx.Provide(admin.NewAccountHandler),
		fx.Provide(admin.NewSystemHandler),
		fx.Provide(admin.NewDashboardHandler),
		fx.Provide(admin.NewPermissionHandler),
		fx.Provide(admin.NewRoleHandler),

		// 用户端(企业)
		fx.Invoke(func(s *core.AppServer, h *merchant.PublicHandler) {
			group := s.Engine.Group("/merchant/api/v1/public")
			group.GET("tax_list", h.CityList)
			group.POST("register", h.Register)
			group.POST("login", h.Login)
			group.POST("mobile", h.MobileLogin)
			group.POST("logout", h.Logout)
			group.POST("sendSms", h.SmsCode)
			group.POST("sendEmail", h.SendEmail)
			group.POST("repass", h.RePassword)
		}),
		fx.Invoke(func(s *core.AppServer, h *merchant.ChinaTaxHandler) {
			group := s.Engine.Group("/merchant/api/v1/tax")
			group.POST("sendSms", h.SendSms)
			group.POST("login", h.Login)
			group.POST("info", h.Info)
			group.POST("account/MobileList", h.MobileList)
			group.POST("account/DefaultMobile", h.DefaultMobile)
			group.POST("account/GetSmsCode", h.GetSmsCode)
			group.POST("account/VerifyOldSmsCode", h.VerifyOldSmsCode)
			group.POST("account/SendSmsCodeMobileChange", h.SendSmsCodeMobileChange)
			group.POST("account/ModifyDefaultMobileByMessage", h.ModifyDefaultMobileByMessage)
		}),
		fx.Invoke(func(s *core.AppServer, h *merchant.MerchantHandler) {
			group := s.Engine.Group("/merchant/api/v1/merchant")
			group.POST("info", h.Info)
			group.POST("update", h.Update)
			group.POST("pass", h.RePass)
		}),
		fx.Invoke(func(s *core.AppServer, h *merchant.CompanyHandler) {
			group := s.Engine.Group("/merchant/api/v1/company")
			group.POST("search", h.Search)
			group.POST("detail", h.Detail)
			group.POST("ocr", h.Ocr)
		}),

		// 财务端(财务公司)
		fx.Invoke(func(s *core.AppServer, h *partner.PublicHandler) {
			group := s.Engine.Group("/partner/api/v1/public")
			group.POST("login", h.Login)
			group.POST("register", h.Register)
		}),
		fx.Invoke(func(s *core.AppServer, h *partner.PartnerHandler) {
			group := s.Engine.Group("/partner/api/v1/partner")
			group.POST("info", h.Info)
			group.POST("update", h.Update)
			group.POST("pass", h.RePass)
		}),

		// 平台端(平台运营)
		fx.Invoke(func(s *core.AppServer, h *admin.PublicHandler) {
			group := s.Engine.Group("/admin/api/v1/public")
			group.GET("/captcha", h.GetCaptcha)
			group.POST("login", h.Login)
			group.POST("logout", h.Logout)
			group.GET("info", h.Info)
		}),

		fx.Invoke(func(s *core.AppServer, h *admin.DashboardHandler) {
			group := s.Engine.Group("/admin/api/v1/dashboard")
			group.GET("welcome", h.Welcome)
		}),

		fx.Invoke(func(s *core.AppServer, h *admin.SystemHandler) {
			group := s.Engine.Group("/admin/api/v1/system")
			group.GET("get", h.GetValue)
			group.POST("update", h.SetValue)
		}),

		fx.Invoke(func(s *core.AppServer, h *admin.CityTaxHandler) {
			group := s.Engine.Group("/admin/api/v1/city_tax")
			group.GET("list", h.List)
			group.POST("create", h.Create)
			group.POST("update", h.Update)
			group.POST("delete", h.Delete)
			group.POST("sort", h.Sort)
		}),

		fx.Invoke(func(s *core.AppServer, h *admin.AccountHandler) {
			group := s.Engine.Group("/admin/api/v1/account")
			group.GET("list", h.List)
			group.POST("create", h.Create)
			group.POST("update", h.Update)
			group.POST("delete", h.Delete)
		}),

		fx.Invoke(func(s *core.AppServer, h *admin.PermissionHandler) {
			group := s.Engine.Group("/admin/api/v1/permission")
			group.GET("list", h.List)
			group.POST("create", h.Create)
			group.POST("update", h.Update)
			group.POST("delete", h.Delete)
		}),
		fx.Invoke(func(s *core.AppServer, h *admin.RoleHandler) {
			group := s.Engine.Group("/admin/api/v1/role")
			group.GET("list", h.List)
			group.POST("create", h.Create)
			group.POST("update", h.Update)
			group.POST("delete", h.Delete)
		}),

		fx.Invoke(func(s *core.AppServer, h *admin.PartnerHandler) {
			group := s.Engine.Group("/admin/api/v1/partner")
			group.GET("list", h.List)
			group.POST("create", h.Create)
			group.POST("update", h.Update)
			group.POST("delete", h.Delete)
		}),

		// 启动应用程序
		fx.Invoke(func(s *core.AppServer) {
			ok := s.Run()
			if ok != nil {
				fmt.Println("Application startup", ok)
			}
		}),
		//注册生命周期回调函数
		fx.Invoke(func(lifecycle fx.Lifecycle, lc *AppLifecycle) {
			lifecycle.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return lc.OnStart(ctx)
				},
				OnStop: func(ctx context.Context) error {
					return lc.OnStop(ctx)
				},
			})
		}),
	)

	// 启动应用程序
	go func() {
		if err := app.Start(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	// 监听退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 关闭应用程序
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Stop(ctx); err != nil {
		log.Fatal(err)
	}
}
