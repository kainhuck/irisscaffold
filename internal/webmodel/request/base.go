package request

type PageReq struct {
	Page     int `json:"page" url:"page" validate:"gte=1"`
	PageSize int `json:"page_size" url:"page_size" validate:"gte=0"`
}
