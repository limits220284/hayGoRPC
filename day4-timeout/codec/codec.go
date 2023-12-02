package codec

// 进行消息的编码和解码过程
import "io"

type (
	Header struct {
		ServiceMethod string // format "Service.Method"
		Seq           uint64 // sequence number chosen by client
		Error         string
	}

	Codec interface {
		io.Closer
		ReadHeader(*Header) error
		ReadBody(interface{}) error
		Write(*Header, interface{}) error
	}
	
	NewCodecFunc func(io.ReadWriteCloser) Codec

	Type string
)

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
