package main

//本文件旨在写出收取链接信息并转化成结构体和将信息发送出去的两个函数
import (
	"bao/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8192)
	fmt.Println("读取客户端发送的数据...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// 将信息发送出去
func writePkg(conn net.Conn, data []byte) (err error) {
	//因为我们这里要发送信息长度，从而确保信息不会丢失
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	//我们为了将数字转换成切片 这里用PutUint32 但注意我们接受参数要uint32所以转化一下
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//此时我们就获得了一个描述信息长度的byte切片buf 我们将他传过去 这里传4个字节 因为uint32用了四个字节
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return err
	}
	//上面传输的是 我们要发送的信息的长度 本次是发送信息本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}
	return
}
