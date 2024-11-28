package public

import (
	"forum/app/user/api/internal/logic/public"
	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// register by email
func RegisterByEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterByEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewRegisterByEmailLogic(r.Context(), svcCtx)
		resp, err := l.RegisterByEmail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
