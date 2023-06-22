package netpoll

import (
	"errors"
	"net"
	"syscall"
	"time"
)

type netFD struct {
	net.Conn
}

func newNetFD(conn net.Conn) *netFD {
	return &netFD{
		Conn: conn,
	}
}

func (c *netFD) Fd() (fd int) {
	return getFd(c.Conn)
}

// SetKeepAlive implements Conn.
// TODO: only tcp conn is ok.
func (c *netFD) SetKeepAlive(second int) error {
	tcpConn, ok := c.Conn.(*net.TCPConn)
	if !ok {
		return nil
	}
	if second > 0 {
		err := tcpConn.SetKeepAlive(true)
		if err != nil {
			return err
		}
		err = tcpConn.SetKeepAlivePeriod(time.Duration(second) * time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}

func getFd(c interface{}) int {
	c2, ok := c.(interface {
		SyscallConn() (syscall.RawConn, error)
	})
	if !ok {
		return 0
	}
	c3, err := c2.SyscallConn()
	if err != nil {
		return 0
	}
	ch := make(chan uintptr, 1)
	err = c3.Control(func(fd uintptr) {
		ch <- fd
	})
	if err != nil {
		return 0
	}
	return int(<-ch)
}

// Various errors contained in OpError.
var (
	errMissingAddress = errors.New("missing address")
)
