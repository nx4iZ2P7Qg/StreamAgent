package logic

import (
	"context"
	"encoding/json"

	"StreamAgent/internal/svc"
	"StreamAgent/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type VerifyConnectStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyConnectStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyConnectStreamLogic {
	return VerifyConnectStreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyConnectStreamLogic) VerifyConnectStream(req types.VerifyConnectStreamReq) (*types.Res, error) {
	re, err := json.Marshal(req)
	if err != nil {
		panic("json marshall error")
	}
	logx.Infof("req = %v", string(re))

	return &types.Res{}, nil
}
