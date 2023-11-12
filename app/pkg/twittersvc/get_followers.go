package twittersvc

import (
	"FuguBackend/app/pkg/core"
	"github.com/dghubble/go-twitter/twitter"
	"log"
)

func (s *TwitterServiceMaster) GetFollower(ctx core.Context) (err error) {
	// 获取用户的follower列表
	followers, err := getUserFollowers(s.xClient, "screenName")
	if err != nil {
		log.Println("Failed to get user followers:", err)
	} else {
		log.Println("User followers:")
		for _, follower := range followers {
			log.Println(follower.ScreenName)
		}
	}
	return nil
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
