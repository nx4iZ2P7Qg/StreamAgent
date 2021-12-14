package logic

import (
	"StreamAgent/internal/cus/db"
	"StreamAgent/internal/svc"
	"StreamAgent/internal/types"
	"context"
	"errors"
	"github.com/tal-tech/go-zero/core/logx"
	"strings"
)

type VerifyPushStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyPushStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyPushStreamLogic {
	return VerifyPushStreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyPushStreamLogic) VerifyPushStream(req types.VerifyPushStreamReq) (*types.Res, error) {
	logx.Infof("req = %v", req)

	var inputToken string
	if strings.Contains(req.Param, "token") {
		// ?token=fakeToken
		inputToken = req.Param[7:]
	}

	var streams []types.Stream
	dbc := db.GetDBC()
	dbc.Where("user_id = ?", "9527").Find(&streams)
	if len(streams) != 1 {
		panic("stream 记录不唯一")
	}
	dbToken := streams[0].Token
	if inputToken != dbToken {
		return nil, errors.New("token unmatched")
	}

	return &types.Res{}, nil
}
