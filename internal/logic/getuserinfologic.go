package logic

import (
	"context"

	"StreamAgent/internal/svc"
	"StreamAgent/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserInfoLogic {
	return GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req types.GetUserInfoReq) (*types.GetUserInfoRes, error) {
	// todo df 需要添加逻辑
	return &types.GetUserInfoRes{
		UserId:   "1",
		Username: "vben",
		RealName: "Vben Admin",
		Avatar:   "https://q1.qlogo.cn/g?b=qq&nk=190848757&s=640",
		Desc:     "manager",
		Password: "123456",
		Token:    "fakeToken1",
		HomePath: "/stream/push",
		Roles: []types.RoleInfo{
			{
				RoleName: "Super Admin",
				Value:    "super",
			},
		},
	}, nil
}
