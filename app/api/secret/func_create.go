package secret

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/services/secret"
	"errors"
	"go.uber.org/zap"
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
		//userID := c.Param("UserID")
		//if userID == "" {
		//	c.AbortWithError(core.Error(
		//		http.StatusBadRequest,
		//		code.ParamBindError,
		//		code.Text(code.ParamBindError)).WithError(errors.New("path param not found")),
		//	)
		//	return
		//}
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
		req := new(createRequest)

		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		innerID, _ := h.hashids.HashidsDecode(hashID)
		s := new(secret.CreateSecretData)
		s.Content = req.Content
		s.ViewLevel = req.ViewLevel
		s.Images = req.Images
		s.AuthorID = innerID[0]
		//s.SecretID, _ = snowflake.GenID()

		sid, err := h.secretService.Create(c, hashID, s)
		if err != nil {
			h.logger.Info("err:", zap.Error(err))
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SecretCreateError,
				code.Text(code.SecretCreateError)).WithError(err),
			)
			return
		}
		hashID, _ = h.hashids.HashidsEncode([]int{int(sid)})

		c.Payload(hashID)
	}
}
