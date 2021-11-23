package config

import (
	"github.com/spf13/viper"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
)

type Config struct {
	Postgres Postgres
	Ssh      Ssh
	Dvr      Dvr
	Minio    Minio
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	Db       string
	Schema   string
}

type Ssh struct {
	Host     string
	Port     int
	User     string
	Password string
}

type Dvr struct {
	Path   string
	Delete bool
}

type Minio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
}

var C Config

func init() {
	viper.SetConfigName("local")
	// 本地开发需要的路径
	viper.AddConfigPath("internal/cus/config")
	// orm_test.go 需要的路径
	viper.AddConfigPath("../config")
	// docker 运行时需要的路径
	viper.AddConfigPath("/app/etc")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}

	var c logx.LogConf
	// 本地开发需要的路径
	conf.MustLoad("etc/log.yaml", &c)
	// client_test.go 需要的路径
	//conf.MustLoad("../../../etc/log.yaml", &c)
	logx.MustSetup(c)
}
