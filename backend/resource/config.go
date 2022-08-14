package resource

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/miscmo/soil/base"
	"log"
	"sync"
)

type Conf struct {
	Log struct {
	}

	MySQL struct {
		Dsn     string
		Address string

		MaxIdleConnCnt int
		MaxOpenConnCnt int

		TableMusic    string
		TableUserInfo string
	}

	MongoDB struct{
		URI string
		IKnowDB string
		IKnowNodeCollection string
		IKnowNoteCollection string
	}
}

var configOnce sync.Once


var config Conf

func GetConfig() *Conf {
	configOnce.Do(func() {
		cmdflag := GetCmdFlag()
		if _, err := toml.DecodeFile(cmdflag.ConfPath, &config); err != nil {
			errMsg := fmt.Sprintf("decode config failed, err ：%s, confPath: %s, conf: %+v", err.Error(), base.CmdFlag.ConfPath, config)
			panic(errMsg)
		}

		log.Printf("decode config succ, config = %v", config)
	})

	return &config
}