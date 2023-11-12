package cave

import (
	"FuguBackend/app/pkg/core"
)

type secretListRequest struct{}

type secretListResponse struct {
	Secrets []SecretEntity `json:"secrets"`
}
type SecretEntity struct {
	Timestamp int64      `json:"timestamp,omitempty"`
	Views     int64      `json:"views,omitempty" gorm:"column:views;type:bigint"`
	SecretID  string     `json:"secretId,omitempty"`
	Content   string     `json:"content,omitempty" gorm:"column:content;type:varchar(255)"`
	Images    []string   `json:"images,omitempty" gorm:"foreignKey:SecretID"`
	Publisher AuthorInfo `json:"publisher,omitempty"`
}

type AuthorInfo struct {
	CaveID     string `json:"caveID"`
	CaveName   string `json:"caveName"`
	CaveBio    string `json:"caveBio"`
	CaveAvatar string `json:"caveAvatar"`
}

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
	return func(ctx core.Context) {

	}
}
