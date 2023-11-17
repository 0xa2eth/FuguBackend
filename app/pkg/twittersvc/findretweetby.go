package twittersvc

import (
	"log"

	"FuguBackend/app/pkg/core"

	"github.com/dghubble/go-twitter/twitter"
)

func (s *TwitterServiceMaster) FindSBReTweetByTweetID(c core.Context, screenName string, tweetID int) (bool, error) {
	// 计数器
	count := 0

	// 设置最大的推文数目（可以设置更大的数目，但是 Twitter API 有限制）
	maxTweets := 3200

	//var allTweets []twitter.Tweet

	// 获取推文
	for count < maxTweets {
		// 设置 API 请求参数
		params := &twitter.UserTimelineParams{
			// 设置获取用户的用户名或用户ID
			ScreenName: screenName,
			// Count 设置获取推文的数量
			Count:           200,
			MaxID:           0,
			ExcludeReplies:  twitter.Bool(true),
			IncludeRetweets: twitter.Bool(true),
		}

		// 发起 API 请求
		tweets, _, err := s.xClient.Timelines.UserTimeline(params)
		if err != nil {
			log.Fatal(err)
		}

		// 如果没有更多推文可获取，则退出循环
		if len(tweets) == 0 {
			break
		}

		for _, tweet := range tweets {
			if tweet.ID == int64(tweetID) {
				return true, nil
			}
		}

		// 将获取的推文添加到 allTweets 切片中
		//allTweets = append(allTweets, tweets...)
		count += len(tweets)

		// 设置下一次 API 请求的 MaxID 参数
		params.MaxID = tweets[len(tweets)-1].ID - 1
	}

	return false, nil
}
