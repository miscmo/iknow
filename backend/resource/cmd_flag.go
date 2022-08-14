package resource

import (
	"flag"
	"sync"
)

var ConfPath string

type CmdFlag struct {
	ConfPath string
}

var cmdflag CmdFlag

var cmdFlagOnce sync.Once

func GetCmdFlag() *CmdFlag {
	cmdFlagOnce.Do(func() {
		flag.StringVar(&cmdflag.ConfPath, "cpath", "./conf/config.toml", "-cpath <path>")
	})

	return &cmdflag
}
