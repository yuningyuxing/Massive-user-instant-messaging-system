package process

import (
	"bao/client/utils"
	"bao/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//本文件作用
//1.显示登陆成功界面
//2.保持和服务器通讯
//3.当读取服务器发送的消息后，就会显示在界面

// 显示登陆成功后的界面
func ShowMenu() {
	fmt.Println("-------恭喜xxx登陆成功-------")
	fmt.Println("-------1. 显示在线用户列表-------")
	fmt.Println("-------2. 发送消息-------")
	fmt.Println("-------3. 信息列表-------")
	fmt.Println("-------4. 退出系统-------")
	fmt.Println("请选择(1-4):")
	var key int
	fmt.Scanln(&key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("你选择退出系统...")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确")
	}
}

// 定义一个和服务器保持通讯的函数
func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//读取到信息后，进行下一步的逻辑处理
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
			fmt.Println(notifyUserStatusMes.UserId, "上线了")
		default:
			fmt.Println("服务器端返回了未知的消息类型")
		}
	}

}
