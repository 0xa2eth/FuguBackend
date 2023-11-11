package user

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/cron"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
	"FuguBackend/app/router/interceptor"
	"FuguBackend/app/services/user"
	"FuguBackend/config"
	"FuguBackend/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// RegisterOrLogin 注册或登陆
	// @Tags API.user
	// @Router /api/user/twitterlogin [post]
	RegisterOrLogin() core.HandlerFunc

	// UserInfo 用户（洞穴非秘密部分）个人信息
	// @Tags API.user
	// @Router /api/user/:UserID [get]
	UserInfo() core.HandlerFunc

	// ModifyInfo 修改个人（洞穴）信息
	// @Tags API.user
	// @Router /api/user/:UserID [put]
	ModifyInfo() core.HandlerFunc

	// GenInviteCode 生成邀请码
	// @Tags API.user
	// @Router /api/user/invitecode [get]
	GenInviteCode() core.HandlerFunc

	// Logout 登出
	// @Tags API.user
	// @Router /api/user/logout [get]
	Logout() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       redis.Repo
	hashids     hash.Hash
	userService user.Service
}

func New(r *Resource) Handler {
	return &handler{
		logger:      r.Logger,
		cache:       r.Cache,
		hashids:     hash.New(config.Get().HashIds.Secret, config.Get().HashIds.Length),
		userService: user.New(r.Db, r.Cache, r.Logger),
	}
}

func (h *handler) i() {}

type Resource struct {
	Mux          core.Mux
	Logger       *zap.Logger
	Db           mysql.Repo
	Cache        redis.Repo
	Interceptors interceptor.Interceptor
	CronServer   cron.Server
}
