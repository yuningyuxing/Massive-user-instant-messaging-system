package main

import (
	"bao/common/message"
	"bao/server/process2"
	"bao/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 本函数 根据客户端发送消息种类不同，决定调用那个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	//我们用switch实现业务逻辑
	switch mes.Type {
	case message.LoginMesType: //处理登陆
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType: //处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("此类消息类型不存在，无法处理.....")
	}
	return
}
func (this *Processor) process0() (err error) {
	for {
		//因为我们要将客户端传来的消息转化为结构体 我们这里将他包装成一个函数
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF { //这里的错误是因为客户端断开了链接，所以我们服务器也断开链接
				fmt.Println("客户端断开链接，因此服务端断开链接")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		//将接受的信息 传到一个函数里，函数用来区别信息类型 并做出对应操作
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
	return
}
