package secret

import (
	"FuguBackend/app/pkg/core"
)

type createRequest struct{}

type createResponse struct{}

// Create 发布新帖子
// @Summary 发布新帖子
// @Description 发布新帖子
// @Tags API.secret
// @Accept application/json
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/secret/:UserID [post]
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
