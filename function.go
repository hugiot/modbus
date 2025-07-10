package modbus

const (
	FuncCodeReadCoils                  = 1
	FuncCodeReadDiscreteInputs         = 2
	FuncCodeReadHoldingRegisters       = 3
	FuncCodeReadInputRegisters         = 4
	FuncCodeWriteSingleCoil            = 5
	FuncCodeWriteSingleRegister        = 6
	FuncCodeWriteMultipleCoils         = 15
	FuncCodeWriteMultipleRegisters     = 16
	FuncCodeReadWriteMultipleRegisters = 23
	FuncCodeMaskWriteRegister          = 22
	FuncCodeReadFIFOQueue              = 24
)

// FuncCodeText returns the text representation of a Modbus function code.
func FuncCodeText(code int) string {
	switch code {
	case FuncCodeReadCoils:
		return "Read Coils"
	case FuncCodeReadDiscreteInputs:
		return "Read Discrete Inputs"
	case FuncCodeReadHoldingRegisters:
		return "Read Holding Registers"
	case FuncCodeReadInputRegisters:
		return "Read Input Registers"
	case FuncCodeWriteSingleCoil:
		return "Write Single Coil"
	case FuncCodeWriteSingleRegister:
		return "Write Single Register"
	case FuncCodeWriteMultipleCoils:
		return "Write Multiple Coils"
	case FuncCodeWriteMultipleRegisters:
		return "Write Multiple Registers"
	case FuncCodeReadWriteMultipleRegisters:
		return "Read/Write Multiple Registers"
	case FuncCodeMaskWriteRegister:
		return "Mask Write Register"
	case FuncCodeReadFIFOQueue:
		return "Read FIFO Queue"
	default:
		return ""
	}
}
