package kvn

import (
	"fmt"
	"net"
)

const NAME = "Kevin"

type Response struct {
	NumHeaders int
}

func (r *Response) Write(conn net.Conn) error {
	// Response Line
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	// Headers
	conn.Write([]byte("Server: KvN/0.1\r\n"))
	conn.Write([]byte("Content-Type: text/html\r\n"))
	// Blank Line
	conn.Write([]byte("\r\n"))
	// Body
	conn.Write([]byte(fmt.Sprintf("THERE WERE %d HEADERS, %s WAS HERE", r.NumHeaders, NAME)))
	conn.Close()
	return nil
}
