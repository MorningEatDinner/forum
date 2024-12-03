package public

import (
	"net/http"

	"forum/app/user/api/internal/logic/public"
	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"forum/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// get mobile verification code
func GetMobileCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMobileCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewGetMobileCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetMobileCode(&req)
		result.HttpResult(r, w, resp, err)
	}
}
