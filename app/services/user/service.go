package user

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/mysql/users"
	"FuguBackend/app/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, adminData *CreateUserData) (id int64, err error)

	Detail(ctx core.Context, searchOneData *SearchOneData) (info *users.Users, err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
	// todo ...
	// logger config snowflake
}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
