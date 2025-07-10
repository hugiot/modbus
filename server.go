package modbus

type Server struct {
	Handler Handler
}

func (s *Server) Start() error {
	return nil
}

func (s *Server) Stop() error {
	return nil
}
