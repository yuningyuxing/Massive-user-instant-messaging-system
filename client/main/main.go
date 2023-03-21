package main

import (
	"bao/client/process"
	"fmt"
	"os"
)

// 定义两个用户 一个表示用户id,一个表示用户密码
var userId int
var userPwd string

func main() {

	//接受用户选择
	var key int
	//判断是否还继续显示菜单
	var loop = true
	for loop {
		fmt.Println("----------------欢迎登陆多人聊天系统----------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			//说明用户要登陆
			fmt.Println("请输入用户的id：")
			fmt.Scanln(&userId) //这里注意输入细节 如果是用Scanf要+\n用来吞换行
			fmt.Println("请输入用户密码：")
			fmt.Scanln(&userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
