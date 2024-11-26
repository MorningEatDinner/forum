// Code generated by goctl. DO NOT EDIT.
package types

type CreatePostReq struct {
	CommunityId int64  `json:"communityId"` // 社区ID
	Title       string `json:"title"`       // 帖子标题
	Content     string `json:"content"`     // 帖子内容
}

type CreatePostResp struct {
	PostId int64 `json:"postId"` // 创建的帖子ID
}

type DeletePostReq struct {
	PostId int64 `path:"postId"` // 要删除的帖子ID
}

type DeletePostResp struct {
}

type GetPostDetailReq struct {
	PostId int64 `path:"postId"` // 帖子ID
}

type GetPostDetailResp struct {
	Post Post `json:"post"` // 帖子详情
}

type GetPostListReq struct {
	Page        int64 `form:"page,optional,default=1"`      // 页码
	PageSize    int64 `form:"pageSize,optional,default=10"` // 每页数量
	CommunityId int64 `form:"communityId,optional"`         // 按社区筛选
	AuthorId    int64 `form:"authorId,optional"`            // 按作者筛选
}

type GetPostListResp struct {
	Total int64  `json:"total"` // 总帖子数
	Posts []Post `json:"posts"` // 帖子列表
}

type Post struct {
	PostId      int64  `json:"postId"`      // 帖子ID
	AuthorId    int64  `json:"authorId"`    // 作者ID
	CommunityId int64  `json:"communityId"` // 社区ID
	Title       string `json:"title"`       // 帖子标题
	Content     string `json:"content"`     // 帖子内容
	CreateTime  string `json:"createTime"`  // 创建时间
	UpdatedTime string `json:"updatedTime"` // 更新时间
}
