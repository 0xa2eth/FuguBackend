package user

import (
	"net/http"

	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/validation"
	"FuguBackend/app/services/user"
)

type createRequest struct {
	Address       string `json:"address" binding:"-"`
	Userid        string `json:"userid" binding:"-"`
	TwitterID     string `json:"twitterID" binding:"required"`
	TwitterName   string `json:"twitterName" binding:"required"`
	TwitterUrl    string `json:"twitterUrl" binding:"-"`
	TwitterAvatar string `json:"twitterAvatar" binding:"-"`
}

type createResponse struct {
	ID        int32  `json:"ID,omitempty"`
	NftNum    int32  `json:"nftNum,omitempty"`
	FtNum     int32  `json:"ftNum,omitempty"`
	TicketNum int32  `json:"ticketNum,omitempty"`
	NickName  string `json:"nickName,omitempty"`
	Bios      string `json:"bios,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

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
		req := new(createRequest)
		res := new(createResponse)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		createData := &user.CreateUserData{
			Address:       req.Address,
			Userid:        req.Userid,
			TwitterID:     req.TwitterID,
			TwitterName:   req.TwitterName,
			TwitterUrl:    req.TwitterUrl,
			TwitterAvatar: req.TwitterAvatar,
		}

		id, err := h.userService.Create(c, createData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminCreateError,
				code.Text(code.AdminCreateError)).WithError(err),
			)
			return
		}

		res.ID = id
		c.Payload(res)
	}
}
