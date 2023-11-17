package twittersvc

import (
	"fmt"
	"time"

	"FuguBackend/app/pkg/core"

	"github.com/dghubble/go-twitter/twitter"
)

func (s *TwitterServiceMaster) GetFollower(ctx core.Context, screenName string, isAll bool) ([]twitter.User, error) {

	return getUserFollowers(s.xClient, screenName, isAll)

}

// 获取用户的follower列表
func getUserFollowers(client *twitter.Client, screenName string, isAll bool) ([]twitter.User, error) {
	param := &twitter.FollowerListParams{
		ScreenName:          screenName,
		Count:               200,
		SkipStatus:          twitter.Bool(true),
		IncludeUserEntities: twitter.Bool(true),
	}
	followers, _, err := client.Followers.List(param)
	if err != nil {
		return nil, err
	}

	var resUsers []twitter.User
	resUsers = append(resUsers, followers.Users...)
	if isAll {
		// 获取下一页的关注对象
		for followers.NextCursor != 0 {
			param.Cursor = followers.NextCursor
			friends, httpRes, err := client.Followers.List(param)
			if err != nil {
				fmt.Println("Failed to get followers...", err)
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
