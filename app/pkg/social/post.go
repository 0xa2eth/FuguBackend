package social

import (
	"context"
	"fmt"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
	"net/url"
)

func PostATweet_ANA() {
	client := buildANAClient()
	tweet, err := client.PostTweet(url.QueryEscape("Hello, Twitter,anaconda!"), nil)
	if err != nil {
		fmt.Println("Failed to post tweet:", err)
		return
	}

	// 输出发布的推文
	fmt.Println("Tweet posted successfully!")
	fmt.Println("Tweet ID:", tweet.IdStr)
	fmt.Println("Tweet Text:", tweet.FullText)
}
func PostATweet_twi() {
	client := buildTWIClient()
	p := &types.CreateInput{
		Text: gotwi.String("This is a test tweet with poll. gotwi"),
		Poll: &types.CreateInputPoll{
			DurationMinutes: gotwi.Int(5),
			Options: []string{
				"Cyan",
				"Magenta",
				"Yellow",
				"Key plate",
			},
		},
	}

	res, err := managetweet.Create(context.Background(), client, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
}
