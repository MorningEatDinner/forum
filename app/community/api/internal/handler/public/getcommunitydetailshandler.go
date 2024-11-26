package public

import (
	"net/http"

	"forum/app/community/api/internal/logic/public"
	"forum/app/community/api/internal/svc"
	"forum/app/community/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// get community details
func GetCommunityDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCommunityDetailsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewGetCommunityDetailsLogic(r.Context(), svcCtx)
		resp, err := l.GetCommunityDetails(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
