package twittersvc

import (
	"fmt"
	"time"

	"FuguBackend/app/pkg/core"

	"github.com/dghubble/go-twitter/twitter"
)

func (s *TwitterServiceMaster) GetFollowing(ctx core.Context, screenName string, isAll bool) ([]twitter.User, error) {

	return getUserFollowing(s.xClient, screenName, isAll)
}

// getUserFollowing 获取用户的following列表
func getUserFollowing(client *twitter.Client, screenName string, isAll bool) ([]twitter.User, error) {
	param := &twitter.FriendListParams{
		ScreenName:          screenName,
		Count:               200,
		SkipStatus:          twitter.Bool(true),
		IncludeUserEntities: twitter.Bool(true),
	}
	resp, _, err := client.Friends.List(param)
	if err != nil {
		return nil, err
	}

	var resUsers []twitter.User
	resUsers = append(resUsers, resp.Users...)
	if isAll {
		// 获取下一页的关注对象
		for resp.NextCursor != 0 {
			param.Cursor = resp.NextCursor
			friends, httpRes, err := client.Friends.List(param)
			if err != nil {
				fmt.Println("Failed to get following:", err)
				return nil, nil
			}
			if httpRes.Status == "429" {
				//请求次数过多，休息一m再请求  这个地方后续要优化
				time.Sleep(time.Minute)
			}
			resUsers = append(resUsers, friends.Users...)
		}
	}

	return resUsers, nil
}
