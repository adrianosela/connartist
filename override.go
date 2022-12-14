package connartist

import (
	"net"
	"time"
)

// WithRead overrides the Read method on the Conn
func (c *Conn) WithRead(fn func(conn net.Conn) func([]byte) (int, error)) *Conn {
	c.read = fn(c.innerConn)
	return c
}

// WithWrite overrides the Write method on the Conn
func (c *Conn) WithWrite(fn func(conn net.Conn) func([]byte) (int, error)) *Conn {
	c.write = fn(c.innerConn)
	return c
}

// WithClose overrides the Close method on the Conn
func (c *Conn) WithClose(fn func(conn net.Conn) func() error) *Conn {
	c.close = fn(c.innerConn)
	return c
}

// WithLocalAddr overrides the LocalAddr method on the Conn
func (c *Conn) WithLocalAddr(fn func(conn net.Conn) func() net.Addr) *Conn {
	c.localAddr = fn(c.innerConn)
	return c
}

// WithRemoteAddr overrides the RemoteAddr method on the Conn
func (c *Conn) WithRemoteAddr(fn func(conn net.Conn) func() net.Addr) *Conn {
	c.remoteAddr = fn(c.innerConn)
	return c
}

// WithSetDeadline overrides the SetDeadline method on the Conn
func (c *Conn) WithSetDeadline(fn func(conn net.Conn) func(t time.Time) error) *Conn {
	c.setDeadline = fn(c.innerConn)
	return c
}

// WithSetReadDeadline overrides the SetReadDeadline method on the Conn
func (c *Conn) WithSetReadDeadline(fn func(conn net.Conn) func(t time.Time) error) *Conn {
	c.setReadDeadline = fn(c.innerConn)
	return c
}

// WithSetWriteDeadline overrides the SetWriteDeadline method on the Conn
func (c *Conn) WithSetWriteDeadline(fn func(conn net.Conn) func(t time.Time) error) *Conn {
	c.setWriteDeadline = fn(c.innerConn)
	return c
}
