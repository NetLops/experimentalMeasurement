package main

import (
	"experimentalMeasurement/socket-problem/frame_decoder"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}
func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:7373")
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", "43.154.236.47:7373")
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", "82.157.193.4:7373")
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", "165.22.56.51:7373")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")

	// fix-length
	//fix_length.ClientTcpFixLength(conn)

	// delimiter based
	//delimiter_based.ClientTcpDelimiter(conn)

	// frame_decoder
	frame_decoder.ClientTcpFrameDecoder(conn)
	defer conn.Close()
}
