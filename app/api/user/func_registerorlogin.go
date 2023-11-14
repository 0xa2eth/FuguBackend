package user

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/password"
	"FuguBackend/app/pkg/validation"
	"FuguBackend/app/proposal"
	"FuguBackend/app/repository/redis"
	"FuguBackend/app/services/user"
	"FuguBackend/config"
	"FuguBackend/pkg/snowflake"
	"errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

type registerOrLoginRequest struct {
	//Address        string `json:"address,omitempty"`
	TwitterID      string `json:"twitterID,omitempty" binding:"required"`
	TwitterAvatar  string `json:"twitterAvatar,omitempty" binding:"required"`
	TwitterName    string `json:"twitterName,omitempty" binding:"required"`
	InvitationCode string `json:"invitationCode,omitempty"`
}

type registerOrLoginResponse struct {
	UserID string `json:"UserID"`
}

// RegisterOrLogin 注册或登陆
// @Summary 注册或登陆
// @Description 注册或登陆
// @Tags API.user
// @Accept application/json
// @Produce json
// @Param Request body registerOrLoginRequest true "请求信息"
// @Success 200 {object} registerOrLoginResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/twitterlogin [post]
func (h *handler) RegisterOrLogin() core.HandlerFunc {
	return func(c core.Context) {
		req := new(registerOrLoginRequest)
		res := new(registerOrLoginResponse)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		//  如果注册过 走login 逻辑 颁发一个令牌 如果没注册过 走注册create逻辑
		search := &user.SearchOneData{TwitterID: req.TwitterID}
		users, err := h.userService.Detail(c, search)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			// 有错误 ,没找到  -> 注册

			genID, _ := snowflake.GenID()
			Avatar := config.UserPortraitBaseUrl + GetDefaultImage()
			createData := &user.CreateUserData{
				Address: "",
				UserID:  int64(genID),
				// 随机名称，和 默认bio，头像
				NickName: password.GetRandomFishName(),
				Bios:     config.DefaultBio,
				Avatar:   Avatar,
				// 推特数据
				TwitterID:     req.TwitterID,
				TwitterName:   req.TwitterName,
				TwitterAvatar: req.TwitterAvatar,
				//InvitedBy:     req.InvitationCode,
			}
			if req.InvitationCode != "" {
				HashID, err := h.cache.Get(req.InvitationCode)
				if err != nil {
					h.logger.Error("cache get invitecode failed...", zap.Error(err))
				}
				ids, _ := h.hashids.HashidsDecode(HashID)
				createData.InvitedBy = ids[0]
			}
			id, err := h.userService.Create(c, createData)
			if err != nil {
				h.logger.Info("err:", zap.Error(err))
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.AdminCreateError,
					code.Text(code.AdminCreateError)).WithError(err),
				)
				return
			}

			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(id)})
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}
			jwt, err3 := password.GenerateJWT(hashId)
			if err3 != nil {
				h.logger.Fatal(" Failed to generate auth token!!! ", zap.Error(err))
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}
			// 用户信息
			sessionUserInfo := &proposal.SessionUserInfo{
				UserID: hashId,
				//UserName:
			}

			// 将用户信息记录到 Redis 中
			err = h.cache.Set(config.RedisKeyPrefixLoginUser+jwt, string(sessionUserInfo.Marshal()), config.LoginSessionTTL, redis.WithTrace(c.Trace()))
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.AdminLoginError,
					code.Text(code.AdminLoginError)).WithError(err),
				)
				return
			}
			c.SetHeader(config.HeaderSignToken, jwt)
			// 启用浏览器的内置XSS保护机制，并阻止页面加载时检测到XSS攻击
			c.SetHeader("X-XSS-Protection", "1; mode=block")
			res.UserID = hashId
			c.Payload(res)
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			// 有错误 internal error
			c.Payload(err)
		}
		// 没错误, 找到, 登录逻辑
		hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(users.Id)})
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsEncodeError,
				code.Text(code.HashIdsEncodeError)).WithError(err),
			)
			return
		}
		jwt, err2 := password.GenerateJWT(hashId)
		if err2 != nil {
			h.logger.Fatal(" Failed to generate auth token!!! ", zap.Error(err2))
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsEncodeError,
				code.Text(code.HashIdsEncodeError)).WithError(err),
			)
			return
		}
		c.SetHeader(config.HeaderSignToken, jwt)
		// 启用浏览器的内置XSS保护机制，并阻止页面加载时检测到XSS攻击
		c.SetHeader("X-XSS-Protection", "1; mode=block")
		res.UserID = hashId
		c.Payload(res)
	}
}
func GetDefaultImage() (suffix string) {
	rand.Seed(time.Now().UnixNano())
	tag := rand.Intn(7)
	switch tag {
	case 1:
		suffix = "01.png"
	case 2:
		suffix = "02.png"
	case 3:
		suffix = "03.png"
	case 4:
		suffix = "04.png"
	case 5:
		suffix = "05.png"
	case 6:
		suffix = "06.png"
	default:
		suffix = "06.png"
	}
	return
}
