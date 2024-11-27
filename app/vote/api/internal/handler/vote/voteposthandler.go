package vote

import (
	"net/http"

	"forum/app/vote/api/internal/logic/vote"
	"forum/app/vote/api/internal/svc"
	"forum/app/vote/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 为帖子投票
func VotePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VotePostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := vote.NewVotePostLogic(r.Context(), svcCtx)
		resp, err := l.VotePost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
