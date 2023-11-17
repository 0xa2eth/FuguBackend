package secret

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/pagination"
	"FuguBackend/app/pkg/twittersvc"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/mysql/secrets"
	"FuguBackend/app/repository/redis"
	"FuguBackend/pkg/hash"

	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(core.Context, string, *CreateSecretData) (int, error)

	List(c core.Context, InnerID int, pageNum, pageSize int, hashFunc hash.Hash) (pagination.PageInfo, error)

	Complaint() error

	BuildNormalSecretsRes(c core.Context, raw []secrets.Secret, hashFunc hash.Hash, viewAble bool) []SecretRes
	GetExtro(c core.Context, siteFriendIds []int, hashFunc hash.Hash, InnerID int) []SecretRes
	GetPostsByAuthorID(c core.Context, id int) (list []*secrets.Secrets, err error)
}

type service struct {
	db     mysql.Repo
	cache  redis.Repo
	logger *zap.Logger
	twSvc  twittersvc.TwitterServiceMaster
}

func New(db mysql.Repo, cache redis.Repo, logger *zap.Logger, svc twittersvc.TwitterServiceMaster) Service {

	return &service{
		db:     db,
		cache:  cache,
		logger: logger,
		twSvc:  svc,
	}
}

func (s *service) i() {}
