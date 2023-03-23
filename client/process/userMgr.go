package process

import (
	"bao/client/model"
	"bao/common/message"
	"fmt"
)

var onlineUsers map[int]*model.User = make(map[int]*model.User, 10)

// 在客户端显示当前在线的用户
func outputOnlineUser() {
	fmt.Println("显示当前在线用户列表")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

// 有人来告诉我 他上线了  我就检查一下 自己维护的在线表里面有没有他 如果没有就添加 否则就不管
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &model.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
}
