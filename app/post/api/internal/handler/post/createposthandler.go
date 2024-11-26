package post

import (
	"net/http"

	"forum/app/post/api/internal/logic/post"
	"forum/app/post/api/internal/svc"
	"forum/app/post/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 创建帖子
func CreatePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreatePostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := post.NewCreatePostLogic(r.Context(), svcCtx)
		resp, err := l.CreatePost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
