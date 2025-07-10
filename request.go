package modbus

type Request struct {
	SlaveID      uint8
	FunctionCode int
	Data         []byte
}
