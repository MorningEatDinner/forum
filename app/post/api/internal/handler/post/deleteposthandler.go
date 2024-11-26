package post

import (
	"net/http"

	"forum/app/post/api/internal/logic/post"
	"forum/app/post/api/internal/svc"
	"forum/app/post/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除帖子
func DeletePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeletePostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := post.NewDeletePostLogic(r.Context(), svcCtx)
		resp, err := l.DeletePost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
