package user

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql/users"
	"FuguBackend/config"
	"go.uber.org/zap"
	"time"
)

type CreateUserData struct {
	UserID        int64  `json:"userid" binding:"-"`
	Address       string `json:"address" binding:"-"`
	TwitterID     string `json:"twitterID" binding:"required"`
	TwitterName   string `json:"twitterName" binding:"required"`
	TwitterAvatar string `json:"twitterAvatar" binding:"-"`
}

func (s *service) Create(ctx core.Context, adminData *CreateUserData) (id int64, err error) {
	model := users.NewModel()
	model.TwitterName = adminData.TwitterName
	model.DeletedAt = time.Now()

	//id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	//if err != nil {
	//	return 0, err
	//}
	err = s.db.GetDbW().Create(model).Error
	if err != nil {
		config.Logger.Error("", zap.Error(err))
	}
	return
}
