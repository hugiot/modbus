package modbus

type FunctionCode uint8

const (
	FunctionCodeReadCoils              FunctionCode = 0x1
	FunctionCodeReadDiscreteInputs     FunctionCode = 0x2
	FunctionCodeReadHoldingRegisters   FunctionCode = 0x3
	FunctionCodeReadInputRegisters     FunctionCode = 0x4
	FunctionCodeWriteSingleCoil        FunctionCode = 0x5
	FunctionCodeWriteSingleRegister    FunctionCode = 0x6
	FunctionCodeWriteMultipleCoils     FunctionCode = 0xF
	FunctionCodeWriteMultipleRegisters FunctionCode = 0x10
)
