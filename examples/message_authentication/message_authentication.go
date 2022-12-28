package message_authentication

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	hmacKey = "somesupersecretstring"
	hmacLen = 32
)

// Read reads a message and verifies its HMAC
func Read(c net.Conn) func([]byte) (int, error) {
	return func(b []byte) (int, error) {
		// buffer big enough to read hmac and fill b
		buf := make([]byte, hmacLen+len(b))

		// read at least one byte more than the hmac length
		n, err := io.ReadAtLeast(c, buf, hmacLen+1)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return n, io.EOF
			}
			if errors.Is(err, io.ErrUnexpectedEOF) {
				return n, fmt.Errorf("bad message received, too short to have HMAC")
			}
			return n, fmt.Errorf("failed to read message: %s", err)
		}

		// split data into hmac and message
		mac, msg := buf[:hmacLen], buf[hmacLen:n]

		// compute hmac for message
		computed := hmac.New(sha256.New, []byte(hmacKey))
		if n, err = computed.Write(msg); err != nil {
			// note: hash.Write() never returns an error as per godoc
			// (https://pkg.go.dev/hash#Hash) but we check it regardless
			return n, err
		}
		sum := computed.Sum(nil)

		// compare received vs computed HMAC
		if string(mac) != string(computed.Sum(nil)) {
			return 0, fmt.Errorf(
				"mac did not match sum: mac(%s)|sum(%s)",
				base64.StdEncoding.EncodeToString(mac),
				base64.StdEncoding.EncodeToString(sum),
			)
		}

		// copy the message onto the given buffer
		return copy(b, msg), nil
	}
}

// Write computes a message's HMAC and writes it along with the message
func Write(c net.Conn) func([]byte) (int, error) {
	return func(b []byte) (int, error) {
		// compute HMAC for message
		computed := hmac.New(sha256.New, []byte(hmacKey))
		if n, err := computed.Write(b); err != nil {
			// note: hash.Write() never returns an error as per godoc
			// (https://pkg.go.dev/hash#Hash) but we check it regardless
			return n, err
		}
		sum := computed.Sum(nil)

		// put together data (${HMAC}${MSG})
		data := append(sum, b...)

		// write data to conn
		n, err := c.Write(data)
		if err != nil {
			return n, fmt.Errorf("failed to write signed message: %s", err)
		}
		return n - hmacLen, nil
	}
}
