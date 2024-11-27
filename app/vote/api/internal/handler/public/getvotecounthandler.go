package public

import (
	"net/http"

	"forum/app/vote/api/internal/logic/public"
	"forum/app/vote/api/internal/svc"
	"forum/app/vote/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取帖子投票信息
func GetVoteCountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetVoteCountReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewGetVoteCountLogic(r.Context(), svcCtx)
		resp, err := l.GetVoteCount(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
