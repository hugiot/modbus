package modbus

type RequestHandler interface {
	HandleCoils(req *CoilsRequest) (res []bool, err error)
	HandleDiscreteInputs(req *DiscreteInputsRequest) (res []bool, err error)
	HandleHoldingRegisters(req *HoldingRegistersRequest) (res []uint16, err error)
	HandleInputRegisters(req *InputRegistersRequest) (res []uint16, err error)
}
