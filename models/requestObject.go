package models

import "fmt"

type RequestObject struct {
	Host   string
	Path   string
	Method string
	Cookie string
}

func (rb *RequestObject) Print() {
	fmt.Println(rb.Host)
	fmt.Println(rb.Path)
	fmt.Println(rb.Method)
	fmt.Println(rb.Cookie)
}
