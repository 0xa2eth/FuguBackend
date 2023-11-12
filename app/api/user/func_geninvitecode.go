package user

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/password"
	"FuguBackend/app/repository/redis"
	"FuguBackend/config"
	"errors"
	"net/http"
)

type genInviteCodeRequest struct{}

type genInviteCodeResponse struct{}

// GenInviteCode 生成邀请码
// @Summary 生成邀请码
// @Description 生成邀请码
// @Tags API.user
// @Accept application/json
// @Produce json
// @Param Request body genInviteCodeRequest true "请求信息"
// @Success 200 {object} genInviteCodeResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/invitecode [get]
func (h *handler) GenInviteCode() core.HandlerFunc {
	return func(c core.Context) {
		value, exists := c.Get("UserID")
		if !exists {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("invalid token")),
			)
			return
		}
		hashID := value.(string)
		inviteCode := password.GenInviteCode(hashID, config.DefaultInviteCodeLength)
		err := h.cache.Set(inviteCode, hashID, config.DefaultInviteCodeTTL, redis.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusInternalServerError,
				code.CacheSetError,
				code.Text(code.CacheSetError)).WithError(err),
			)
			return
		}
		c.Payload(inviteCode)
	}
}
