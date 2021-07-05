package utils

type HttpResp struct {
	Success bool        `json:"success"`
	ErrMsg  string      `json:"code"`
	Data    interface{} `json:"data"`
}

func NewHttpErrResp(errMsg string) *HttpResp {
	return &HttpResp{ErrMsg: errMsg}
}

func NewHttpSuccessResp(data interface{}) *HttpResp {
	return &HttpResp{
		Success: true,
		Data:    data,
	}
}
