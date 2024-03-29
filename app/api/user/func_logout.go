package user

import (
	"errors"
	"net/http"

	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/redis"
	"FuguBackend/config"
)

type logoutResponse struct {
	UserID string `json:"userID"` // 用户账号
	Logout bool   `json:"logout"`
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
		uid, exists := c.Get("UserID")
		if uid == "" || !exists {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("invalid token")),
			)
			return
		}

		if !h.cache.Del(config.RedisKeyPrefixLoginUser+c.GetHeader(config.HeaderSignToken), redis.WithTrace(c.Trace())) {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminLogOutError,
				code.Text(code.AdminLogOutError)).WithError(errors.New("cache del err")),
			)
			return
		}
		res := new(logoutResponse)
		res.UserID = uid.(string)
		res.Logout = true

		c.Payload(res)
	}
}
