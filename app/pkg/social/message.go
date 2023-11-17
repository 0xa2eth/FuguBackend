package social

import (
	"fmt"
	"github.com/michimani/gotwi"

	dmpkg "github.com/ChimeraCoder/anaconda"
)

const (
	ConsumerKey    = "QK9pCxEpBrkMZD993E3bIfoVQ"
	ConsumerSecret = "nUcqMMQfq0XbtjcyuAVmQrimNjL8zBOV5gBvheaMC5uGxGdC8v"
	AccessToken    = "1521091305107243008-QboEaJjeaiJOhcAT4hAhabDQ5hS0wa"
	AccessSecret   = "VhiSWkGKtw1NQwuHuDiMAoLP50cXqPEExgDfkksmYZwVM"
)

func buildANAClient() *dmpkg.TwitterApi {
	dmpkg.SetConsumerKey(ConsumerKey)
	dmpkg.SetConsumerSecret(ConsumerSecret)
	return dmpkg.NewTwitterApi(AccessToken, AccessSecret)
}
func buildTWIClient() *gotwi.Client {
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           AccessToken,
		OAuthTokenSecret:     AccessSecret,
	}

	c, err := gotwi.NewClient(in)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return c
}
func DirectMessage_ANA(sb string, content string) {
	client := buildANAClient()
	recipientScreenName := sb // 收件人的 Twitter 用户名
	messageText := content    // 要发送的私信内容

	_, err := client.PostDMToScreenName(messageText, recipientScreenName)
	if err != nil {
		fmt.Println("发送私信失败:", err)
		return
	}

	fmt.Println("私信发送成功！")
}
