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
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	//将长度写入byte切片
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}
	return
}
