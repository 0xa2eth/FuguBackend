package social

import (
	"fmt"
	dmpkg "github.com/ChimeraCoder/anaconda"
)

func buildClient() *dmpkg.TwitterApi {
	dmpkg.SetConsumerKey("YOUR_CONSUMER_KEY")
	dmpkg.SetConsumerSecret("YOUR_CONSUMER_SECRET")
	return dmpkg.NewTwitterApi("YOUR_ACCESS_TOKEN", "YOUR_ACCESS_TOKEN_SECRET")
}
func DirectMessageTo(sb string, content string) {
	client := buildClient()
	recipientScreenName := sb // 收件人的 Twitter 用户名
	messageText := content    // 要发送的私信内容

	_, err := client.PostDMToScreenName(messageText, recipientScreenName)
	if err != nil {
		fmt.Println("发送私信失败:", err)
		return
	}

	fmt.Println("私信发送成功！")
}
