package user

import (
	"time"

	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql/users"

	"gorm.io/gorm"
)

type CreateUserData struct {
	UserID        int64  `json:"userid" binding:"-"`
	Address       string `json:"address" binding:"-"`
	NickName      string `json:"nickName,omitempty" gorm:"column:nick_name;type:varchar(255)"`
	Bios          string `json:"bios,omitempty" gorm:"column:bios;type:varchar(255)"`
	Avatar        string `json:"avatar,omitempty" gorm:"column:avatar;type:varchar(255)"`
	TwitterID     string `json:"twitterID" binding:"required"`
	TwitterName   string `json:"twitterName" binding:"required"`
	TwitterAvatar string `json:"twitterAvatar" binding:"-"`
	InvitedBy     int    `json:"invitedBy,omitempty"`
}
type modifyData struct {
}

func (s *service) Create(ctx core.Context, data *CreateUserData) (id int64, err error) {
	model := users.NewModel()
	model.TwitterName = data.TwitterName
	model.TwitterId = data.TwitterID
	model.TwitterAvatar = data.TwitterAvatar
	model.NickName = data.NickName
	model.Avatar = data.Avatar
	model.Bios = data.Bios
	model.RegisTime = time.Now().Unix()
	//todo 数据库可以改下字段名
	model.InvitedBy = int64(data.InvitedBy)
	//model.DeletedAt = time.Now()

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	err = s.db.GetDbW().Table("users").
		Where("id = ?", id).
		Update("cavepoint", gorm.Expr("cavepoint + ?", 1)).Error
	err = s.db.GetDbW().Table("users").
		Where("id = ?", data.InvitedBy).
		Update("earnedpoint", gorm.Expr("earnedpoint + ?", 1)).Error
	return
}
