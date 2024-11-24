package public

import (
	"net/http"

	"forum/app/user/api/internal/logic/public"
	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// check if mobile exists
func CheckMobileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckMobileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewCheckMobileLogic(r.Context(), svcCtx)
		resp, err := l.CheckMobile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
