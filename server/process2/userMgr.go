package process2

//本文件用于维护用户在线列表  可以实现对onlineUsers的增删改查
import (
	"fmt"
)

// 因为全局只有一个 所以我们直接定义为全局变量
var (
	userMgr *UserMgr
)

// 本结构体就是用来维护用户在线列表的
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 初始化UserMgr
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUsers的添加
func (this *UserMgr) AddOnlienUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// 删除某个用户
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// 返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// 根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok { //此时说明要查找的这个用户不存在
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
