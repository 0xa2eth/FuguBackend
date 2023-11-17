package cave

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/twittersvc"
	"FuguBackend/app/repository/cron"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
	"FuguBackend/app/router/interceptor"
	"FuguBackend/app/services/cave"
	"FuguBackend/config"
	"FuguBackend/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// SecretList 洞穴内秘密列表
	// @Tags API.cave
	// @Router /api/cave/:CaveID [get]
	SecretList() core.HandlerFunc

	// Top top5洞穴
	// @Tags API.cave
	// @Router /api/cave/top [get]
	Top() core.HandlerFunc

	// RecommendCave 推荐的洞穴
	// @Tags API.cave
	// @Router /api/cave/recommend [get]
	RecommendCave() core.HandlerFunc

	// VerifyTask 效验洞穴任务是否完成
	// @Tags API.cave
	// @Router /api/cave/verifytask [get]
	VerifyTask() core.HandlerFunc

	MayIView() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       redis.Repo
	hashids     hash.Hash
	caveService cave.Service
}

func New(r *Resource) Handler {
	return &handler{
		logger:      r.Logger,
		cache:       r.Cache,
		hashids:     hash.New(config.Get().HashIds.Secret, config.Get().HashIds.Length),
		caveService: cave.New(r.Db, r.Cache, r.Logger, r.TwitterServer),
	}
}

func (h *handler) i() {}

type Resource struct {
	Mux           core.Mux
	Logger        *zap.Logger
	Db            mysql.Repo
	Cache         redis.Repo
	Interceptors  interceptor.Interceptor
	CronServer    cron.Server
	TwitterServer twittersvc.TwitterServiceMaster
}
