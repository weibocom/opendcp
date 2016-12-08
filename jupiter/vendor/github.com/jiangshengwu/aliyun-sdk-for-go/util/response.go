package util

type Response interface {
}

type BaseResponse struct {
	RequestId string `json:"RequestId"`
}

type ErrorResponse struct {
	BaseResponse
	HostId  string `json:"HostId"`
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

type PageResponse struct {
	TotalCount int `json:"TotalCount"`
	PageNumber int `json:"PageNumber"`
	PageSize   int `json:"PageSize"`
}
