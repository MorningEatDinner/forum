package vote

import (
	"net/http"

	"forum/app/vote/api/internal/logic/vote"
	"forum/app/vote/api/internal/svc"
	"forum/app/vote/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 撤销用户对帖子的投票
func RevokeVoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RevokeVoteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := vote.NewRevokeVoteLogic(r.Context(), svcCtx)
		resp, err := l.RevokeVote(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
