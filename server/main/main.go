package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//注意延时关闭conn
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process0()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误 err=", err)
		return
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
