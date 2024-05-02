package helper

type BaseResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseData(code int, message string, data interface{}) *BaseResponse {
	var response BaseResponse

	response.Code = code
	response.Message = message

	if data != nil {
		response.Data = data
	}

	return &response
}
