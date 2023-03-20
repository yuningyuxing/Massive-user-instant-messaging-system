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
		//打印接受的信息 是一个结构体
		fmt.Println("mes=", mes)
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
