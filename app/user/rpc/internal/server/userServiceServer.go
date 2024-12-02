// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"forum/app/user/rpc/internal/logic"
	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) CheckMobile(ctx context.Context, in *pb.CheckMobileRequest) (*pb.CheckMobileResponse, error) {
	l := logic.NewCheckMobileLogic(ctx, s.svcCtx)
	return l.CheckMobile(in)
}

func (s *UserServiceServer) GetCaptcha(ctx context.Context, in *pb.CaptchaRequest) (*pb.CaptchaResponse, error) {
	l := logic.NewGetCaptchaLogic(ctx, s.svcCtx)
	return l.GetCaptcha(in)
}

func (s *UserServiceServer) GetMobileCode(ctx context.Context, in *pb.GetMobileCodeRequest) (*pb.GetMobileCodeResponse, error) {
	l := logic.NewGetMobileCodeLogic(ctx, s.svcCtx)
	return l.GetMobileCode(in)
}

func (s *UserServiceServer) GetEmailCode(ctx context.Context, in *pb.GetEmailCodeRequest) (*pb.GetEmailCodeResponse, error) {
	l := logic.NewGetEmailCodeLogic(ctx, s.svcCtx)
	return l.GetEmailCode(in)
}


func (s *UserServiceServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServiceServer) RegisterByEmail(ctx context.Context, in *pb.RegisterByEmailRequest) (*pb.RegisterByEmailResponse, error) {
	l := logic.NewRegisterByEmailLogic(ctx, s.svcCtx)
	return l.RegisterByEmail(in)
}


func (s *UserServiceServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServiceServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}

func (s *UserServiceServer) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	l := logic.NewRefreshTokenLogic(ctx, s.svcCtx)
	return l.RefreshToken(in)
}

func (s *UserServiceServer) GetUserDetail(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	l := logic.NewGetUserDetailLogic(ctx, s.svcCtx)
	return l.GetUserDetail(in)
}

func (s *UserServiceServer) UpdateUserInfo(ctx context.Context, in *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	l := logic.NewUpdateUserInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserInfo(in)
}

func (s *UserServiceServer) UpdateMobile(ctx context.Context, in *pb.UpdateMobileRequest) (*pb.UpdateMobileResponse, error) {
	l := logic.NewUpdateMobileLogic(ctx, s.svcCtx)
	return l.UpdateMobile(in)
}

func (s *UserServiceServer) UpdateEmail(ctx context.Context, in *pb.UpdateEmailRequest) (*pb.UpdateEmailResponse, error) {
	l := logic.NewUpdateEmailLogic(ctx, s.svcCtx)
	return l.UpdateEmail(in)
}

func (s *UserServiceServer) UpdatePassword(ctx context.Context, in *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	l := logic.NewUpdatePasswordLogic(ctx, s.svcCtx)
	return l.UpdatePassword(in)
}


func (s *UserServiceServer) GetUserList(ctx context.Context, in *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	l := logic.NewGetUserListLogic(ctx, s.svcCtx)
	return l.GetUserList(in)
}
