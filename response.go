package responser

type Response interface {
	GetData() interface{}
	GetHTTPStatus() int
	GetMessage() string
	GetError() error
	GetMeta() interface{}
}
