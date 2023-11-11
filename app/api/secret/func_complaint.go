package secret

import (
	"FuguBackend/app/pkg/core"
)

type complaintRequest struct{}

type complaintResponse struct{}

// Complaint 投诉
// @Summary 投诉
// @Description 投诉
// @Tags API.secret
// @Accept application/json
// @Produce json
// @Param Request body complaintRequest true "请求信息"
// @Success 200 {object} complaintResponse
// @Failure 400 {object} code.Failure
// @Router /api/secret/complaint/:SecretID [get]
func (h *handler) Complaint() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
