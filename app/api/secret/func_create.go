package secret

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/services/secret"
	"errors"
	"net/http"
)

type createRequest struct {
	Content   string   `json:"content,omitempty"`
	Images    []string `json:"images,omitempty"`
	ViewLevel int      `json:"viewLevel,omitempty"`
}

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
	return func(c core.Context) {
		userID := c.Param("UserID")
		if userID == "" {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(errors.New("path param not found")),
			)
			return
		}
		req := new(createRequest)

		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		s := new(secret.CreateSecretData)
		s.Content = req.Content
		s.ViewLevel = req.ViewLevel
		s.Images = req.Images

	}
}
