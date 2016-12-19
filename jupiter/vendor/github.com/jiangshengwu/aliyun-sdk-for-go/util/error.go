package util

import "fmt"

type SdkError struct {
	Resp ErrorResponse
	Url  string
}

func (err *SdkError) Error() string {
	return fmt.Sprintf("Aliyun SDK Request Error:\nURL: %s\nRequestId: %s\nHostId: %s\nCode: %s\nMessage: %s",
		err.Url, err.Resp.RequestId, err.Resp.HostId, err.Resp.Code, err.Resp.Message)
}
