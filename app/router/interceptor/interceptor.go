package interceptor

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/twittersvc"
	"FuguBackend/app/proposal"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
	"FuguBackend/app/services/user"

	"go.uber.org/zap"
)

var _ Interceptor = (*interceptor)(nil)

type Interceptor interface {
	// CheckLogin 验证是否登录
	CheckLogin(ctx core.Context) (info proposal.SessionUserInfo, err core.BusinessError)

	// CheckJWT 验证 JWT
	CheckJWT() core.HandlerFunc

	// CheckRBAC 验证 RBAC 权限是否合法
	//CheckRBAC() core.HandlerFunc

	// CheckSignature 验证签名是否合法，对用签名算法 pkg/signature
	CheckSignature() core.HandlerFunc

	// i 为了避免被其他包实现
	i()
}

type interceptor struct {
	logger *zap.Logger
	cache  redis.Repo
	db     mysql.Repo
	//authorizedService authorized.Service
	userService user.Service
}

func New(logger *zap.Logger, cache redis.Repo, db mysql.Repo, twsvc twittersvc.TwitterServiceMaster) Interceptor {
	return &interceptor{
		logger: logger,
		cache:  cache,
		db:     db,
		//authorizedService: authorized.New(db, cache),
		userService: user.New(db, cache, logger, twsvc),
	}
}

func (i *interceptor) i() {}
