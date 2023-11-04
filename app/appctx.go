package app

import (
	"FuguBackend/config"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
	"sync"
)

type ServiceCtx struct {
	C   *config.Config
	Db  *gorm.DB
	Rds *redis.Repo
	Aws *s3.S3
}

var serviceCtx *ServiceCtx
var once sync.Once

func NewServiceContext() (*ServiceCtx, error) {
	//初始化日志
	config.LoadConfig()
	//初始化mysql
	initMsql, _ := mysql.New()
	db, err := initMsql.NewMysqlDB()
	if err != nil {
		config.Logger.Error("INIT MYSQL CONFIG FAILED", zap.Error(err))
		return nil, err
	}
	//初始化redis
	rds := redis.InitRds(config.Conf.RedisC.RedisHost, config.Conf.RedisC.Password, 1)

	//初始化aws
	awsClient, err := awsc.AWSInit()
	if err != nil {
		config.Logger.Error("INIT aws CONFIG FAILED", zap.Error(err))
		return nil, err
	}
	//初始化etcd
	etcdClient, err := etcd.ETCDInit()

	serviceCtxs := NewServerCtx(WithDB(db), WithRds(rds), WithAws(awsClient), WithEtcd(etcdClient))
	serviceCtxs.C = config.Conf
	return serviceCtxs, nil
}

func GetServiceCtx() *ServiceCtx {
	once.Do(func() {
		serviceCtx, _ = NewServiceContext()
	})
	return serviceCtx
}
