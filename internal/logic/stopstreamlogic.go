package logic

import (
	"StreamAgent/internal/cus/db"
	"context"

	"StreamAgent/internal/svc"
	"StreamAgent/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StopStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStopStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) StopStreamLogic {
	return StopStreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StopStreamLogic) StopStream(req types.StopStreamReq) (*types.Res, error) {
	logx.Infof("stop stream %v", req)

	dbc := db.GetDBC()
	var streams []types.Stream
	dbc.Where("user_id = ?", req.UserId).Find(&streams)
	if len(streams) == 1 {
		streams[0].Status = "finished"
		dbc.Save(&streams[0])
	} else {
		panic("推流记录不唯一")
	}
	return &types.Res{}, nil
}
