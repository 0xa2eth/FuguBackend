package cave

import (
	"FuguBackend/app/services/cave"
	"net/http"

	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/pagination"
	"FuguBackend/app/pkg/validation"
	"FuguBackend/pkg/errors"
)

type secretListRequest struct {
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
	Order    string `form:"order"`
}
type secretListUri struct {
	CaveID string `uri:"CaveID"`
}
type secretListResponse struct {
	Data pagination.PageInfo `json:"secrets"`
}

//type SecretEntity struct {
//	Timestamp int64      `json:"timestamp,omitempty"`
//	Views     int64      `json:"views,omitempty" gorm:"column:views;type:bigint"`
//	SecretID  string     `json:"secretId,omitempty"`
//	Content   string     `json:"content,omitempty" gorm:"column:content;type:varchar(255)"`
//	Images    []string   `json:"images,omitempty" gorm:"foreignKey:SecretID"`
//	Publisher AuthorInfo `json:"publisher,omitempty"`
//}
//
//type AuthorInfo struct {
//	CaveID     string `json:"caveID"`
//	CaveName   string `json:"caveName"`
//	CaveBio    string `json:"caveBio"`
//	CaveAvatar string `json:"caveAvatar"`
//}

// SecretList 洞穴内秘密列表
// @Summary 洞穴内秘密列表
// @Description 洞穴内秘密列表
// @Tags API.cave
// @Accept application/json
// @Produce json
// @Param Request body secretListRequest true "请求信息"
// @Success 200 {object} secretListResponse
// @Failure 400 {object} code.Failure
// @Router /api/cave/:CaveID [get]
func (h *handler) SecretList() core.HandlerFunc {
	return func(c core.Context) {
		hashid, exists := c.Get("UserID")
		if !exists {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("invalid token")),
			)
			return
		}
		InnerID, err := h.hashids.HashidsDecode(hashid.(string))
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsEncodeError,
				code.Text(code.HashIdsEncodeError)).WithError(err),
			)
			return
		}
		req := new(secretListRequest)
		if err = c.ShouldBindQuery(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}
		uri := new(secretListUri)
		if err = c.ShouldBindURI(uri); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}
		var CaveSecret []cave.Secretcave
		if hashid != uri.CaveID {
			//h.caveService.ListCaveSecrets(c core.Context, InnerID, caveID int,
			//	pageNum, pageSize int, hashFunc hash.Hash) ([]Secretcave, error)

		} else {
			CaveSecret, err = h.caveService.ListMySecrets(c, req.PageNum, req.PageSize, InnerID[0], h.hashids)
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusOK,
					code.GetSecretsError,
					code.Text(code.GetSecretsError)).WithError(err),
				)
				return
			}
		}

		page := pagination.PageInfo{}
		page.PageNum = req.PageNum
		page.PageSize = req.PageSize
		page.Data = CaveSecret
		res := new(secretListResponse)
		res.Data = page
		c.Payload(res)
	}
}
