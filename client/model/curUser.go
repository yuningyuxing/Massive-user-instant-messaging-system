package model

import (
	"bao/common/message"
	"net"
)

// 这是专门用来发送信息的结构体
type CurUser struct {
	Conn net.Conn
	message.User
}
