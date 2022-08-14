package utils

import "fmt"

type RspCode int64

const (
	RspSuccess RspCode = 1000
	RspInvaildParam RspCode = 1001
)

var rspMsg = map[RspCode]string{
	RspSuccess:         "success",
	RspInvaildParam:   "invaild param",
}

func (rsp RspCode)Msg() string {
	msg, ok := rspMsg[rsp]
	if ok {
		return msg
	}

	return fmt.Sprintf("unkown failed, %v", rsp)
}