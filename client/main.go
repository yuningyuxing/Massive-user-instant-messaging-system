package main

import (
	"fmt"
	"os"
)

func main() {

	var key int
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
			loop = false
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
	//根据用户的输入，显示新的提示信息
	if key == 1 {

	}
}
