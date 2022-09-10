package main

import (
	"experimentalMeasurement/socket-problem/frame_decoder"
	"fmt"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":7373")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(lis net.Listener) {
		err3 := lis.Close()
		if err3 != nil {

		}
	}(lis)

	log.Println("Waiting for client...") // 启动后，等待客户端访问

	for {
		conn, err2 := lis.Accept() // 监听客户端
		if err2 != nil {
			fmt.Println("conn error:", err2)
			return
		}
		log.Println(conn.RemoteAddr().String(), "tcp connection success")

		// fix length
		//fix_length.ServerTcpFixLength(conn)

		// delimiter
		//go delimiter_based.ServerTcpDelimiter(conn)

		// frame_decoder
		go frame_decoder.ServerTcpFrameDecoder(conn)
	}
}
