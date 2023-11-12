package twittersvc

import (
	"FuguBackend/app/pkg/core"
	"github.com/dghubble/go-twitter/twitter"
	"log"
)

func (s *TwitterServiceMaster) GetFollowing(ctx core.Context) (err error) {
	// 获取用户的following列表
	following, err := getUserFollowing(s.xClient, "screenName")
	if err != nil {
		log.Println("Failed to get user following:", err)
	} else {
		log.Println("User following:")
		for _, user := range following {
			log.Println(user.ScreenName)
		}
	}
	return nil
}

// 获取用户的following列表
func getUserFollowing(client *twitter.Client, screenName string) ([]twitter.User, error) {
	following, _, err := client.Friends.List(&twitter.FriendListParams{
		ScreenName: screenName,
	})
	if err != nil {
		return nil, err
	}
	return following.Users, nil
}
