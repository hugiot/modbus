package modbus

type ResponseWriter interface {
	Success() error
	SuccessWithUint16Slice([]uint16) error
	Error(err error) error
}

type Response struct {
	Request     *Request
	ValueLength byte   `json:"valueLength"` // value length
	Values      []byte `json:"values"`      // value bytes
	ErrorCode   byte   `json:"errorCode"`
}
