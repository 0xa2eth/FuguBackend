package user

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql/user"
)

type CreateUserData struct {
	Username string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
	Password string // 密码
}

func (s *service) Create(ctx core.Context, adminData *CreateUserData) (id int32, err error) {
	model := user.NewModel()

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
