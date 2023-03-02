package response

type ApiResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Error     string      `json:"error"`
	Data      interface{} `json:"data"`
	RequestID string      `json:"request_id"`
}

type PageResponse struct {
	Count   int64       `json:"count"`
	Results interface{} `json:"results"`
}
