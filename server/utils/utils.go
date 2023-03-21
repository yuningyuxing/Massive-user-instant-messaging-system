package utils

//这里我们重构server代码 本包主要用来将一些常用的工具的函数，结构体封装
import (
	"bao/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 这里将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8192)
	fmt.Println("读取客户端发送的数据...")
	//这里先读一次数据的长度
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}
	//我们拿到要读的长度后 将他转化为uint32 这样我们就直到下次读取要读取多长的信息了
	var pkgLen uint32
	//这里调函数将byte切片转化为uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	//这里检查一下接受到的信息和要接受的信息长度是否一致
	if n != int(pkgLen) || err != nil {
		return
	}
	//将接受的信息反序列化 注意这里我们要取地址
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}
	return
}
