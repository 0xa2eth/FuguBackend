package user

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/mysql/users"
)

type ModifyData struct {
	Avatar   string `json:"avatar,omitempty"`
	NickName string `json:"nickName,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

func (s *service) Modify(ctx core.Context, id int64, modifyData *ModifyData) (err error) {

	// 1 新建洞穴 转发一个推特
	//content := config.TwitterPostPreFix + modifyData.NickName + config.TwitterPostSuffix
	//postUrl, err := s.twSvc.Post(ctx, content)
	//s.logger.Info(fmt.Sprintf("===========postUrl:%v ==============", postUrl))
	//
	//if err != nil {
	//	s.logger.Error("twitter a post failed... ", zap.Error(err))
	//	return err
	//}
	//
	//_, tweetIDstr, _ := s.twSvc.GetTweetIDByUrl(postUrl)
	//retweetUrl := config.RetweetPrefix + tweetIDstr
	// 2 入库
	data := map[string]interface{}{
		"nick_name":      modifyData.NickName,
		"avatar":         modifyData.Avatar,
		"bios":           modifyData.Bio,
		"caveretweeturl": "retweetUrl",
	}
	qb := users.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
