package secret

import (
	"errors"
	"net/http"

	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/validation"
)

type listRequest struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
	Order    int `form:"order"`
}

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
		if !exists {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("invalid token")),
			)
			return
		}
		HashID, err := h.hashids.HashidsDecode(value.(string))
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsEncodeError,
				code.Text(code.HashIdsEncodeError)).WithError(err),
			)
			return
		}
		req := new(listRequest)
		if err := c.ShouldBindQuery(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		list, err := h.secretService.List(c, HashID[0], req.PageNum, req.PageSize, h.hashids)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusOK,
				code.AdminCreateError,
				code.Text(code.AdminCreateError)).WithError(err),
			)
			return
		}
		c.Payload(list)

	}
}
