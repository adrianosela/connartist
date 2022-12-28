package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

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

	c, err := net.Dial(protocol, address)
	if err != nil {
		log.Fatalf("could not dial %s to %s", protocol, address)
	}

	conn := connartist.New(c).
		WithRead(message_authentication.Read).
		WithWrite(message_authentication.Write)
	defer conn.Close()

	for {
		fmt.Print(">> ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read from stdin: %s", input)
		}

		writer := bufio.NewWriter(conn)
		if _, err = writer.WriteString(input); err != nil {
			log.Fatalf("failed to write to writer: %s", err)
		}
		if err = writer.Flush(); err != nil {
			log.Fatalf("failed to flush writer: %s", err)
		}

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
		}

		fmt.Print(msg)
	}
}
