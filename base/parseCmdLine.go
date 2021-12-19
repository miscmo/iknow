package base

import "flag"

var ConfPath string

var CmdFlag = struct {
	ConfPath string
}{}

func init() {
	flag.StringVar(&CmdFlag.ConfPath, "conf", "./conf/config.toml", "-conf path")
}
