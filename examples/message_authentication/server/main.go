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

		// FIXME: why does bufio.NewWriter(conn) not work?
		// FIXME: catch and handle (n int, err error)
		conn.Write([]byte(fmt.Sprintf("[%s] %s\n", time.Now().Format(time.RFC3339), data)))
	}
}
