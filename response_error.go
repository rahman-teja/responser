package responser

type ErrorResponse struct {
	Data       interface{}
	HttpStatus int
	Message    string
	Err        error
}

func NewErrorResponse(data interface{}, httpstatus int, message string, err error) *ErrorResponse {
	return &ErrorResponse{
		Data:       data,
		HttpStatus: httpstatus,
		Message:    message,
		Err:        err,
	}
}

func (s ErrorResponse) GetData() interface{} {
	return s.Data
}

func (s ErrorResponse) GetHTTPStatus() int {
	return s.HttpStatus
}

func (s ErrorResponse) GetMessage() string {
	return s.Message
}

func (s ErrorResponse) GetError() error {
	return s.Err
}

func (s ErrorResponse) GetMeta() interface{} {
	return nil
}
