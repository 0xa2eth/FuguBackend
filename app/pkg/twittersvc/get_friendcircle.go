package twittersvc

import (
	"FuguBackend/app/pkg/core"
	"github.com/dghubble/go-twitter/twitter"
)

func (s *TwitterServiceMaster) RefreshFriendCircle(ctx core.Context, sbName string) ([]twitter.User, error) {
	var friendCircle []twitter.User

	follower, _ := s.GetFollower(ctx, sbName)
	friendCircle = append(friendCircle, follower...)

	following, _ := s.GetFollowing(ctx, sbName)
	friendCircle = append(friendCircle, following...)

	return friendCircle, nil

}
