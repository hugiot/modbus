package modbus

import (
	"github.com/goburrow/serial"
	"sync"
)

type serialPort struct {
	serial.Config
	port serial.Port
	mu   sync.Mutex
}

func (s *serialPort) Connect() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.connect()
}

func (s *serialPort) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.close()
}

func (s *serialPort) Read(p []byte) (n int, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err = s.connect(); err != nil {
		return 0, err
	}

	return s.port.Read(p)
}

func (s *serialPort) Write(p []byte) (n int, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err = s.connect(); err != nil {
		return 0, err
	}

	return s.port.Write(p)
}

func (s *serialPort) connect() error {
	if s.port != nil {
		return nil
	}

	port, err := serial.Open(&s.Config)
	if err != nil {
		return err
	}
	s.port = port

	return nil
}

func (s *serialPort) close() error {
	if s.port == nil {
		return nil
	}

	if err := s.port.Close(); err != nil {
		return err
	}
	s.port = nil

	return nil
}

func createSerialPort(address string) serialPort {
	return serialPort{
		Config: serial.Config{
			Address:  address,
			BaudRate: BaudRate9600,
			DataBits: DataBits8,
			StopBits: StopBits1,
			Parity:   ParityEven,
		},
	}
}
