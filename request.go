package modbus

type Request struct {
	SlaveID  uint8        // slave address
	Function FunctionCode // function code
	Address  uint16       // starting address
}

type CoilsRequest struct {
	Request
	Quantity uint16
	Values   []bool
}

type DiscreteInputsRequest struct {
	Request
	Quantity uint16
}

type HoldingRegistersRequest struct {
	Request
	Quantity uint16
	Values   []uint16
}

type InputRegistersRequest struct {
	Request
	Quantity uint16
}
