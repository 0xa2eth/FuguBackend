package twittersvc

import (
	"fmt"
	"strconv"

	"FuguBackend/app/pkg/core"

	"github.com/dghubble/go-twitter/twitter"
	"go.uber.org/zap"
)

func (s *TwitterServiceMaster) Post(ctx core.Context, content string) (string, error) {
	tweet, err := sendTweet(s.xClient, content)
	if err != nil {
		s.logger.Error("Failed to send tweet:", zap.Error(err))
		return "", err
	} else {
		tweetURL := getTweetURL(tweet.User.ScreenName, tweet.ID)
		s.logger.Info(fmt.Sprintf("Tweet URL:%v", tweetURL))
		return tweetURL, nil
	}
}

// 发送推文
func sendTweet(client *twitter.Client, tweetText string) (*twitter.Tweet, error) {
	params := &twitter.StatusUpdateParams{
		Status: tweetText,
	}
	tweet, _, err := client.Statuses.Update(tweetText, params)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

// 构建推文的URL
func getTweetURL(screenName string, tweetID int64) string {
	return "https://twitter.com/" + screenName + "/status/" + strconv.FormatInt(int64(tweetID), 10)

}
