package user

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/services/user"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type userInfoRequest struct {
	UserID string `json:"UserID"`
}

type userInfoResponse struct {
	UserID    string `json:"userId,omitempty" gorm:"column:userid;type:bigint"`
	TicketNum int    `json:"ticketNum,omitempty" gorm:"column:ticketnum;type:bigint"`
	CaveFans  int    `json:"caveFans,omitempty" gorm:"column:cavefans;type:bigint"`
	//TwitterFans int    `json:"twitterFans,omitempty" gorm:"column:twitterfans;type:bigint"`
	//LastLogin   int    `json:"lastLogin,omitempty" gorm:"column:lastlogin;type:bigint"`
	//RegisTime   int    `json:"regisTime,omitempty" gorm:"column:registime;type:bigint"`
	EarnedPoint int `json:"earnedPoint,omitempty" gorm:"column:earned_point;type:bigint"`
	CavePoint   int `json:"CavePoint,omitempty" gorm:"column:cave_point;type:bigint"`
	Views       int `json:"views,omitempty" gorm:"column:views;type:bigint"`
	//Tag            int    `json:"tag" gorm:"column:tag;type:int"`
	NickName string `json:"nickName,omitempty" gorm:"column:nick_name;type:varchar(255)"`
	Bios     string `json:"bios,omitempty" gorm:"column:bios;type:varchar(255)"`
	Avatar   string `json:"avatar,omitempty" gorm:"column:avatar;type:varchar(255)"`
	//Address  string `json:"address,omitempty" gorm:"column:address;type:varchar(255)"`
	//TwitterID      string `json:"twitterID,omitempty" gorm:"column:twitter_id;type:varchar(255)"`
	//TwitterAvatar  string `json:"twitterAvatar,omitempty" gorm:"column:twitter_avatar;type:varchar(255)"`
	//TwitterName    string `json:"twitterName,omitempty" gorm:"column:twitter_name;type:varchar(255)"`
	//CaveReTweetUrl string `json:"CaveReTweetUrl" gorm:"column:caveretweeturl;type:varchar(255)"`
	NumberOfPosts int            `json:"numberOfPosts,omitempty"`
	FollowedCaves []FollowedCave `json:"followedCaves,omitempty"`
}
type FollowedCave struct {
	CaveID     string `json:"caveID"`
	CaveName   string `json:"caveName"`
	CaveAvatar string `json:"caveAvatar"`
}

// UserInfo 用户（洞穴非秘密部分）个人信息
// @Summary 用户（洞穴非秘密部分）个人信息
// @Description 用户（洞穴非秘密部分）个人信息
// @Tags API.user
// @Accept application/json
// @Produce json
// @Param Request body userInfoRequest true "请求信息"
// @Success 200 {object} userInfoResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/:UserID [get]
func (h *handler) UserInfo() core.HandlerFunc {
	return func(ctx core.Context) {
		receive := struct {
			ID string `uri:"UserID" binding:"required"`
		}{}
		err := ctx.ShouldBindURI(&receive)
		if err != nil {
			h.logger.Error("", zap.Error(err))
		}

		InnerID, err := h.hashids.HashidsDecode(receive.ID)
		fmt.Println(InnerID)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsEncodeError,
				code.Text(code.HashIdsEncodeError)).WithError(err),
			)
			return
		}

		searchOneData := new(user.SearchOneData)

		searchOneData.Id = InnerID[0]

		info, err := h.userService.Detail(ctx, searchOneData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminDetailError,
				code.Text(code.AdminDetailError)).WithError(err),
			)
			return
		}

		//_, err = h.cache.Get(config.RedisKeyPrefixLoginUser+password.GenerateLoginToken(searchOneData.Id)+":menu", redis.WithTrace(ctx.Trace()))
		//if err != nil {
		// ctx.AbortWithError(core.Error(
		//    http.StatusBadRequest,
		//    code.AdminDetailError,
		//    code.Text(code.AdminDetailError)).WithError(err),
		// )
		// return
		//}

		fmt.Println("info : ", info)
		res := userInfoResponse{
			UserID:      receive.ID,
			TicketNum:   int(info.Ticketnum),
			CaveFans:    int(info.Cavefans),
			EarnedPoint: int(info.EarnedPoint),
			CavePoint:   int(info.CavePoint),
			Views:       int(info.Views),
			NickName:    info.NickName,
			Bios:        info.Bios,
			Avatar:      info.Avatar,
		}

		var num int
		// todo
		res.NumberOfPosts = num
		ctx.Payload(res)
	}
}
