package user

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, adminData *CreateUserData) (id int32, err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
