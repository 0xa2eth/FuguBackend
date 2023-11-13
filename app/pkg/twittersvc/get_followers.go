package twittersvc

import (
	"FuguBackend/app/pkg/core"
	"github.com/dghubble/go-twitter/twitter"
)

func (s *TwitterServiceMaster) GetFollower(ctx core.Context, screenName string) ([]twitter.User, error) {

	return getUserFollowers(s.xClient, screenName)

}

// 获取用户的follower列表
func getUserFollowers(client *twitter.Client, screenName string) ([]twitter.User, error) {
	followers, _, err := client.Followers.List(&twitter.FollowerListParams{
		ScreenName: screenName,
	})
	if err != nil {
		return nil, err
	}
	return followers.Users, nil
}
