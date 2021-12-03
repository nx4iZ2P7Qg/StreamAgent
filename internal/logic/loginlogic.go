package logic

import (
	"context"

	"StreamAgent/internal/svc"
	"StreamAgent/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginRes, error) {
	// todo df 需要添加逻辑
	return &types.LoginRes{
		Roles: []types.RoleInfo{
			{
				RoleName: "Super Admin",
				Value:    "super",
			},
		},
		UserId:   "1",
		Username: "vben",
		Token:    "fakeToken1",
		RealName: "Vben Admin",
		Desc:     "manager",
	}, nil
}
