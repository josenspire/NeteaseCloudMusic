package models

import (
	"fmt"
)

type RequestObject struct {
	Host   string
	Path   string
	Method string
	Cookie string
	Data   interface{}
}

func (rb *RequestObject) Print() {
	fmt.Println(rb.Host)
	fmt.Println(rb.Path)
	fmt.Println(rb.Method)
	fmt.Println(rb.Cookie)
	fmt.Println(rb.Data)
}

// func main() {
// 	origin := `{"Username": "james", "Password": "123456"}`
//
// 	json.Unmarshal(origin, )
// }
