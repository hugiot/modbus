package modbus

import (
	"errors"
	"github.com/hugiot/modbus/crc"
)

type MBAPHeader struct {
	TransactionIdentifier uint16 `json:"transactionIdentifier"` // transaction identifier
	ProtocolIdentifier    uint16 `json:"protocolIdentifier"`    // protocol identifier
	Length                uint16 `json:"length"`                // length of the following data
	UnitIdentifier        byte   `json:"unitIdentifier"`        // unit identifier
}

type Request struct {
	MBAPHeader
	SlaveId      byte   `json:"slaveId"`      // slave address
	FunctionCode byte   `json:"functionCode"` // function code
	Address      uint16 `json:"address"`      // start address
	Quantity     uint16 `json:"quantity"`     // quantity
	ValueLength  byte   `json:"valueLength"`  // value length
	Values       []byte `json:"values"`       // value bytes
	original     []byte // original ADU
}

func (r *Request) Original() []byte {
	return r.original
}

func (r *Request) GetBoolSlice() []bool {
	results := bytesToBoolSlice(r.Values)
	if results == nil {
		return nil
	}

	if len(results) >= int(r.Quantity) {
		return results[:r.Quantity]
	}

	return results
}

func (r *Request) GetUint16Slice() []uint16 {
	results := bytesToUint16Slice(r.Values)
	if results == nil {
		return nil
	}

	if len(results) >= int(r.Quantity) {
		return results[:r.Quantity]
	}

	return results
}

func ReadRTURequest(adu []byte) (*Request, error) {
	if len(adu) == 0 {
		return nil, errors.New("modbus: empty application data unit")
	}

	if len(adu) < rtuMinSize {
		return nil, errors.New("modbus: RTU data unit too short")
	}

	// crc
	length := len(adu)
	sum := crc.Sum(adu[0 : length-2])
	checksum := uint16(adu[length-1])<<8 | uint16(adu[length-2])
	if checksum != sum {
		return nil, errors.New("modbus: crc checksum mismatch")
	}

	// slave id
	slaveId := adu[0]

	// protocol data unit
	code, address, quantity, valueLength, values, err := parseProtocolDataUnit(adu[1 : length-2])
	if err != nil {
		return nil, err
	}

	return &Request{
		SlaveId:      slaveId,
		FunctionCode: code,
		Address:      address,
		Quantity:     quantity,
		ValueLength:  valueLength,
		Values:       values,
		original:     adu,
	}, nil
}

func ReadTCPRequest(adu []byte) (*Request, error) {
	// todo something
	return nil, nil
}
