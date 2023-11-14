package secret

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/pagination"
	"FuguBackend/app/pkg/twittersvc"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/mysql/secrets"
	"FuguBackend/app/repository/redis"
	"FuguBackend/pkg/hash"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(core.Context, *CreateSecretData) (int64, error)

	List(c core.Context, InnerID int, pageNum, pageSize int, hashFunc hash.Hash) (pagination.PageInfo, error)

	Complaint() error

	BuildRes(c core.Context, raw []secrets.Secrets, hashFunc hash.Hash, viewAble bool) (list []Secret)
}

type service struct {
	db     mysql.Repo
	cache  redis.Repo
	logger *zap.Logger
	twSvc  twittersvc.TwitterServiceMaster
}

func New(db mysql.Repo, cache redis.Repo, logger *zap.Logger, svc twittersvc.TwitterServiceMaster) Service {

	return &service{
		db:     db,
		cache:  cache,
		logger: logger,
		twSvc:  svc,
	}
}

func (s *service) i() {}
