package user

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql"
	"FuguBackend/app/repository/mysql/users"
)

type SearchOneData struct {
	Id        int // 用户ID
	TwitterID string
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *users.Users, err error) {

	qb := users.NewQueryBuilder()
	//first, err := qb.First(s.db.GetDbR().WithContext(ctx.RequestContext()))

	// 登陆的的时候 userid != 0  twitterid = "" 走这个
	if searchOneData.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, int64(searchOneData.Id))
	}
	// 注册的时候 userid = 0  twitterid != " 走这个
	if searchOneData.TwitterID != "" {
		qb.WhereTwitterId(mysql.EqualPredicate, searchOneData.TwitterID)
	}

	//if searchOneData.Nickname != "" {
	//	qb.WhereNickname(mysql.EqualPredicate, searchOneData.Nickname)
	//}
	//
	//if searchOneData.Mobile != "" {
	//	qb.WhereMobile(mysql.EqualPredicate, searchOneData.Mobile)
	//}
	//
	//if searchOneData.Password != "" {
	//	qb.WherePassword(mysql.EqualPredicate, searchOneData.Password)
	//}
	//
	//if searchOneData.IsUsed != 0 {
	//	qb.WhereIsUsed(mysql.EqualPredicate, searchOneData.IsUsed)
	//}

	info, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return info, err
}
