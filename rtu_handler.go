package modbus

import (
	"github.com/goburrow/modbus"
	"github.com/goburrow/serial"
	"io"
	"sync"
	"time"
)

//type Handler interface {
//	ReadCoils(address, quantity uint16) (results []bool, err error)
//	ReadDiscreteInputs(address, quantity uint16) (results []bool, err error)
//	WriteSingleCoil(address uint16, value bool) error
//	WriteMultipleCoils(address, quantity uint16, value []bool) error
//	ReadInputRegisters(address, quantity uint16) (results []uint16, err error)
//	ReadHoldingRegisters(address, quantity uint16) (results []uint16, err error)
//	WriteSingleRegister(address, value uint16) error
//	WriteMultipleRegisters(address, quantity uint16, value []uint16) error
//}

type RTUHandler struct {
	SlaveID     uint8
	Address     string
	BaudRate    int
	DataBits    int
	StopBits    int
	Parity      string
	Timeout     time.Duration
	MinInterval time.Duration

	mu        sync.RWMutex
	functions map[int]HandlerFunc
	handler   *modbus.RTUClientHandler
	port      serial.Port
}

func (h *RTUHandler) HandleFunc(functionCode int, f HandlerFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.functions[functionCode] = f
}

func (h *RTUHandler) Serve() error {
	if err := h.connect(); err != nil {
		return err
	}
	return nil
}

func (h *RTUHandler) connect() error {
	if h.port == nil {
		port, err := serial.Open(&serial.Config{
			Address:  h.Address,
			BaudRate: h.BaudRate,
			DataBits: h.DataBits,
			StopBits: h.StopBits,
			Parity:   h.Parity,
			Timeout:  h.Timeout,
		})
		if err != nil {
			return err
		}

		h.port = port
	}
	return nil
}

func (h *RTUHandler) getHandlerFunc(code int) (HandlerFunc, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	f, ok := h.functions[code]
	return f, ok
}

func (h *RTUHandler) acceptRequest() {
	timer := time.NewTimer(h.MinInterval)
	defer timer.Stop()

	var buffer [256]byte
	for {
		if h.port == nil {
			break
		}

		// read data
		n, err := io.ReadAtLeast(h.port, buffer[:], 1)
		if err != nil {
			buffer = [256]byte{}
			timer.Reset(h.MinInterval)
			continue
		}
	}
}
