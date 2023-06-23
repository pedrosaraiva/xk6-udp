package udp

import (
	"io"
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
	if err := conn.SetWriteDeadline(time.Now().Add(writeTimeout)); err != nil {
		return err
	}

	_, err := conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (udp *UDP) Read(conn net.Conn, size int, readTimeout time.Duration) ([]byte, error) {
	if err := conn.SetReadDeadline(time.Now().Add(readTimeout)); err != nil {
		return nil, err
	}

	buf := make([]byte, 0, size)
	tmp := make([]byte, size)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err == net.ErrClosed || err == io.EOF {
				break // Connection closed, exit the loop
			}
			return nil, err
		}
		buf = append(buf, tmp[:n]...)
		if len(buf) >= size {
			break // Read enough data, exit the loop
		}
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
