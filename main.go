package main

import (
	"bufio"
	"errors"
	"fmt"
	. "http/kvn"
	"net"

	"github.com/davecgh/go-spew/spew"
)

func Must[T any](ret T, e error) T {
	if e != nil {
		panic(e)
	}
	return ret
}

func businessLogic(r Request) (Response, error) {
	return Response{
		NumHeaders: len(r.Headers),
	}, nil
}

func handleConnection(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if ok {
				WriteErrorResponse(conn, err)
			} else {
				WriteErrorResponse(conn, errors.New("an unknown error was encountered"))
			}
		}
	}()
	reader := bufio.NewScanner(bufio.NewReader(conn))
	rql := Must(GetRequestLine(reader))
	headers := Must(GetHeaders(reader))
	reply := Must(businessLogic(NewRequest(rql, headers)))
	spew.Println(rql, headers)
	reply.Write(conn)
	spew.Println("Connection has been handled!")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			panic("Dude we hit the recover lol")
		}
	}()

	ln := Must(net.Listen("tcp", ":8080"))
	spew.Println("Listener created!")
	for {
		conn := Must(ln.Accept())
		spew.Println("New connection accepted!")
		go handleConnection(conn)
	}

}

func WriteErrorResponse(conn net.Conn, err error) {
	conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n"))
	// Headers
	conn.Write([]byte("Server: KvN/0.1\r\n"))
	conn.Write([]byte("Content-Type: text/html\r\n"))
	// Blank Line
	conn.Write([]byte("\r\n"))
	// Body
	conn.Write([]byte(fmt.Sprintf("THERE WAS AN ERROR ENCOUNTERED: %s", err)))
	conn.Close()
}
