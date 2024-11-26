package admin

import (
	"net/http"

	"forum/app/community/api/internal/logic/admin"
	"forum/app/community/api/internal/svc"
	"forum/app/community/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// update community details
func UpdateCommunityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCommunityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewUpdateCommunityLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCommunity(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
