package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/adrianosela/connartist"
	"github.com/adrianosela/connartist/examples/message_authentication"
)

var (
	protocol string
	address  string
)

func main() {
	flag.StringVar(&protocol, "protocol", "tcp", "listener protocol to use")
	flag.StringVar(&address, "address", "localhost:1234", "listener address (i.e. HOST:PORT) to use")
	rand.Seed(time.Now().Unix())

	l, err := net.Listen(protocol, address)
	if err != nil {
		log.Fatalf("could not start %s listener on %s", protocol, address)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("could not accept new connection %s listener on %s", protocol, address)
		}

		conn := connartist.New(c).
			WithRead(message_authentication.Read).
			WithWrite(message_authentication.Write)
		defer conn.Close()

		go handleConn(rand.Intn(1000), conn)
	}
}

func handleConn(clientID int, conn net.Conn) {
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
		}

		fmt.Printf("[%d] %s", clientID, data)

		writer := bufio.NewWriter(conn)
		_, err = writer.WriteString(fmt.Sprintf("[%s] %s", time.Now().Format(time.RFC3339), data))
		if err != nil {
			log.Printf("failed to write to writer for client id %d: %s", clientID, err)
			return
		}
		if err = writer.Flush(); err != nil {
			log.Printf("failed to flush writer for client id %d: %s", clientID, err)
			return
		}
	}
}
