package bilibili

import (
	"github.com/eric2788/common-utils/request"
)

func XRespDecoder(b []byte, v interface{}) error {
	var resp XResp
	err := request.JsonDecoder(b, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return &APIError{
			Code:    resp.Code,
			Message: resp.Message,
		}
	}

	return request.JsonDecoder(b, v)
}

func V1RespDecoder(b []byte, v interface{}) error {
	var resp V1Resp
	err := request.JsonDecoder(b, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return &APIError{
			Code:    resp.Code,
			Message: resp.Message,
		}
	}

	return request.JsonDecoder(b, v)
}
