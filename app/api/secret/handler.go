package secret

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
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

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:        logger,
		cache:         cache,
		hashids:       hash.New(config.Get().HashIds.Secret, config.Get().HashIds.Length),
		secretService: secret.New(db, cache),
	}
}

func (h *handler) i() {}
