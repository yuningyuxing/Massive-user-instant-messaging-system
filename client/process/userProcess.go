package process

//本文件的作用
//1.处理和用户相关的业务
//2.登陆
//3.注册....
import (
	"bao/client/utils"
	"bao/common/message"
	"encoding/json"
	"fmt"
	"net"
)

// 将方法绑定到结构体中 暂时不需要字段
type UserProcess struct {
}

// 同在main包下 写login函数
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
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
	tf := &utils.Transfer{
		Conn: conn,
	}

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
	//fmt.Println(mes.Data)
	//然后将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return err
	}
	tf.WritePkg(data)

	//现在我们要处理服务器端返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//此时我们需要在客户端启动一个协程，用来保持和服务器通讯，如果服务器有数据推送给客户端 则接受并显示在客户端的终端
		go serverProcessMes(conn)
		//显示我们登陆成功的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

// 写一个注册函数  注册逻辑和登录逻辑和大致相同
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册信息发送错误 err=", err)
		return
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请前往登陆")
	} else {
		fmt.Println(registerResMes.Error)
	}
	return
}
