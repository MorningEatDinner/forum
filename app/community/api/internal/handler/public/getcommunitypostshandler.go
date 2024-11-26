package public

import (
	"net/http"

	"forum/app/community/api/internal/logic/public"
	"forum/app/community/api/internal/svc"
	"forum/app/community/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// get all posts in a community
func GetCommunityPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCommunityPostsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewGetCommunityPostsLogic(r.Context(), svcCtx)
		resp, err := l.GetCommunityPosts(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
