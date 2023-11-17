package cave

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/twittersvc"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
	"FuguBackend/pkg/hash"

	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	VerifyRetweetTask(core.Context, int, int) (bool, error)

	VerifyFollowTask(core.Context, int) (bool, error)
	ListCaveSecrets(c core.Context, InnerID, caveID int, pageNum, pageSize int, hashFunc hash.Hash) ([]Secretcave, error)

	ListMySecrets(c core.Context, InnerID int, pageNum, pageSize int, hashFunc hash.Hash) ([]Secretcave, error)
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
