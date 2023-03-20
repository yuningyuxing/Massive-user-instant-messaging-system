package main

import "fmt"

// 同在main包下 写login函数
func login(userId int, userPwd string) (err error) {
	//下面要开始定协议
	fmt.Printf(" userId = %d userPwd=%s\n", userId, userPwd)
	return nil
}
