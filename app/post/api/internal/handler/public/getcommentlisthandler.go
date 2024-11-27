package public

import (
	"forum/app/post/api/internal/logic/public"
	"forum/app/post/api/internal/svc"
	"forum/app/post/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取评论列表
func GetCommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCommentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewGetCommentListLogic(r.Context(), svcCtx)
		resp, err := l.GetCommentList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
