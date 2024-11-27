package comment

import (
	"net/http"

	"forum/app/comment/api/internal/logic/comment"
	"forum/app/comment/api/internal/svc"
	"forum/app/comment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 创建评论
func CreateCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewCreateCommentLogic(r.Context(), svcCtx)
		resp, err := l.CreateComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
