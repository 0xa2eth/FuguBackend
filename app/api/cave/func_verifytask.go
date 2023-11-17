package cave

import (
	"errors"
	"net/http"

	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/validation"
)

type verifyTaskRequest struct {
	CaveID string `uri:"CaveID" binding:"required"`
	Type   int    `uri:"Type" binding:"required"`
}

type verifyTaskResponse struct {
	verify bool `json:"verify"`
}

// VerifyTask 验证任务是否完成
// @Summary 验证任务是否完成
// @Description 验证任务是否完成
// @Tags API.cave
// @Accept application/json
// @Produce json
// @Param Request body verifyTaskRequest true "请求信息"
// @Success 200 {object} verifyTaskResponse
// @Failure 400 {object} code.Failure
// @Router /api/cave/verifytask [get]
func (h *handler) VerifyTask() core.HandlerFunc {
	return func(c core.Context) {
		req := new(verifyTaskRequest)
		res := new(verifyTaskResponse)
		if err := c.ShouldBindURI(&req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}
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
		targetID := req.CaveID
		uid, _ := h.hashids.HashidsDecode(hashID)
		targetid, _ := h.hashids.HashidsDecode(targetID)

		if req.Type == 1 {
			// 关注的
			// 一 找人
			// 二  找到要转发的推特  查数据库 找到tweetid
			complete, err := h.caveService.VerifyFollowTask(c, uid[0])
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusOK,
					code.VerifyTaskError,
					code.Text(code.VerifyTaskError)).WithError(err),
				)
				return
			}
			res.verify = complete
		} else if req.Type == 2 {
			// 转推的
			// 先找到 要转发的推特的id 再找到这个人的推特,然后从他发的推特列表里找这个要转发的id
			// 一 找人
			// 二  找到要转发的推特  查数据库 找到tweetid
			complete, err := h.caveService.VerifyRetweetTask(c, uid[0], targetid[0])
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusOK,
					code.VerifyTaskError,
					code.Text(code.VerifyTaskError)).WithError(err),
				)
				return
			}
			res.verify = complete
		} else {
			c.AbortWithError(core.Error(
				http.StatusOK,
				code.VerifyTaskError,
				code.Text(code.VerifyTaskError)).WithError(errors.New("invalid verify type")),
			)
		}

		c.Payload(res)

	}
}
