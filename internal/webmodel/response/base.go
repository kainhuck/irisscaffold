package response

type ApiResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Error     string      `json:"error"`
	Data      interface{} `json:"data"`
	RequestID string      `json:"request_id"`
}
