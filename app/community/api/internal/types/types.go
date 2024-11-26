// Code generated by goctl. DO NOT EDIT.
package types

type Community struct {
	CommunityId   int64  `json:"communityId"`
	CommunityName string `json:"communityName"`
	Introduction  string `json:"introduction,omitempty"`
	CreateTime    string `json:"createTime"`
	UpdatedTime   string `json:"updatedTime"`
}

type CreateCommunityReq struct {
	CommunityName string `json:"communityName"`
	Introduction  string `json:"introduction,optional"`
}

type CreateCommunityResp struct {
	CommunityId int64 `json:"communityId"`
}

type GetAllCommunitiesReq struct {
	Page     int `form:"page,optional,default=1"`
	PageSize int `form:"pageSize,optional,default=10"`
}

type GetAllCommunitiesResp struct {
	Total       int64       `json:"total"`
	Communities []Community `json:"communities"`
}

type GetCommunityDetailsReq struct {
	CommunityId int64 `form:"communityId"`
}

type GetCommunityDetailsResp struct {
	Community Community `json:"community"`
}

type GetCommunityPostsReq struct {
	CommunityId int64 `form:"communityId"`
	Page        int   `form:"page,optional,default=1"`
	PageSize    int   `form:"pageSize,optional,default=10"`
}

type GetCommunityPostsResp struct {
	Total int64  `json:"total"`
	Posts []Post `json:"posts"`
}

type Post struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type UpdateCommunityReq struct {
	CommunityId   int64  `json:"communityId"`
	CommunityName string `json:"communityName,optional"`
	Introduction  string `json:"introduction,optional"`
}

type UpdateCommunityResp struct {
	CommunityId int64 `json:"communityId"`
}