package twittersvc

import "FuguBackend/app/pkg/core"

func (s *TwitterServiceMaster) FindIsFollowing(c core.Context, name string, beFollowedName string) (isFind bool, err error) {
	// 绝大多数都是新关注者 200以内
	following, _ := s.GetFollowing(c, name, false)
	for _, people := range following {
		if people.Name == beFollowedName {
			return true, nil
		}
	}
	return false, nil

}
