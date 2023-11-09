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

//var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 新增管理员
	// @Tags API.admin
	// @Router /api/admin [post]
	Create() core.HandlerFunc

	// List 管理员列表
	// @Tags API.admin
	// @Router /api/admin [get]
	List() core.HandlerFunc

	// Detail 个人信息
	// @Tags API.admin
	// @Router /api/admin/info [get]
	Detail() core.HandlerFunc

	// Delete 删除管理员
	// @Tags API.admin
	// @Router /api/admin/{id} [delete]
	Delete() core.HandlerFunc
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
		secretService: secret.New(r.Db, r.Cache),
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
