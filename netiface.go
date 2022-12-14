package connartist

import (
	"net"
	"time"
)

// statement ensures Conn implements net.Conn at compile-time
var _ net.Conn = (*Conn)(nil)

// Read runs the Read method on the conn
func (c *Conn) Read(b []byte) (n int, err error) {
	return c.read(b)
}

// Write runs the Write method on the conn
func (c *Conn) Write(b []byte) (n int, err error) {
	return c.write(b)
}

// Close runs the Close method on the conn
func (c *Conn) Close() error {
	return c.close()
}

// LocalAddr runs the LocalAddr method on the conn
func (c *Conn) LocalAddr() net.Addr {
	return c.localAddr()
}

// RemoteAddr runs the RemoteAddr method on the conn
func (c *Conn) RemoteAddr() net.Addr {
	return c.remoteAddr()
}

// SetDeadline runs the SetDeadline method on the conn
func (c *Conn) SetDeadline(t time.Time) error {
	return c.setDeadline(t)
}

// SetReadDeadline runs the SetReadDeadline method on the conn
func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.setReadDeadline(t)
}

// SetWriteDeadline runs the SetWriteDeadline method on the conn
func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.setWriteDeadline(t)
}
