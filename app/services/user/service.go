package user

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/twittersvc"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/mysql/users"
	"FuguBackend/app/repository/redis"

	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, userData *CreateUserData) (id int64, err error)

	Detail(ctx core.Context, searchOneData *SearchOneData) (info *users.Users, err error)

	Modify(ctx core.Context, id int64, modifyData *ModifyData) (err error)
}

type service struct {
	db     mysql.Repo
	cache  redis.Repo
	logger *zap.Logger
	twSvc  twittersvc.TwitterServiceMaster
}

func New(
	db mysql.Repo,
	cache redis.Repo,
	logger *zap.Logger,
	svc twittersvc.TwitterServiceMaster) Service {
	return &service{
		db:     db,
		cache:  cache,
		logger: logger,
		twSvc:  svc,
	}
}

func (s *service) i() {}
