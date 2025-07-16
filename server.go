package modbus

type ServerHandler interface {
	Listen() error
	Close() error
	HandleFunc(handler func(adu []byte))
	Decode(adu []byte) (request *Request, err error)
	Send(response Response) (err error)
}
