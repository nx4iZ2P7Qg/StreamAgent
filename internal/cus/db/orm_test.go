package db

import (
	"agent/internal/cus/config"
	"agent/internal/types"
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

func Test(t *testing.T) {
	db := GetDBC()
	schema := config.C.Postgres.Schema

	// schema 需要手动创建
	// create schema agent;

	// 建表
	db.Exec(fmt.Sprintf("drop table %s.dvrs", schema))

	// 指定 schema
	gorm.DefaultTableNameHandler = func(db *gorm.DB, table string) string {
		return schema + "." + table
	}

	db.AutoMigrate(&types.Dvr{})

	// 添加数据
	db.Create(&types.Dvr{
		ClientId: 683,
		Ip: "100.100.100.100",
		Vhost: "srs_vhost_anime",
		App: "live",
		Stream: "secret",
		Param: "?token=xxx&salt=yyy",
		Cwd: "/usr/local/srs",
		File: "./objs/nginx/html/live/livestream.1420254068776.flv",
		Base: "livestream.1420254068776.flv",
		ContentName: "2021年7月电影推荐",
	})

	var dvr types.Dvr
	// 主键查询
	if err := db.First(&dvr, 1).Error; err != nil {
		panic("error first")
	}
	//// 关联查询
	//if err := db.Model(&dvr).Related(&dvr.SessionFilter).Related(&dvr.TimeStamp).Error; err != nil {
	//	panic("error model related")
	//}
	//log.Info(dvr)
	// 条件查询
	var dvr2 types.Dvr
	if err := db.Where("content_name like ?", "%电影%").First(&dvr2).Error; err != nil {
		panic("error where")
	}
	// 更新
	if err := db.Model(&dvr2).Update("Ip", "100.100.100.200").Error; err != nil {
		panic("error model update")
	}
	// 删除
	if err := db.Delete(&dvr2).Error; err != nil {
		panic("error delete")
	}
}