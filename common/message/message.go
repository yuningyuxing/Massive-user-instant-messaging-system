package message

const (
	LoginMesType   = "LoginMes"
	LoginReMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的类型
}
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}
type LoginResMes struct {
	Code  int    `json:"code"'` //返回状态码 500表示用户未错误，200表示登陆成功
	Error string `json:"error"` //返回错误信息
}
