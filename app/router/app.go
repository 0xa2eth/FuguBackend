package router

import (
	"fmt"
	"net/http"

	"FuguBackend/app/alert"
	"FuguBackend/app/metrics"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/twittersvc"
	"FuguBackend/app/repository/cron"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
	"FuguBackend/app/router/interceptor"
	"FuguBackend/config"
	"FuguBackend/pkg/env"
	"FuguBackend/pkg/logger"
	"FuguBackend/pkg/timeutil"

	"go.uber.org/zap"
)

type App struct {
	//config *config.Config
	//router *gin.Engine
	Server *http.Server
}

func NewApp() (*App, error) {
	config.LoadConfig()

	server, err := NewHTTPServer()
	if err != nil {
		return nil, err
	}

	httpServer := &http.Server{
		Addr:    config.ProjectPort,
		Handler: server.Mux,
	}

	return &App{
		Server: httpServer,
	}, nil
}

func (p *App) AppStart() {

	config.Lg.Info("[INFO]", zap.String("service is", " starting ..."), zap.Any("address:", config.Conf.Server.Addr))

	if err := p.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		config.Lg.Fatal("http server startup err", zap.Error(err))
	}
}

func (p *App) AppClose() error {
	//p.appCtx.Rds.Close()
	//db, _ := p.appCtx.Db.DB()
	//db.Close()
	return nil
}

type Resource struct {
	Mux           core.Mux
	Logger        *zap.Logger
	Db            mysql.Repo
	Cache         redis.Repo
	Interceptors  interceptor.Interceptor
	CronServer    cron.Server
	TwitterServer twittersvc.TwitterServiceMaster
}
type Server struct {
	Mux        core.Mux
	Db         mysql.Repo
	Cache      redis.Repo
	CronServer cron.Server
}

func NewHTTPServer() (*Server, error) {
	//if logger == nil {
	//	return nil, errors.New("logger required")
	//}

	r := new(Resource)
	//r.Logger = logger
	r.Logger = config.Lg
	r.TwitterServer = twittersvc.NewTwitterServiceMaster()

	//_, ok := file.IsExists(config.ProjectInstallMark)
	//if !ok { // 未安装
	//	openBrowserUri += "/install"
	//} else { // 已安装

	// 初始化 DB
	dbRepo, err := mysql.New()
	if err != nil {
		config.Lg.Fatal("new db err", zap.Error(err))
	}
	r.Db = dbRepo

	// 初始化 Cache
	cacheRepo, err := redis.New()
	if err != nil {
		config.Lg.Fatal("new cache err", zap.Error(err))
	}
	r.Cache = cacheRepo

	//// 初始化 CRON Server
	//cronServer, err := cron.New(cronLogger, dbRepo, cacheRepo)
	//if err != nil {
	//	logger.Fatal("new cron err", zap.Error(err))
	//}
	//cronServer.Start()
	//r.cronServer = cronServer
	//}

	//openBrowserUri := config.ProjectDomain + config.ProjectPort
	mux, err := core.New(config.Lg,
		//core.WithEnableOpenBrowser(openBrowserUri),
		core.WithEnableCors(),
		core.WithEnableRate(),
		core.WithAlertNotify(alert.NotifyHandler(config.Lg)),
		core.WithRecordMetrics(metrics.RecordHandler(config.Lg)),
	)

	if err != nil {
		panic(err)
	}

	r.Mux = mux
	r.Interceptors = interceptor.New(config.Lg, r.Cache, r.Db, r.TwitterServer)

	// 设置 API 路由
	SetApiRouter(r)

	s := new(Server)
	s.Mux = mux
	s.Db = r.Db
	s.Cache = r.Cache
	s.CronServer = r.CronServer

	return s, nil
}
func GetTraceLogger() (accessLogger, cronLogger *zap.Logger) {
	// 初始化 access logger
	accessLogger, err := logger.NewJSONLogger(
		//logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", config.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(config.ProjectAccessLogFile),
	)
	if err != nil {
		panic(err)
	}

	// 初始化 cron logger
	cronLogger, err = logger.NewJSONLogger(
		//logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", config.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(config.ProjectCronLogFile),
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = accessLogger.Sync()
		_ = cronLogger.Sync()
	}()
	return accessLogger, cronLogger
}
