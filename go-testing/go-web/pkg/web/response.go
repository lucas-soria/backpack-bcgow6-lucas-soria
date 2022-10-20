package web

import "strconv"

type Response struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int64, data interface{}, err string) Response {
	if code < 300 {
		return Response{
			Code:  strconv.FormatInt(code, 10),
			Data:  data,
			Error: "",
		}
	}
	return Response{
		Code:  strconv.FormatInt(code, 10),
		Data:  nil,
		Error: err,
	}
}
