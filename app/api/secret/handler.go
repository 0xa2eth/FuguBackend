package secret

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/cron"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
	"FuguBackend/app/router/interceptor"
	"FuguBackend/app/services/secret"
	"FuguBackend/config"
	"FuguBackend/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 发布新帖子
	// @Tags API.secret
	// @Router /api/secret/:UserID [post]
	Create() core.HandlerFunc

	// List 可见的秘密列表
	// @Tags API.secret
	// @Router /api/secret/viewable [get]
	List() core.HandlerFunc

	// Complaint 投诉
	// @Tags API.secret
	// @Router /api/secret/complaint/:SecretID [get]
	Complaint() core.HandlerFunc
}

type handler struct {
	logger        *zap.Logger
	cache         redis.Repo
	hashids       hash.Hash
	secretService secret.Service
}

func New(r *Resource) Handler {
	return &handler{
		logger:        r.Logger,
		cache:         r.Cache,
		hashids:       hash.New(config.Get().HashIds.Secret, config.Get().HashIds.Length),
		secretService: secret.New(r.Db, r.Cache, r.Logger),
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
