package udp

import (
	"net"
	"time"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/udp", new(UDP))
}

type UDP struct{}

func (udp *UDP) Connect(addr string) (net.Conn, error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (udp *UDP) Write(conn net.Conn, data []byte, writeTimeout time.Duration) error {
	if err := conn.SetWriteDeadline(time.Now().Add(writeTimeout * time.Second)); err != nil {
		return err
	}

	_, err := conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (udp *UDP) Read(conn net.Conn, size int, readTimeout time.Duration) ([]byte, error) {
	if err := conn.SetReadDeadline(time.Now().Add(60 * time.Second)); err != nil {
		return nil, err
	}

	buf := make([]byte, size)
	_, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (udp *UDP) WriteLn(conn net.Conn, data []byte, writeTimeout time.Duration) error {
	return udp.Write(conn, append(data, '\n'), writeTimeout)
}

func (udp *UDP) Close(conn net.Conn) error {
	err := conn.Close()
	if err != nil {
		return err
	}
	return nil
}
