package public

import (
	"net/http"

	"forum/app/post/api/internal/logic/public"
	"forum/app/post/api/internal/svc"
	"forum/app/post/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取帖子详情
func GetPostDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewGetPostDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetPostDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
