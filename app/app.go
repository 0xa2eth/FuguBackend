package app

import (
	"FuguBackend/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	config *config.Config
	router *gin.Engine
	appCtx *ServiceCtx
}

func NewApp(config *config.Config, router *gin.Engine, server *ServiceCtx) (*App, error) {
	return &App{
		config: config,
		router: router,
		appCtx: server,
	}, nil
}

func (p *App) AppStart() error {
	config.Logger.Info("[INFO]", zap.String("service is", " starting ..."), zap.Any("address:", p.config.Server.Addr))
	if err := p.router.Run(p.config.Server.Addr); err != nil {
		config.Logger.Error("start service faild", zap.Error(err))
		return err
	}
	return nil
}

func (p *App) AppClose() error {
	p.appCtx.Rds.Close()
	db, _ := p.appCtx.Db.DB()
	db.Close()
	return nil
}
