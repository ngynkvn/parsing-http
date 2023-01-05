package kvn

import (
	"bufio"
	"errors"
	"net"
	"strings"
)

type RequestLine struct {
	Method   string
	Path     string
	Protocol string
}

type HeaderFields struct {
	Headers map[string]string
}

type Request struct {
	RequestLine
	HeaderFields
}

func NewRequest(r RequestLine, hf HeaderFields) Request {
	return Request{
		RequestLine:  r,
		HeaderFields: hf,
	}
}

type WritableResponse interface {
	Write(net.Conn) error
}

func (hf *HeaderFields) Add(k, v string) {
	hf.Headers[k] = v
}

const SPACE = " "

func GetRequestLine(sc *bufio.Scanner) (RequestLine, error) {
	sc.Scan()
	text := strings.Split(sc.Text(), SPACE)
	if len(text) != 3 {
		return RequestLine{}, errors.New("unexpected request line format")
	}
	return RequestLine{
		Method:   text[0],
		Path:     text[1],
		Protocol: text[2],
	}, nil
}

func GetHeaders(sc *bufio.Scanner) (HeaderFields, error) {
	hf := HeaderFields{
		Headers: make(map[string]string),
	}
	for sc.Scan() {
		rawHeader := sc.Text()
		if len(rawHeader) == 0 {
			return hf, nil
		}
		before, after, found := strings.Cut(rawHeader, ": ")
		if !found {
			return hf, errors.New("could not find header seperator")
		}
		hf.Add(before, after)
	}
	return hf, nil
}
