package modbus

import (
	"encoding/binary"
	"errors"
	"fmt"
)

func byteToBoolSlice(b byte) []bool {
	var bs [8]bool
	for i := 0; i < 8; i++ {
		if b>>i&1 == 1 {
			bs[i] = true
		}
	}
	return bs[:]
}

func bytesToBoolSlice(bs []byte) []bool {
	if len(bs) == 0 {
		return nil
	}

	var results []bool
	for _, b := range bs {
		results = append(results, byteToBoolSlice(b)...)
	}
	return results
}

func bytesToUint16Slice(bs []byte) []uint16 {
	if len(bs) == 0 {
		return nil
	}

	var results []uint16
	for i := 0; i < len(bs); i += 2 {
		if i+1 < len(bs) {
			results = append(results, binary.BigEndian.Uint16(bs[i:i+2]))
		} else {
			results = append(results, uint16(bs[i]))
		}
	}
	return results
}

func boolSliceToBytes(v []bool) []byte {
	return nil
}

func uint16SliceToBytes(v []uint16) []byte {
	if len(v) == 0 {
		return nil
	}

	result := make([]byte, 0, 2*len(v))
	for _, val := range v {
		result = binary.BigEndian.AppendUint16(result, val)
	}

	return result
}

// parseProtocolCommonDataUnit
// [functionCode, CRCCheck) include function code, exclude CRC check
func parseProtocolDataUnit(adu []byte) (code byte, address, quantity uint16, valueLength byte, values []byte, err error) {
	if len(adu) == 0 {
		err = errors.New("modbus: Illegal data value")
		return
	}

	// function code & data
	code = adu[0]
	switch int(code) {
	case FuncCodeReadCoils, FuncCodeReadDiscreteInputs, FuncCodeReadHoldingRegisters, FuncCodeReadInputRegisters:
		if len(adu) != 5 {
			err = errors.New("modbus: Illegal data value")
			return
		}
		address = binary.BigEndian.Uint16(adu[1:3])
		quantity = binary.BigEndian.Uint16(adu[3:5])
	case FuncCodeWriteSingleCoil, FuncCodeWriteSingleRegister: // write single register
		if len(adu) != 5 {
			err = errors.New("modbus: Illegal data value")
			return
		}
		address = binary.BigEndian.Uint16(adu[1:3])
		quantity = 1
		values = adu[3:5]
	case FuncCodeWriteMultipleCoils: // write multiple coils
		if len(adu) < 7 {
			err = errors.New("modbus: Illegal data value")
			return
		}
		address = binary.BigEndian.Uint16(adu[1:3])
		quantity = binary.BigEndian.Uint16(adu[3:5])
		valueLength = adu[5]
		values = adu[6:]
	case FuncCodeWriteMultipleRegisters: // write multiple registers
		if len(adu) < 8 {
			err = errors.New("modbus: Illegal data value")
			return
		}
		address = binary.BigEndian.Uint16(adu[1:3])
		quantity = binary.BigEndian.Uint16(adu[3:5])
		valueLength = adu[5]
		values = adu[6:]
	default:
		err = fmt.Errorf("modbus: unsupported function code [%d]", code)
	}

	return
}
