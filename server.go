package modbus

type Server interface {
	Start() error
	Stop() error
}

type serverImpl struct {
}

func (s *serverImpl) Start() error {
	//TODO implement me
	panic("implement me")
}

func (s *serverImpl) Stop() error {
	//TODO implement me
	panic("implement me")
}

func NewServer() Server {
	return nil
}
