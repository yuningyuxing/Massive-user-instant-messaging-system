package process

import (
	"bao/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message, kind int) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}
	if kind == 1 {
		info := fmt.Sprintf("用户id:\t%d 对大家说:%s", smsMes.UserId, smsMes.Content)
		fmt.Println(info)
	} else if kind == 2 {
		info := fmt.Sprintf("用户id:\t%d 对你说:%s", smsMes.UserId, smsMes.Content)
		fmt.Println(info)
	}
}
