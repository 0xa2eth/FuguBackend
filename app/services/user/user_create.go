package user

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql/user"
)

type CreateUserData struct {
	Address       string `json:"address" binding:"-"`
	Userid        string `json:"userid" binding:"-"`
	TwitterID     string `json:"twitterID" binding:"required"`
	TwitterName   string `json:"twitterName" binding:"required"`
	TwitterUrl    string `json:"twitterUrl" binding:"-"`
	TwitterAvatar string `json:"twitterAvatar" binding:"-"`
}

func (s *service) Create(ctx core.Context, adminData *CreateUserData) (id int32, err error) {
	model := user.NewModel()

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
