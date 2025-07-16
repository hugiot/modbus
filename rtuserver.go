package modbus

import (
	"encoding/binary"
	"github.com/goburrow/modbus"
	"io"
)

type RTUServerHandler struct {
	serialPort
	core    *modbus.RTUClientHandler
	handler func(adu []byte)
}

func (h *RTUServerHandler) Listen() error {
	// open serial port
	err := h.serialPort.Connect()
	if err != nil {
		return err
	}
	defer h.serialPort.Close()

	var buffer [256]byte
	for {
		n, err := h.serialPort.Read(buffer[:])
		if err != nil {
			if err == io.EOF {
				return io.EOF
			}
		}

		if n == 0 {
			continue
		}

		data := make([]byte, n)
		copy(data, buffer[:n])
		h.handler(data)
	}
}

func (h *RTUServerHandler) Close() error {
	return h.serialPort.Close()
}

func (h *RTUServerHandler) HandleFunc(handler func(adu []byte)) {
	h.handler = handler
}

func (h *RTUServerHandler) Decode(adu []byte) (request *Request, err error) {
	return ReadRTURequest(adu)
}

func (h *RTUServerHandler) Send(response Response) (err error) {
	adu, err := h.encode(response)
	if err != nil {
		return err
	}

	_, err = h.serialPort.Write(adu)
	return err
}

func (h *RTUServerHandler) encode(response Response) (adu []byte, err error) {
	h.core.SlaveId = response.Request.SlaveId

	if response.ErrorCode != 0 {
		return h.core.Encode(&modbus.ProtocolDataUnit{
			FunctionCode: response.Request.FunctionCode | 0x80,
			Data:         []byte{response.ErrorCode},
		})
	}

	var data []byte

	switch int(response.Request.FunctionCode) {
	case FuncCodeReadHoldingRegisters: // read holding registers
		data = append(data, response.ValueLength)
		data = append(data, response.Values...)
	case FuncCodeWriteSingleRegister: // write single register
		data = binary.BigEndian.AppendUint16(data, response.Request.Address)
		data = append(data, response.Request.Values...)
	case FuncCodeWriteMultipleRegisters: // write multiple registers
		data = binary.BigEndian.AppendUint16(data, response.Request.Address)
		data = binary.BigEndian.AppendUint16(data, response.Request.Quantity)
	}

	return h.core.Encode(&modbus.ProtocolDataUnit{
		FunctionCode: response.Request.FunctionCode,
		Data:         data,
	})
}

func NewRTUServerHandler(address string) *RTUServerHandler {
	return &RTUServerHandler{
		serialPort: createSerialPort(address),
		core:       modbus.NewRTUClientHandler(address),
	}
}
