package process2

import (
	"bao/common/message"
	"bao/server/model"
	"bao/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

// 给结构体绑定一个serverProcessLogin 专门处理登陆请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200
		//此时因为登陆成功 我们要将用户放到userMgr(用户在线列表)中
		this.UserId = loginMes.UserId
		userMgr.AddOnlienUser(this)
		//通知其他用户我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前在线用户的id 放入到loginResMes.UsersId 遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserId = append(loginResMes.UserId, id)
		}
		fmt.Println(user.UserName, "登陆成功")
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
	//因为我们采用的是分层模式 所以要先创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 专门处理注册请求
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	//注意这里为什么要强转呢？  因为我们从客户端发送过来的信息.User虽然样子和我们server的User一样 但本质上他们不是同种结构体需要我们强转类型
	err = model.MyUserDao.Register((*model.User)(&registerMes.User))
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."

		}
	} else {
		registerResMes.Code = 200
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 通知所有在线用户 我上线了
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//注意这个函数的一些细节 首先我们为了实现能遍历到所有当前在线的客户端 我们遍历得到的是UserProcess 是一个客户端 包含他的链接
	//所以此时我们调用up.Not 的时候 就用的时要发送到客户端的链接 并且我们传的参数时自己 因为要给别人说为上线了
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

// 这个函数的逻辑比较简单 就是把自己上线的消息 封装在结构体里 也就是经典的mes 然后序列化传过去即可
func (this *UserProcess) NotifyMeOnline(UserId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = UserId
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
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
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}
