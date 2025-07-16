package modbus

import (
	"errors"
	"sync"
)

type HandlerFunc func(w ResponseWriter, r *Request)

type Slave struct {
	SlaveId   byte
	handler   ServerHandler
	functions map[int]HandlerFunc
	mu        sync.RWMutex

	Coils            *coilStorage
	DiscreteInputs   *coilStorage
	HoldingRegisters *registerStorage
	InputRegisters   *registerStorage
}

func (s *Slave) Listen() error {
	return s.handler.Listen()
}

func (s *Slave) Close() error {
	return s.handler.Close()
}

func (s *Slave) HandleFuncCode(code int, f HandlerFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.functions[code] = f
}

func (s *Slave) handleFunc(adu []byte) {
	request, err := s.handler.Decode(adu)
	if err != nil {
		return
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	f, ok := s.functions[int(request.FunctionCode)]
	if !ok {
		_ = s.handler.Send(Response{
			Request:   request,
			ErrorCode: 1,
		})
	}

	responseWriter := newResponseWriterImpl(s.handler, request)
	f(responseWriter, request)
}

type responseWriterImpl struct {
	handler ServerHandler
	request *Request
}

func (r *responseWriterImpl) Success() error {
	switch r.request.FunctionCode {
	case FuncCodeReadHoldingRegisters: // read holding registers
		return errors.New("no response data")
	case FuncCodeWriteSingleRegister, FuncCodeWriteMultipleRegisters:
		return r.handler.Send(Response{
			Request: r.request,
		})
	default:
		return errors.New("unsupported function code")
	}
}

func (r *responseWriterImpl) SuccessWithUint16Slice(data []uint16) error {
	switch r.request.FunctionCode {
	case FuncCodeReadHoldingRegisters: // read holding registers
		return r.handler.Send(Response{
			Request:     r.request,
			ValueLength: byte(len(data) * 2),
			Values:      uint16SliceToBytes(data),
			ErrorCode:   0,
		})
	case FuncCodeWriteSingleRegister, FuncCodeWriteMultipleRegisters:
		return r.handler.Send(Response{
			Request: r.request,
		})
	default:
		return errors.New("unsupported function code")
	}
}

func (r *responseWriterImpl) Error(err error) error {
	return r.handler.Send(Response{
		Request:   r.request,
		ErrorCode: 4,
	})
}

func newResponseWriterImpl(handler ServerHandler, request *Request) *responseWriterImpl {
	return &responseWriterImpl{
		handler: handler,
		request: request,
	}
}

func NewSlave(handler ServerHandler) *Slave {
	slave := &Slave{
		SlaveId:          0,
		handler:          handler,
		functions:        make(map[int]HandlerFunc),
		mu:               sync.RWMutex{},
		Coils:            &coilStorage{},
		DiscreteInputs:   &coilStorage{},
		HoldingRegisters: &registerStorage{},
		InputRegisters:   &registerStorage{},
	}

	handler.HandleFunc(slave.handleFunc)
	return slave
}
