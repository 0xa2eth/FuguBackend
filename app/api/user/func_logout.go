package user

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/redis"
	"FuguBackend/config"
	"errors"
	"net/http"
)

type logoutResponse struct {
	UserID string `json:"userID"` // 用户账号
}

// Logout 用户登出
// @Summary 用户登出
// @Description 用户登出
// @Tags API.user
// @Accept application/json
// @Produce json
// @Success 200 {object} logoutResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/logout [get]
// @Security LoginToken
func (h *handler) Logout() core.HandlerFunc {
	return func(c core.Context) {
		uid := c.Param("UserID")
		if uid == "" {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("invalid token")),
			)
			return
		}
		res := new(logoutResponse)
		res.UserID = c.SessionUserInfo().UserID

		if !h.cache.Del(config.RedisKeyPrefixLoginUser+c.GetHeader(config.HeaderSignToken), redis.WithTrace(c.Trace())) {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminLogOutError,
				code.Text(code.AdminLogOutError)).WithError(errors.New("cache del err")),
			)
			return
		}

		c.Payload(res)
	}
}
