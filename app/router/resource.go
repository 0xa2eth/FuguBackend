package router

import (
	"FuguBackend/app/alert"
	"FuguBackend/app/metrics"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/cron"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
	"FuguBackend/app/router/interceptor"
	"FuguBackend/config"
	"FuguBackend/pkg/errors"

	"go.uber.org/zap"
)

//type ServiceCtx struct {
//	C   *config.Config
//	Db  *gorm.DB
//	Rds *redis.Repo
//	Aws *s3.S3
//}
//
//var serviceCtx *ServiceCtx
//var once sync.Once
//
//func NewServiceContext() (*ServiceCtx, error) {
//	//初始化日志
//	config.LoadConfig()
//	//初始化mysql
//	initMsql, _ := mysql.New()
//	db, err := initMsql.NewMysqlDB()
//	if err != nil {
//		config.Logger.Error("INIT MYSQL CONFIG FAILED", zap.Error(err))
//		return nil, err
//	}
//	//初始化redis
//	rds := redis.InitRds(config.Conf.RedisC.RedisHost, config.Conf.RedisC.Password, 1)
//
//	//初始化aws
//	awsClient, err := awsc.AWSInit()
//	if err != nil {
//		config.Logger.Error("INIT aws CONFIG FAILED", zap.Error(err))
//		return nil, err
//	}
//	//初始化etcd
//	etcdClient, err := etcd.ETCDInit()
//
//	serviceCtxs := NewServerCtx(WithDB(db), WithRds(rds), WithAws(awsClient), WithEtcd(etcdClient))
//	serviceCtxs.C = config.Conf
//	return serviceCtxs, nil
//}
//
//func GetServiceCtx() *ServiceCtx {
//	once.Do(func() {
//		serviceCtx, _ = NewServiceContext()
//	})
//	return serviceCtx
//}

type Resource struct {
	Mux          core.Mux
	Logger       *zap.Logger
	Db           mysql.Repo
	Cache        redis.Repo
	Interceptors interceptor.Interceptor
	CronServer   cron.Server
}
type Server struct {
	Mux        core.Mux
	Db         mysql.Repo
	Cache      redis.Repo
	CronServer cron.Server
}

func NewHTTPServer(logger *zap.Logger, cronLogger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	r := new(Resource)
	r.Logger = logger

	openBrowserUri := config.ProjectDomain + config.ProjectPort

	//_, ok := file.IsExists(config.ProjectInstallMark)
	//if !ok { // 未安装
	//	openBrowserUri += "/install"
	//} else { // 已安装

	// 初始化 DB
	dbRepo, err := mysql.New()
	if err != nil {
		logger.Fatal("new db err", zap.Error(err))
	}
	r.Db = dbRepo

	// 初始化 Cache
	cacheRepo, err := redis.New()
	if err != nil {
		logger.Fatal("new cache err", zap.Error(err))
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

	mux, err := core.New(logger,
		core.WithEnableOpenBrowser(openBrowserUri),
		core.WithEnableCors(),
		core.WithEnableRate(),
		core.WithAlertNotify(alert.NotifyHandler(logger)),
		core.WithRecordMetrics(metrics.RecordHandler(logger)),
	)

	if err != nil {
		panic(err)
	}

	r.Mux = mux
	r.Interceptors = interceptor.New(logger, r.Cache, r.Db)

	// 设置 API 路由
	SetApiRouter(r)

	s := new(Server)
	s.Mux = mux
	s.Db = r.Db
	s.Cache = r.Cache
	s.CronServer = r.CronServer

	return s, nil
}
