package twittersvc

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/redis"
	"github.com/dghubble/go-twitter/twitter"
	"go.uber.org/zap"
)

var _ TwitterService = (*TwitterServiceMaster)(nil)

type TwitterService interface {
	i()

	Post(core.Context, string) (string, error)

	GetPosts(ctx core.Context, screenName string, count int) ([]twitter.Tweet, error)

	GetTweetIDByUrl(url string) (int, error)

	GetFollower(ctx core.Context, screenName string) ([]twitter.User, error)

	GetFollowing(ctx core.Context, screenName string) ([]twitter.User, error)

	RefreshFriendCircle(ctx core.Context, sbName string) ([]twitter.User, error)

	DiffFriendCircle(ctx core.Context, InnerID int, screenName string, ids []string) (diff FDiff, err error)

	FindSBReTweetByTweetID(core.Context, string, int) (isFind bool, err error)

	FindIsFollower(c core.Context, name string, tweetID int) (isFind bool, err error)

	DirectMessage(ctx core.Context, recipientScreenName, message string) (err error)
}

type TwitterServiceMaster struct {
	xClient *twitter.Client
	db      mysql.Repo
	cache   redis.Repo
	logger  *zap.Logger
}

func NewTwitterServiceMaster(db mysql.Repo,
	cache redis.Repo,
	logger *zap.Logger) TwitterServiceMaster {

	return TwitterServiceMaster{
		db:      db,
		cache:   cache,
		logger:  logger,
		xClient: buildClient(),
	}
}

func (s *TwitterServiceMaster) i() {}
