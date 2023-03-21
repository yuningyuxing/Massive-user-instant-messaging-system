package main

//本文件主要用来构建login函数
import (
	"bao/common/message"
	"encoding/json"
	"fmt"
	"net"
)

// 同在main包下 写login函数
func login(userId int, userPwd string) (err error) {
	//下面要开始定协议
	//1.链接到服务器
	//客户端申请向目标建立链接
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return err
	}
	//延时关闭链接
	defer conn.Close()

	//申请一个用来描述消息信息的结构体
	var mes message.Message
	//消息类型
	mes.Type = message.LoginMesType

	//用来存储用户的具体信息
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//将得到的用户信息序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return err
	}
	//注意因为我们用来描述消息的结构体.Data是string类型 而序列化后是type切片 所以我们转化一下
	mes.Data = string(data)
	//然后将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return err
	}
	writePkg(conn, data)

	//现在我们要处理服务器端返回的消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
