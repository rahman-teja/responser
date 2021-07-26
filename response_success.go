package responser

import "net/http"

type SuccessResponse struct {
	Data       interface{}
	Meta       interface{}
	HttpStatus int
	Message    string
}

type TotalDataMeta struct {
	TotalData  int `json:"totalData"`
	TotalPages int `json:"totalPages"`
}

func NewSuccessResponse(data, meta interface{}, message string) *SuccessResponse {
	return &SuccessResponse{
		Data:       data,
		Meta:       meta,
		HttpStatus: http.StatusOK,
		Message:    message,
	}
}

func NewSuccessInsertResponse(data, meta interface{}, message string) *SuccessResponse {
	return &SuccessResponse{
		Data:       data,
		Meta:       meta,
		HttpStatus: http.StatusCreated,
		Message:    message,
	}
}

func (s SuccessResponse) GetData() interface{} {
	return s.Data
}

func (s SuccessResponse) GetHTTPStatus() int {
	return s.HttpStatus
}

func (s SuccessResponse) GetMessage() string {
	return s.Message
}

func (s SuccessResponse) GetError() error {
	return nil
}

func (s SuccessResponse) GetMeta() interface{} {
	return s.Meta
}
