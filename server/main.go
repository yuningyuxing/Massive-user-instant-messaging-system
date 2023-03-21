package main

import (
	"bao/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8192)
	fmt.Println("读取客户端发送的数据...")
	//这里先读一次数据的长度
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}
	//我们拿到要读的长度后 将他转化为uint32 这样我们就直到下次读取要读取多长的信息了
	var pkgLen uint32
	//这里调函数将byte切片转化为uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	n, err := conn.Read(buf[:pkgLen])
	//这里检查一下接受到的信息和要接受的信息长度是否一致
	if n != int(pkgLen) || err != nil {
		return
	}
	//将接受的信息反序列化 注意这里我们要取地址
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	//将长度写入byte切片
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}
	return
}

// 这个函数用来处理登陆
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	var loginMes message.LoginMes
	//从mes取出mes.Data 并反序列化成LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//现在我们要申明一个用来返回登陆信息的结构体
	var resMes message.Message
	resMes.Type = message.LoginReMesType
	//这里存的是返回信息的信息本身
	var loginResMes message.LoginResMes
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//我们规定状态码200 表示合法登陆
		loginResMes.Code = 200
	} else {
		//状态码500 表示不合法登陆
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册后再使用"
	}
	//先将loginResMes序列化成切片后再转成string从而作为数据本身赋给resMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) fail=", err)
		return
	}
	resMes.Data = string(data)
	//再将resMes序列化后发送回去
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail=", err)
		return
	}
	//发送data回客户端
	err = writePkg(conn, data)
	return
}

// 本函数 根据客户端发送消息种类不同，决定调用那个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	//我们用switch实现业务逻辑
	switch mes.Type {
	case message.LoginMesType: //处理登陆
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType: //处理注册
	default:
		fmt.Println("此类消息类型不存在，无法处理.....")
	}
	return
}
func process(conn net.Conn) {
	//注意延时关闭conn
	defer conn.Close()
	for {
		//因为我们要将客户端传来的消息转化为结构体 我们这里将他包装成一个函数
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF { //这里的错误是因为客户端断开了链接，所以我们服务器也断开链接
				fmt.Println("客户端断开链接，因此服务端断开链接")
				return
			} else {
				fmt.Println("readPkg err=", err)
				return
			}
		}
		//将接受的信息 传到一个函数里，函数用来区别信息类型 并做出对应操作
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}

}

func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听...")
	//调用net.Listen来监听指定服务器的端口
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//循环等待客户端链接
	for {
		fmt.Println("等待客户端来链接服务器....")
		//在这里一直等待 没有就阻塞
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			return
		}
		//为每一个客户端 开一个协程去服务
		go process(conn)
	}
}
