package secret

import (
	"FuguBackend/app/pkg/core"
)

type listRequest struct{}

type listResponse struct{}

// List 可见的秘密列表
// @Summary 可见的秘密列表
// @Description 可见的秘密列表
// @Tags API.secret
// @Accept application/json
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/secret/viewable [get]
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		value, exists := c.Get("UserID")
		//
		// 先拿id  再刷新下朋友圈 于上次存cache的对比 找到新增的和去处的再根据这些去拿可见的秘密

	}
}
