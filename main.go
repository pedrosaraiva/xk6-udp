package udp

import (
	"net"

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

func (udp *UDP) Write(conn net.Conn, data []byte) error {
	_, err := conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (udp *UDP) Read(conn net.Conn, size int) ([]byte, error) {
	buf := make([]byte, size)
	_, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (udp *UDP) WriteLn(conn net.Conn, data []byte) error {
	return udp.Write(conn, append(data, []byte("\n")...))
}

func (tcp *UDP) Close(conn net.Conn) error {
	err := conn.Close()
	if err != nil {
		return err
	}
	return nil
}
