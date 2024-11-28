package public

import (
	"forum/app/user/api/internal/logic/public"
	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// get email verification code
func GetEmailCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetEmailCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewGetEmailCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetEmailCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
