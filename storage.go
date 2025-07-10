package modbus

import "sync"

type Storage struct {
	coils            [65536]bool
	discreteInputs   [65536]bool
	holdingRegisters [65536]uint16
	inputRegisters   [65536]uint16
	cmu              *sync.Mutex
	dmu              *sync.Mutex
	hmu              *sync.Mutex
	imu              *sync.Mutex
}

func (s *Storage) ReadCoils(address, quantity uint16) (results []bool, err error) {
	s.cmu.Lock()
	defer s.cmu.Unlock()

	results = make([]bool, quantity)
	copy(results, s.coils[address:address+quantity])
	return results, nil
}

func (s *Storage) ReadDiscreteInputs(address, quantity uint16) (results []bool, err error) {
	s.dmu.Lock()
	defer s.dmu.Unlock()

	results = make([]bool, quantity)
	copy(results, s.discreteInputs[address:address+quantity])
	return results, nil
}

func (s *Storage) WriteSingleCoil(address uint16, value bool) error {
	s.cmu.Lock()
	defer s.cmu.Unlock()

	s.coils[address] = value
	return nil
}

func (s *Storage) WriteMultipleCoils(address, quantity uint16, value []bool) error {
	s.cmu.Lock()
	defer s.cmu.Unlock()

	copy(s.coils[address:address+quantity], value)
	return nil
}

func (s *Storage) ReadInputRegisters(address, quantity uint16) (results []uint16, err error) {
	s.imu.Lock()
	defer s.imu.Unlock()

	results = make([]uint16, quantity)
	copy(results, s.inputRegisters[address:address+quantity])
	return results, nil
}

func (s *Storage) ReadHoldingRegisters(address, quantity uint16) (results []uint16, err error) {
	s.hmu.Lock()
	defer s.hmu.Unlock()

	results = make([]uint16, quantity)
	copy(results, s.holdingRegisters[address:address+quantity])
	return results, nil
}

func (s *Storage) WriteSingleRegister(address, value uint16) error {
	s.hmu.Lock()
	defer s.hmu.Unlock()

	s.holdingRegisters[address] = value
	return nil
}

func (s *Storage) WriteMultipleRegisters(address, quantity uint16, value []uint16) error {
	s.hmu.Lock()
	defer s.hmu.Unlock()

	copy(s.holdingRegisters[address:address+quantity], value)
	return nil
}

func (s *Storage) SetCoils(address, quantity uint16, value []bool) {
	s.cmu.Lock()
	defer s.cmu.Unlock()

	copy(s.coils[address:address+quantity], value)
}

func (s *Storage) SetDiscreteInputs(address, quantity uint16, value []bool) {
	s.dmu.Lock()
	defer s.dmu.Unlock()

	copy(s.discreteInputs[address:address+quantity], value)
}

func (s *Storage) SetHoldingRegisters(address, quantity uint16, value []uint16) {
	s.hmu.Lock()
	defer s.hmu.Unlock()

	copy(s.holdingRegisters[address:address+quantity], value)
}

func (s *Storage) SetInputRegisters(address, quantity uint16, value []uint16) {
	s.imu.Lock()
	defer s.imu.Unlock()

	copy(s.inputRegisters[address:address+quantity], value)
}
