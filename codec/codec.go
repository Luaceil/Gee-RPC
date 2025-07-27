package codec

import "io"

type Header struct {
	ServingMethod string
	Seq           uint64
	Error         string
}

type codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}
