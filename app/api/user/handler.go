package user

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/redis"
	"FuguBackend/app/router"
	"FuguBackend/app/services/user"
	"FuguBackend/config"
	"FuguBackend/pkg/hash"

	"go.uber.org/zap"
)

//var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Login 管理员登录
	// @Tags API.admin
	// @Router /api/login [post]
	Login() core.HandlerFunc

	// Detail 个人信息
	// @Tags API.admin
	// @Router /api/admin/info [get]
	Detail() core.HandlerFunc

	// ModifyPersonalInfo 修改个人信息
	// @Tags API.admin
	// @Router /api/admin/modify_personal_info [patch]
	ModifyPersonalInfo() core.HandlerFunc

	// Create 新增管理员
	// @Tags API.admin
	// @Router /api/admin [post]
	Create() core.HandlerFunc

	// List 管理员列表
	// @Tags API.admin
	// @Router /api/admin [get]
	List() core.HandlerFunc

	// Delete 删除管理员
	// @Tags API.admin
	// @Router /api/admin/{id} [delete]
	Delete() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       redis.Repo
	hashids     hash.Hash
	userService user.Service
}

func New(r *router.Resource) Handler {
	return &handler{
		logger:      r.Logger,
		cache:       r.Cache,
		hashids:     hash.New(config.Get().HashIds.Secret, config.Get().HashIds.Length),
		userService: user.New(r.Db, r.Cache, r.Logger),
	}
}

func (h *handler) i() {}
