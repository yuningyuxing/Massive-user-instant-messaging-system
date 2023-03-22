package main

import (
	"bao/server/model"
	"fmt"
	"net"
	"time"
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

// 本函数用来完成对UserDao的初始化任务
func initUserDao() {
	//注意我们一定先初始化连接池才能初始化UserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	//服务器启动时，我们去初始化redis链接池
	initPool("127.0.0.1:6379", 16, 0, 300*time.Second)
	initUserDao()
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
