package modbus

type HandlerFunc func(writer ResponseWriter, request *Request)

type Handler interface {
	Serve() error
}
