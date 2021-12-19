package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/miscmo/soil/base"
	"github.com/miscmo/soil/blog"
	"github.com/miscmo/soil/dao"
	"github.com/miscmo/soil/music"
)

type Conf struct {
	Log struct {
	}

	MySQL struct {
		Dns           string
		Address       string
		TableUserInfo string
	}
}

func init() {
	dao.InitDB()
}

var conf Conf

func MainPage(c *gin.Context) {
	c.HTML(200, "index.html", "")
}

func main() {
	if _, err := toml.DecodeFile(base.CmdFlag.ConfPath, &conf); err != nil {
		errMsg := fmt.Sprintf("读取配置文件失败：%s", err.Error())
		log.Fatal(errMsg)
	}

	log.Printf("config: %+v", conf)

	queryDB := dao.TableUserInfo

	var userInfoRecord []dao.TableMusicModel

	if err := queryDB.Where("1=1").Find(&userInfoRecord).Error; err != nil {
		errMsg := fmt.Sprintf("执行SQL失败：%s", err.Error())
		log.Fatal(errMsg)
	}

	log.Printf("%+v", userInfoRecord)

	r := gin.Default()
	r.LoadHTMLGlob("ui/*.html")
	r.Static("/static", "./ui/static")
	r.GET("/", MainPage)
	r.GET("/MusicHome", music.MusicHome)
	r.GET("/MusicList", music.MusicList)
	r.POST("/AddMusic", music.AddMusic)
	r.GET("/PlayMusic/:id", music.PlayMusic)
	r.GET("/BlogHome", blog.BlogHome)
	r.GET("/BlogList", blog.BlogList)
	r.Run(":8080")
}
