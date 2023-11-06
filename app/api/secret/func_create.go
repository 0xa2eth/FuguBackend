package secret

import (
	"net/http"

	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/validation"
)

type createRequest struct {
	AuthorID  int      `json:"authorID"`
	Content   string   `json:"content"`
	Images    []string `json:"images"`
	Timestamp int64    `json:"timestamp"`
}

type createResponse struct{}

// Create 新增管理员
// @Summary 新增管理员
// @Description 新增管理员
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := &createRequest{}
		//res := &createResponse{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}
		//createData := &secret.Service{
		//	Address:       req.Address,
		//	Userid:        req.Userid,
		//	TwitterID:     req.TwitterID,
		//	TwitterName:   req.TwitterName,
		//	TwitterUrl:    req.TwitterUrl,
		//	TwitterAvatar: req.TwitterAvatar,
		//}
		//
		//id, err := h.userService.Create(c, createData)
		//if err != nil {
		//	c.AbortWithError(core.Error(
		//		http.StatusBadRequest,
		//		code.AdminCreateError,
		//		code.Text(code.AdminCreateError)).WithError(err),
		//	)
		//	return
		//}
		//
		//res.ID = id
		//c.Payload(res)

	}
}
