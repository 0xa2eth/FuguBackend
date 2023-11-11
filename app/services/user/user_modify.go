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
	data := map[string]interface{}{
		"nick_name": modifyData.NickName,
		"avatar":    modifyData.Avatar,
		"bios":      modifyData.Bio,
	}

	qb := users.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
