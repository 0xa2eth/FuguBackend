package app

import (
	"FuguBackend/app/router"
	"FuguBackend/config"
	"FuguBackend/pkg/env"
	"FuguBackend/pkg/logger"
	"FuguBackend/pkg/timeutil"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type App struct {
	//config *config.Config
	//router *gin.Engine
	Server *http.Server
}

func NewApp() (*App, error) {
	// 初始化 access logger
	accessLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", config.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(config.ProjectAccessLogFile),
	)
	if err != nil {
		panic(err)
	}

	// 初始化 cron logger
	cronLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
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
	server, err := router.NewHTTPServer(accessLogger, cronLogger)
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

	config.Logger.Info("[INFO]", zap.String("service is", " starting ..."), zap.Any("address:", config.Conf.Server.Addr))

	if err := p.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		config.Logger.Fatal("http server startup err", zap.Error(err))
	}
}

func (p *App) AppClose() error {
	//p.appCtx.Rds.Close()
	//db, _ := p.appCtx.Db.DB()
	//db.Close()
	return nil
}
