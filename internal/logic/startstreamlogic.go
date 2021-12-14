package logic

import (
	"StreamAgent/internal/cus/config"
	"StreamAgent/internal/cus/db"
	"context"
	"fmt"

	"StreamAgent/internal/svc"
	"StreamAgent/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StartStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) StartStreamLogic {
	return StartStreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartStreamLogic) StartStream(req types.StartStreamReq) (*types.StartStreamRes, error) {
	logx.Infof("start stream %v", req)

	var stream types.Stream
	stream.UserId = req.UserId
	stream.Category = req.Category
	stream.Title = req.Title
	stream.Status = "started"
	// todo df 生成 token
	stream.Token = "fakeToken"

	dbc := db.GetDBC()
	var streams []types.Stream
	dbc.Where("user_id = ?", req.UserId).Find(&streams)
	if len(streams) > 1 {
		panic("推流记录不唯一")
	} else {
		if len(streams) == 1 {
			stream.ID = streams[0].ID
		}
		dbc.Save(&stream)
	}
	return &types.StartStreamRes{
		Server: config.C.Stream.Server,
		Key:    fmt.Sprintf("/secret?token=%v", stream.Token),
	}, nil
}
