package delimiter_based

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func ServerTcpDelimiter(conn net.Conn) {
	fmt.Println("Server delimiter based")
	reader := bufio.NewReader(conn)
	for {
		slice, err := reader.ReadSlice('\n')
		log.Println(slice)
		log.Println(err)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			fmt.Printf("slice %s", slice)
			break
		}
		if err != nil {
			log.Println("delimiter based: ", err)
			continue
		}
		fmt.Printf("slice %s", slice)
	}
}
