package twittersvc

import (
	"log"

	"FuguBackend/app/pkg/core"

	"github.com/dghubble/go-twitter/twitter"
)

func (s *TwitterServiceMaster) GetPosts(ctx core.Context, screenName string, count int) ([]twitter.Tweet, error) {

	tweets, err := getUserTweets(s.xClient, screenName, count) // 获取前count条推文
	if err != nil {
		log.Println("Failed to get user tweets:", err)
	} else {
		log.Println("User tweets:")
		for _, tweet := range tweets {
			log.Println(tweet.Text)
		}
	}
	return tweets, nil
}

// 获取用户的推文
func getUserTweets(client *twitter.Client, screenName string, count int) ([]twitter.Tweet, error) {
	tweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: screenName,
		Count:      count,
	})
	if err != nil {
		return nil, err
	}
	return tweets, nil
}
