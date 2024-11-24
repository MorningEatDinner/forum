package profile

import (
	"net/http"

	"forum/app/user/api/internal/logic/profile"
	"forum/app/user/api/internal/svc"
	"forum/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// update email
func UpdateEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := profile.NewUpdateEmailLogic(r.Context(), svcCtx)
		resp, err := l.UpdateEmail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
