package twittersvc

import (
	"FuguBackend/app/pkg/core"
	"github.com/dghubble/go-twitter/twitter"
)

func (s *TwitterServiceMaster) GetFollowing(ctx core.Context, screenName string) ([]twitter.User, error) {

	return getUserFollowing(s.xClient, screenName)
}

// getUserFollowing 获取用户的following列表
func getUserFollowing(client *twitter.Client, screenName string) ([]twitter.User, error) {
	following, _, err := client.Friends.List(&twitter.FriendListParams{
		ScreenName: screenName,
	})
	if err != nil {
		return nil, err
	}
	return following.Users, nil
}
