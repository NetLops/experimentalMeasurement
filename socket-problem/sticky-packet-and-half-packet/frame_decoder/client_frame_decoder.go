package frame_decoder

import (
	"experimentalMeasurement/socket-problem/protocol"
	"fmt"
	"log"
	"math/rand"
	"net"
)

// ClientTcpFrameDecoder length field based frame decoder
func ClientTcpFrameDecoder(conn net.Conn) {
	log.Println("client, length field based frame decoder")
	for i := 0; i < 10; i++ {
		userName := randStringRunes(6)
		words := "{\"Name\":\"" + userName + "20211217\",\"Meta\":\"golang\",\"Content\":\"message\"}"
		_, err := conn.Write(protocol.Packet([]byte(words)))
		if err != nil {
			fmt.Println(err, ", 写入字符串错误 index=", i)
			return
		}
		fmt.Println(words) // 打印发送出去的信息
	}
	log.Println("send over")

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
