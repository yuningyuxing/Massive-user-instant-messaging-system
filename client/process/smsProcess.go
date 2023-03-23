package process

import (
	"bao/client/model"
	"bao/client/utils"
	"bao/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

var CurUser model.CurUser

// 将要发送的消息发送给服务器
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType
	//创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	data, err := json.Marshal(smsMes)
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
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}
	return
}
func (this *SmsProcess) SendGroupOnlineMes(content string, userId int) (err error) {
	var mes message.Message
	mes.Type = message.SmsOnlineMesType
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.AimUserId = userId
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	data, err := json.Marshal(smsMes)
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
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}
	return

	return
}
