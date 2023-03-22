package model

// 本文件定义一个User结构体
type User struct {
	//注意我们这里可的tag必须和用户信息中json字符串的key对应 这样才能保证序列化和反序列化成功
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
