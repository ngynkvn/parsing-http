package main

import (
	"bufio"
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
