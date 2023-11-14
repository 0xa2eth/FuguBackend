package cave

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/services/user"
	"FuguBackend/config"
)

func (s *service) VerifyFollowTask(c core.Context, uid int) (bool, error) {
	userSvc := user.New(s.db, s.cache, s.logger, s.twSvc)

	person, _ := userSvc.Detail(c, &user.SearchOneData{
		Id: uid,
	})
	return s.twSvc.FindIsFollowing(c, person.TwitterName, config.FuguTwitterName)
}
