package cave

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/services/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *service) VerifyRetweetTask(c core.Context, tocheck int, target int) (bool, error) {

	// 先找到 要转发的推特的id 再找到这个人的推特,然后从他发的推特列表里找这个要转发的id
	// 一  找人
	// 二  找到要转发的推特  查数据库 找到tweetid
	userSvc := user.New(s.db, s.cache, s.logger, s.twSvc)

	person, _ := userSvc.Detail(c, &user.SearchOneData{
		Id: tocheck,
	})
	cave, _ := userSvc.Detail(c, &user.SearchOneData{
		Id: target,
	})

	tweetID, _, _ := s.twSvc.GetTweetIDByUrl(cave.Caveretweeturl)

	isFind, err := s.twSvc.FindSBReTweetByTweetID(c, person.TwitterName, int(tweetID))
	if err != nil {
		s.logger.Error("FindSBReTweetByTweetID failed ...", zap.Error(err))
		return false, err
	}
	if isFind {

		err = s.db.GetDbW().Table("users").
			Where("id = ?", tocheck).
			Update("earnedpoint", gorm.Expr("earnedpoint + ?", 1)).Error
		err = s.db.GetDbW().Table("users").
			Where("id = ?", target).
			Update("cavepoint", gorm.Expr("cavepoint + ?", 1)).Error
	}
	return isFind, err
}
