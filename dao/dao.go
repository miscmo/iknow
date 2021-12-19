package dao

import (
	"fmt"
	"log"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/miscmo/soil/base"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Conf = struct {
	MySQL struct {
		Dsn     string
		Address string

		MaxIdleConnCnt int
		MaxOpenConnCnt int

		TableMusic    string
		TableUserInfo string
	}
}{}

var TableMusic *gorm.DB
var TableUserInfo *gorm.DB

func InitDB() {
	if _, err := toml.DecodeFile(base.CmdFlag.ConfPath, &Conf); err != nil {
		errMsg := fmt.Sprintf("解析配置文件 %s 失败：%s", base.CmdFlag.ConfPath, err.Error())
		log.Fatal(errMsg)
	}

	log.Printf("Conf = %v", Conf)

	dsn := strings.Replace(Conf.MySQL.Dsn, "{addr}", Conf.MySQL.Address, 1)
	log.Printf("dsn = %s", dsn)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		errMsg := fmt.Sprintf("连接数据库 %s 失败：%s", dsn, err.Error())
		log.Fatal(errMsg)
	}

	db.DB().SetMaxOpenConns(Conf.MySQL.MaxOpenConnCnt)
	db.DB().SetMaxIdleConns(Conf.MySQL.MaxIdleConnCnt)

	TableMusic = db.Table(Conf.MySQL.TableMusic)
	TableUserInfo = db.Table(Conf.MySQL.TableUserInfo)

	log.Printf("初始化数据库成功 %v %v", TableMusic, TableUserInfo)
}

type TableMusicModel struct {
	Id          int
	Name        string
	Passwd      string
	Email       string
	Phone       string
	Permissions int
}
