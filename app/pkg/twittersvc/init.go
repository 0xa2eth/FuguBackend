package twittersvc

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	ConsumerKey    = "QK9pCxEpBrkMZD993E3bIfoVQ"
	ConsumerSecret = "nUcqMMQfq0XbtjcyuAVmQrimNjL8zBOV5gBvheaMC5uGxGdC8v"
	AccessToken    = "1521091305107243008-QboEaJjeaiJOhcAT4hAhabDQ5hS0wa"
	AccessSecret   = "VhiSWkGKtw1NQwuHuDiMAoLP50cXqPEExgDfkksmYZwVM"
)

var FuGuTwitterClient *twitter.Client

func buildClient() *twitter.Client {
	//// 设置OAuth1认证配置
	//config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	//token := oauth1.NewToken(AccessToken, AccessSecret)
	//httpClient := config.Client(oauth1.NoContext, token)
	//
	//// 创建Twitter客户端
	//FuGuTwitterClient = twitter.NewClient(httpClient)

	//==========================================================

	// oauth2 configures a client that uses app credentials to keep a fresh token
	fmt.Println("=======================oauth2 Client============================")
	config := &clientcredentials.Config{
		ClientID:     ConsumerKey,
		ClientSecret: ConsumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	FuGuTwitterClient = twitter.NewClient(httpClient)

	fmt.Println("Build TwitterClient Success!")
	return FuGuTwitterClient
}
