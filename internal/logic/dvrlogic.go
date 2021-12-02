package logic

import (
	"StreamAgent/internal/cus/config"
	"StreamAgent/internal/cus/db"
	"StreamAgent/internal/cus/minio"
	"StreamAgent/internal/cus/ssh"
	"context"
	"fmt"
	"path/filepath"
	"time"

	"StreamAgent/internal/svc"
	"StreamAgent/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DvrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDvrLogic(ctx context.Context, svcCtx *svc.ServiceContext) DvrLogic {
	return DvrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DvrLogic) Dvr(req types.DvrRequest) (*types.DvrResponse, error) {
	logx.Infof("req = %s\n", req)

	// todo df contentName 来自于用户开始直播的节目名称，需要从数据库中取出
	contentName := "contentName"

	// 挂载到 pod 的 dvr nfs 路径，用以上传及删除 dvr 文件
	path := config.C.Dvr.Path

	//base := filepath.Base(req.File)
	//ext := filepath.Ext(req.File)

	// File 是 docker 中的路径，此路径挂载到了 nfs，我们访问的是 nfs 中的路径
	//cmd := fmt.Sprintf("mv %v/%v %v/%v%v", path, base, path, contentName, ext)
	//logx.Infof("cmd = %v\n", cmd)
	//ssh.Cmd(cmd)

	// dvr 记录入库，维护内容
	var dvr types.Dvr
	dvr.ClientId = req.ClientId
	dvr.Ip = req.Ip
	dvr.Vhost = req.Vhost
	dvr.App = req.App
	dvr.Stream = req.Stream
	dvr.Param = req.Param
	dvr.Cwd = req.Cwd
	dvr.File = req.File
	dvr.Base = filepath.Base(req.File)
	dvr.ContentName = contentName

	dbc := db.GetDBC()
	dbc.Create(&dvr)

	// 上传 dvr 到 minio
	dvrPath := fmt.Sprintf("%v/%v", path, dvr.Base)
	bucketName := time.Now().Format("2006-01-02")
	err := minio.Upload(bucketName, dvr.Base, dvrPath)
	if err != nil {
		panic(err)
	}

	// 删除 nfs dvr
	if config.C.Dvr.Delete == true {
		cmd := fmt.Sprintf("rm %v", dvrPath)
		re := ssh.Cmd(cmd)
		if re != "" {
			panic(re)
		}
	}

	re := types.DvrResponse{
		ErrorCode: 0,
	}
	return &re, nil
}
