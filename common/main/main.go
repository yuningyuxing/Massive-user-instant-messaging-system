package main

//用于测试的小工具
import (
	"bao/common/message"
	"encoding/json"
	"fmt"
)

func main() {
	s1 := message.LoginMes{
		UserId:   10,
		UserPwd:  "123456",
		UserName: "刘桑",
	}
	res, err := json.Marshal(s1)
	if err != nil {
		return
	}
	data := string(res)
	fmt.Println(data)
}

//"{\"userId\":10,\"userPwd\":\"123456\",\"userName\":\"刘桑\"}"
