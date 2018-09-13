package models

import "fmt"

type Response struct {
	Code int
	Data interface{}
	Message string
}

func (rs *Response) SetResponse(code int, data interface{}, message string) {
	rs.Code = code
	rs.Data = data
	rs.Message = message
}

func (rs *Response) PrintResponse() {
	fmt.Println(rs.Code)
	fmt.Println(rs.Data)
	fmt.Println(rs.Message)
}
