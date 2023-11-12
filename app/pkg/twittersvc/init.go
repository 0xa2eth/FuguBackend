package twittersvc

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const (
	ConsumerKey    = "consumerKey"
	ConsumerSecret = "consumerSecret"
	AccessToken    = "accessToken"
	AccessSecret   = ""
)

var FuGuTwitterClient *twitter.Client

func buildClient() *twitter.Client {
	// 设置OAuth1认证配置
	config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	token := oauth1.NewToken(AccessToken, AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// 创建Twitter客户端
	FuGuTwitterClient = twitter.NewClient(httpClient)
	return FuGuTwitterClient
}
