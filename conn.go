package connartist

import (
	"net"
	"time"
)

// Conn represents a net.Conn which extends the
// behavior of another net.Conn implementation.
type Conn struct {
	innerConn net.Conn

	read             func([]byte) (n int, err error)
	write            func([]byte) (n int, err error)
	close            func() error
	localAddr        func() net.Addr
	remoteAddr       func() net.Addr
	setDeadline      func(t time.Time) error
	setReadDeadline  func(t time.Time) error
	setWriteDeadline func(t time.Time) error
}

// New returns a new Conn extending the given net.Conn
func New(conn net.Conn) *Conn {
	return &Conn{
		innerConn: conn,

		read:             conn.Read,
		write:            conn.Write,
		close:            conn.Close,
		localAddr:        conn.LocalAddr,
		remoteAddr:       conn.RemoteAddr,
		setDeadline:      conn.SetDeadline,
		setReadDeadline:  conn.SetReadDeadline,
		setWriteDeadline: conn.SetWriteDeadline,
	}
}

// New wraps a Conn within a new Conn to allow for multiple of the same
// function to be chained/layered together. This is particularly useful
// when dealing with messages that include multiple headers.
//
// e.g. c.WithRead(fna).New().WithRead(fnb).New() and so on...
func (c *Conn) New() *Conn {
	return &Conn{
		innerConn: c,

		read:             c.Read,
		write:            c.Write,
		close:            c.Close,
		localAddr:        c.LocalAddr,
		remoteAddr:       c.RemoteAddr,
		setDeadline:      c.SetDeadline,
		setReadDeadline:  c.SetReadDeadline,
		setWriteDeadline: c.SetWriteDeadline,
	}
}
