package twittersvc

import (
	"FuguBackend/app/pkg/core"
	"github.com/dghubble/go-twitter/twitter"
)

var _ TwitterService = (*TwitterServiceMaster)(nil)

type TwitterService interface {
	i()

	Post(core.Context, string) (string, error)

	GetPosts(ctx core.Context, screenName string, count int) ([]twitter.Tweet, error)

	GetTweetIDByUrl(url string) (int64, error)

	GetFollower(ctx core.Context, screenName string) ([]twitter.User, error)

	GetFollowing(ctx core.Context, screenName string) ([]twitter.User, error)

	GetFriendCircle(ctx core.Context, sbName string) ([]twitter.User, error)

	FindSBReTweetByTweetID(core.Context, string, int) (isFind bool, err error)

	FindIsFollower(c core.Context, name string, tweetID int) (isFind bool, err error)

	DirectMessage(ctx core.Context, recipientScreenName, message string) (err error)
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
