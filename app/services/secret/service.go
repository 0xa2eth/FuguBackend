package secret

import (
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
}

type service struct {
	db     mysql.Repo
	cache  redis.Repo
	logger *zap.Logger
}

func New(db mysql.Repo, cache redis.Repo, logger *zap.Logger) Service {

	return &service{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}

func (s *service) i() {}
