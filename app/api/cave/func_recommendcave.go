package cave

import (
	"FuguBackend/app/pkg/core"
)

type recommendCaveRequest struct{}

type recommendCaveResponse struct {
}

// RecommendCave 推荐的洞穴
// @Summary 推荐的洞穴
// @Description 推荐的洞穴
// @Tags API.cave
// @Accept application/json
// @Produce json
// @Param Request body recommendCaveRequest true "请求信息"
// @Success 200 {object} recommendCaveResponse
// @Failure 400 {object} code.Failure
// @Router /api/cave/recommend [get]
func (h *handler) RecommendCave() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
