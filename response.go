package modbus

type ResponseWriter interface {
	Write([]byte) (int, error)
}
