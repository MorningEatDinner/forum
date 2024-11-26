// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	post "forum/app/post/api/internal/handler/post"
	public "forum/app/post/api/internal/handler/public"
	"forum/app/post/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 创建帖子
				Method:  http.MethodPost,
				Path:    "/posts",
				Handler: post.CreatePostHandler(serverCtx),
			},
			{
				// 删除帖子
				Method:  http.MethodDelete,
				Path:    "/posts/:postId",
				Handler: post.DeletePostHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/post/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 获取帖子列表
				Method:  http.MethodGet,
				Path:    "/posts",
				Handler: public.GetPostListHandler(serverCtx),
			},
			{
				// 获取帖子详情
				Method:  http.MethodGet,
				Path:    "/posts/:postId",
				Handler: public.GetPostDetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/post/v1"),
	)
}