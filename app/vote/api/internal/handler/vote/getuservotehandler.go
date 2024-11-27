package vote

import (
	"net/http"

	"forum/app/vote/api/internal/logic/vote"
	"forum/app/vote/api/internal/svc"
	"forum/app/vote/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取用户对帖子投票的记录
func GetUserVoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserVoteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := vote.NewGetUserVoteLogic(r.Context(), svcCtx)
		resp, err := l.GetUserVote(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
