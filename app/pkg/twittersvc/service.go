package twittersvc

import (
	"FuguBackend/app/pkg/core"
	"github.com/dghubble/go-twitter/twitter"
)

var _ TwitterService = (*TwitterServiceMaster)(nil)

type TwitterService interface {
	i()

	Post(ctx core.Context) (err error)

	DirectMessage(ctx core.Context) (err error)

	GetFollower(ctx core.Context) (err error)
	GetFollowing(ctx core.Context) (err error)
	GetPosts(c core.Context, name string, count int) ([]twitter.Tweet, error)

	FindSBReTweetByTweetID(c core.Context, name string, tweetID int) (isFind bool, err error)

	FindIsFollowerSB(c core.Context, name string, tweetID int) (isFind bool, err error)

	GetTweetIDByUrl(url string) (int64, error)
}

type TwitterServiceMaster struct {
	xClient *twitter.Client
}

func NewTwitterServiceMaster() TwitterServiceMaster {

	return TwitterServiceMaster{
		xClient: buildClient(),
	}
}

func (s *TwitterServiceMaster) i() {}
